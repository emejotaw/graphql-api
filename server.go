package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/emejotaw/graphql-api/config"
	"github.com/emejotaw/graphql-api/graph"
	"github.com/emejotaw/graphql-api/internal/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := config.NewDatabase()
	productService := service.NewProductService(db)
	orderService := service.NewOrderService(db)
	orderProductService := service.NewOrderProductService(db)
	resolver := graph.NewResolver(productService, orderService, orderProductService)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
