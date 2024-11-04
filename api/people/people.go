package people

import (
	"encoding/json"
	"errors"
	"regexp"
	"starwars/api/core"
	"starwars/api/film"
	"starwars/api/vehicle"
	"strconv"

	"github.com/graphql-go/graphql"
)

type PeopleSwapi struct {
	Name       string
	Birth_year string
	Eye_color  string
	Gender     string
	Hair_color string
	Height     string
	Mass       string
	Skin_color string
	Homeworld  string
	Films      []string
	Species    []string
	Starships  []string
	Vehicles   []string
	Url        string
	Created    string
	Edited     string
}

type People struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Gender   string   `json:"gender"`
	Films    []string `json:"films"`
	Vehicles []string `json:"vehicle"`
}

func Create(content PeopleSwapi) People {
	people := new(People)
	idRegex := regexp.MustCompile("[0-9]{1,}")
	rawId := idRegex.Find([]byte(content.Url))
	id, _ := strconv.Atoi(string(rawId))
	people.Name = content.Name
	people.Gender = content.Gender
	people.Films = content.Films
	people.Vehicles = content.Vehicles
	people.Id = id
	return *people
}

var GraphqlQueries = graphql.Fields{
	"people": &graphql.Field{
		Type:        GraphqlSchema,
		Description: "Get people with given id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if !ok {
				return nil, errors.New("missing id param")
			}

			url := core.SwapiUrl + "/people/" + strconv.Itoa(id)
			content, err := core.GetOrFetch(url, p.Context, *core.RedisClient, 0)

			if err != nil {
				return nil, err
			}

			var peopleSwapi PeopleSwapi
			parsingErr := json.Unmarshal([]byte(content), &peopleSwapi)

			if parsingErr != nil {
				return nil, parsingErr
			}

			people := Create(peopleSwapi)
			return people, nil
		},
	},
	"peoples": &graphql.Field{
		Type:        graphql.NewList(GraphqlSchema),
		Description: "Get people that matches the name param",
		Args: graphql.FieldConfigArgument{
			"search": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			search, ok := p.Args["search"].(string)
			if !ok {
				return nil, errors.New("missing search param")
			}

			url := core.SwapiUrl + "/people?search=" + search
			content, _ := core.GetOrFetch(url, p.Context, *core.RedisClient, 0)

			var peopleSwapi core.PaginatedResponse[PeopleSwapi]
			parsingErr := json.Unmarshal([]byte(content), &peopleSwapi) // Todo: Avoid casting to byte again

			if parsingErr != nil {
				return nil, parsingErr
			}

			var peoples []People = make([]People, len(peopleSwapi.Results))

			for i := range peopleSwapi.Results {
				peopleSwapi := peopleSwapi.Results[i]
				core.RedisClient.Set(p.Context, peopleSwapi.Url, peopleSwapi, 0)
				peoples[i] = Create(peopleSwapi)
			}

			return peoples, nil
		},
	},
}

var GraphqlSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "people",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
			"films": &graphql.Field{
				Type: graphql.NewList(film.GraphqlSchema),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					people, ok := p.Source.(People)
					if !ok {
						return nil, errors.New("could not parse source to people")
					}

					films, err := core.FetchRelationships(
						people.Films, 
						core.SwapiUrl, 
						core.RedisClient, 
						p.Context, 
						film.Create,
					)
					if err != nil {
						return nil, err
					}

					return films, nil
				},
			},
			"vehicles": &graphql.Field{
				Type: graphql.NewList(vehicle.GraphqlSchema),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					people, ok := p.Source.(People)
					if !ok {
						return nil, errors.New("could not parse source to people")
					}

					vehicles, err := core.FetchRelationships(
						people.Vehicles, 
						core.SwapiUrl, 
						core.RedisClient, 
						p.Context, 
						vehicle.Create,
					)
					if err != nil {
						return nil, err
					}

					return vehicles, nil
				},
			},
		},
	},
)
