package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	c := generated.Config{Resolvers: &resolvers.Resolver{UserResolver: resolvers.NewUserResolver()}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.Authenticate(srv.ServeHTTP))

	log.Printf("connect to graphql: http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
