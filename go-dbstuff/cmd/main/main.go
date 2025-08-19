package main

import (
	"log"
	"net/http"

	"github.com/Wanderer0074348/GoServeIt/go-dbstuff/pkg/config"
	"github.com/Wanderer0074348/GoServeIt/go-dbstuff/pkg/routes"
	"github.com/gorilla/mux"
)

const port = ":8000"

func main() {
	r := mux.NewRouter()

	config.ConnectDb()
	db := config.GetDB()

	routes.RegisterBookStoreRoutes(r, db)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(port, r))
}
