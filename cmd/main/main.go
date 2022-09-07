package main

import (
	"bookmanagementapp/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe("localhost:9010", r)

}
