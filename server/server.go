package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	gqlgen_clean_architecture "github.com/ymmooot/gqlgen-clean-architecture"
	"github.com/ymmooot/gqlgen-clean-architecture/app/config/dig"
	"github.com/ymmooot/gqlgen-clean-architecture/app/interfaceadapter/resolver"
)

const defaultPort = "8080"

func main() {

	c, _ := dig.BuildDigDependencies()

	err := c.Invoke(func(r *resolver.Resolver) error {
		port := getPort()

		http.Handle("/", handler.Playground("GraphQL playground", "/query"))
		http.Handle("/query", handler.GraphQL(
			gqlgen_clean_architecture.NewExecutableSchema(
				gqlgen_clean_architecture.Config{
					Resolvers: r,
				},
			),
		))

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))

		return nil
	})

	if err != nil {
		panic(err)
	}

}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}
