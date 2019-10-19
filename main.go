package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/201RichK/graphql/graphql-test/queries"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	log "github.com/sirupsen/logrus"
)

func main() {

	shematConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Rootquery",
			Fields: queries.GetRootField(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: nil,
		}),
	}

	shemat, err := graphql.NewSchema(shematConfig)
	if err != nil {
		log.Error("Failed to create a new shemat", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &shemat,
	})

	http.Handle("/", httpHandler)

	port := os.Getenv("port")
	if port == "" {
		port = ":8080"
	}
	log.Info("listen on port 8080")

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(port), nil))
}
