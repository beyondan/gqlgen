package main

import (
	"log"
	"net/http"

	"github.com/beyondan/gqlgen/example/scalars"
	"github.com/beyondan/gqlgen/graphql/handler"
	"github.com/beyondan/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", handler.NewDefaultServer(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
