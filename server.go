package main

import (
	"github.com/kons16/graphql-todo/repository/redis"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kons16/graphql-todo/graph"
	"github.com/kons16/graphql-todo/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	rdMap, err := redis.NewRedisDB()
	if err != nil {
		log.Fatal(err)
	}

	taskRepo := redis.NewTaskRepository(rdMap)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{TaskRepo: taskRepo}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
