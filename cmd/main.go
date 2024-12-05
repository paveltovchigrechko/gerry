package main

import (
	"fmt"
	"log"

	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	cfg, err := config.LoadConfig("gqlgen.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.LoadSchema()
	if err != nil {
		log.Fatal(err)
	}

	for _, type_ := range cfg.Schema.Types {
		fmt.Println(type_.Name)
		for _, field := range type_.Fields {
			fmt.Println(field.Name)
		}
		fmt.Println()
	}
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
