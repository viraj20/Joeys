package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	v1 "Joeys/pkg/api/v1"
	api "Joeys/pkg/api"
	commit "Joeys/pkg/api/v1/commit"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health-check", api.Get).Methods(http.MethodGet)
	apiV1SubRoute := r.PathPrefix("/api/v1").Subrouter()
	apiV1SubRoute.HandleFunc("/", v1.VersionGet)
	allCommitSubRoute := apiV1SubRoute.PathPrefix("/commits").Subrouter()
	allCommitSubRoute.HandleFunc("/", commit.Get)
	log.Fatal(http.ListenAndServe(":8080", r))
}
