package vehicle

import "github.com/graphql-go/graphql"

type VehicleSwapi struct {
	Name                   string
	Model                  string
	Vehicle_class          string
	Manufacturer           string
	Length                 string
	Cost_in_credits        string
	Crew                   string
	Passengers             string
	Max_atmosphering_speed string
	Cargo_capacity         string
	Consumables            string
	Films                  []string
	Pilots                 []string
	Url                    string
	Created                string
	Edited                 string
}


type Vehicle struct {
	Name         string `json:"name"`
	Model        string `json:"model"`
	VehicleClass string `json:"vehicleClass"`
	Manufacturer string `json:"manufacturer"`
}

func Create(content VehicleSwapi) Vehicle {
	vehicle := new(Vehicle)
	vehicle.Name = content.Name
	vehicle.Manufacturer = content.Manufacturer
	vehicle.Model = content.Model
	vehicle.VehicleClass = content.Vehicle_class
	return *vehicle
}

var GraphqlSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "vehicle",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"model": &graphql.Field{
				Type: graphql.String,
			},
			"vehicleClass": &graphql.Field{
				Type: graphql.String,
			},
			"manufacturer": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)