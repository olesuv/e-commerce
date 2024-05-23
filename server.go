package main

import (
	"context"
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
	"server.go/handlers"
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

	debug := true

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}
	if appEnv == "production" {
		debug = false
	}

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            debug,
	}).Handler)

	builder := configs.NewRedisClientBuilder()
	builder.WithAddr("127.0.0.1:" + os.Getenv("REDIS_PORT")).WithPassword(os.Getenv("REDIS_PASSWORD"))

	rdb, err := builder.Build()
	if err != nil {
		log.Fatal("server: redis connection failed, details: ", err)
	}
	if rdb.Ping(context.Background()).Err() != nil {
		log.Fatal("server: redis ping failed")
	}

	verifyEmailHandler := handlers.NewUserEmailVerificationHandler(rdb)
	router.HandleFunc("/verify", verifyEmailHandler.VerifyUserEmailHandler)

	c := generated.Config{Resolvers: &resolvers.Resolver{
		UserResolver:  resolvers.NewUserResolver(rdb),
		OrderResolver: resolvers.NewOrderResolver(),
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to graphql: http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
