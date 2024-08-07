package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/infra/handler/graphql/graph"

	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/service"
)


func graphql(port string, svc service.ServiceOrderInterface ) {
	const defaultPort = "8082"
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Service: svc,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
