package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountUrl  string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogUrrl string `envconfig:"CATALOG_SERVICE_URL"`
	OrderUrl    string `envconfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	s, err := NewGraphQLServer(cfg.AccountUrl, cfg.CatalogUrrl, cfg.OrderUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	http.Handle("/graphql", handler.New(s.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("ramkrishna", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
