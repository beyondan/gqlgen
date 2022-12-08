package main

import (
	"log"
	"net/http"

	"github.com/beyondan/gqlgen/example/selection"
	"github.com/beyondan/gqlgen/graphql/handler"
	"github.com/beyondan/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Selection Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
