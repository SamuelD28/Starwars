package film

import "github.com/graphql-go/graphql"

type FilmSwapi struct {
	Title         string
	Episode_id    int
	Opening_crawl string
	Director      string
	Producer      string
	Release_date  string
	Species       []string
	Starships     []string
	Vehicles      []string
	Characters    []string
	Planets       []string
	Url           string
	Created       string
	Edited        string
}

type Film struct {
	Title       string `json:"title"`
	EpisodeId   int    `json:"episodeId"`
	ReleaseDate string `json:"releaseDate"`
}

func Create(content FilmSwapi) Film {
	film := new(Film)
	film.Title = content.Title
	film.EpisodeId = content.Episode_id
	film.ReleaseDate = content.Release_date
	return *film
}

var GraphqlSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "film",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"episodeId": &graphql.Field{
				Type: graphql.Int,
			},
			"releaseDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
