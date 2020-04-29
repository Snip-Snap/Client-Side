package main

import (
	// This is a named import from another local package. Need for dbconn methods.
	"api"
	"api/auth"
	"api/generated"

	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
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
	router := chi.NewRouter()

	router.Use(auth.Middleware(api.DB))

	srv := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &api.Resolver{}}))
	router.Handle("/", playground.Handler("Barbershop", "/query"))
	router.Handle("/query", srv)

	print("\nRunning on localhost port 8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}
