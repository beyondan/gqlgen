package main

import (
	"log"
	"net/http"

	todo "github.com/beyondan/gqlgen/example/config"
	"github.com/beyondan/gqlgen/graphql/handler"
	"github.com/beyondan/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
