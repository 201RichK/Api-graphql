package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/201RichK/graphql/db"
	mutation "github.com/201RichK/graphql/mutations"
	"github.com/201RichK/graphql/queries"
	_ "github.com/201RichK/graphql/types"
	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	log "github.com/sirupsen/logrus"
)

func main() {

	shematConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootField(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutation.GetRootField(),
		}),
	}

	shemat, err := graphql.NewSchema(shematConfig)
	if err != nil {
		panic(err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &shemat,
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
	log.Info("listen on port 8090")

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(port), nil))
}
