package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/201RichK/graphql/db"
	mutation "github.com/201RichK/graphql/mutations"
	"github.com/201RichK/graphql/root"
	_ "github.com/201RichK/graphql/types"
	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	log "github.com/sirupsen/logrus"
)

func main() {
	shematConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "query",
			Fields: root.GetMotCleField(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "mutation",
			Fields: mutation.GetRootField(),
		}),
	}

	shemat, err := graphql.NewSchema(shematConfig)
	if err != nil {
		panic(err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema:     &shemat,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/")
	if err != nil {
		panic(err)
	}
	http.Handle("/", httpHandler)
	http.Handle("/graphiql", graphiqlHandler)

	port := os.Getenv("port")
	if port == "" {
		port = ":8080"
	}
	log.Info("listen on port 8080")

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(port), nil))
}
