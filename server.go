package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/go-graphql-api-starter/config"
	"github.com/go-graphql-api-starter/data/db"
	"github.com/go-graphql-api-starter/graph"
	"github.com/go-graphql-api-starter/graph/generated"
)

const defaultPort = "4000"

func init() {
	config.InitSettings()
}

func main() {

	db, err := db.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
