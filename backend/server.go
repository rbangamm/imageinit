package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/kamva/mgm/v3"
	config2 "github.com/rbangamm/imageinit/config"
	"github.com/rbangamm/imageinit/graph"
	"github.com/rbangamm/imageinit/graph/generated"
	"github.com/rbangamm/imageinit/repository/user"
	userservice "github.com/rbangamm/imageinit/service/user"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	config := config2.Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	port := config.ServerConfig.Port

	router := chi.NewRouter()

	uri := fmt.Sprintf("mongodb://localhost:%s", config.MongoConfig.Port)
	err = mgm.SetDefaultConfig(nil, "imageinit", options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	userRepo := user.NewRepository()
	userService := userservice.NewService(&config, userRepo)

	router.Use(userservice.Middleware(userService))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: graph.NewResolver(userService),
	}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
