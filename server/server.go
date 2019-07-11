package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	gqlgen_clean_architecture "github.com/ymmooot/gqlgen-clean-architecture"
	"github.com/ymmooot/gqlgen-clean-architecture/app/infrastructure"
	"github.com/ymmooot/gqlgen-clean-architecture/app/interfaceadapter/gateway"
	"github.com/ymmooot/gqlgen-clean-architecture/app/interfaceadapter/resolver"
	"github.com/ymmooot/gqlgen-clean-architecture/app/usecase"
)

const defaultPort = "8080"

func main() {
	port := getPort()

	sql := infrastructure.NewSqlHandler() // todo: connection pool
	gateway := gateway.NewArticleGateway(sql)
	usecase := usecase.NewArticleUseCaseImpl(gateway)
	resolver := resolver.NewResolver(usecase)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		gqlgen_clean_architecture.NewExecutableSchema(
			gqlgen_clean_architecture.Config{
				Resolvers: resolver,
			},
		),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}
