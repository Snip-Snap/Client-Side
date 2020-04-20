package main

import (
	// This is a named import from another local package. Need for dbconn methods.
	"api"
	"api/generated"
	"log"
	"net/http"
	"os"
	"github.com/99designs/gqlgen/handler"

)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	print("connecting to psql")
	api.ConnectPSQL()
	defer api.ClosePSQL()

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &api.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
