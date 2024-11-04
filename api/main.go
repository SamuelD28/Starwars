package main

import (
	"context"
	"net/http"
	"starwars/api/people"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "root",
					Fields: people.GraphqlQueries,
				}),
		},
	)

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		graphqlHandler.ContextHandler(context.Background(), writer, request)
	})

	http.ListenAndServe(":8080", nil)
}
