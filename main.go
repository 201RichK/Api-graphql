package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/201RichK/graphql/db"
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
			Fields: root.GetMutationField(),
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
		port = ":8090"
	}
	log.Warnf("listen on port http://127.0.O.1%s", port)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(port), nil))
}
