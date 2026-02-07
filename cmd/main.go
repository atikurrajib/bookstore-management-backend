package main

import (
	"log"
	"net/http"

	"github.com/atikurrajib/go-bookstore/pkg/config"
	"github.com/atikurrajib/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	config.Connect()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}
