package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"server.go/configs"
	"server.go/graph/generated"
	"server.go/graph/resolvers"
	"server.go/middleware"
)

const defaultPort = "8080"

func main() {
	configs.LoadEnv()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(middleware.Middleware())

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	c := generated.Config{Resolvers: &resolvers.Resolver{UserResolver: resolvers.NewUserResolver()}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to graphql: http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
