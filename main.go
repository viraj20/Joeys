package main

import (
	"log"
	"net/http"

	api "Joeys/pkg/api"
	v1 "Joeys/pkg/api/v1"
	build "Joeys/pkg/api/v1/build"
	commit "Joeys/pkg/api/v1/commit"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health-check", api.Get).Methods(http.MethodGet)
	apiV1SubRoute := r.PathPrefix("/api/v1").Subrouter()
	apiV1SubRoute.HandleFunc("/", v1.VersionGet)
	allCommitSubRoute := apiV1SubRoute.PathPrefix("/commit").Subrouter()
	allCommitSubRoute.HandleFunc("/", commit.Get)
	buildSubRoute := apiV1SubRoute.PathPrefix("/build").Subrouter()
	buildSubRoute.HandleFunc("/", build.Post).Methods(http.MethodPost)
	buildSubRoute.HandleFunc("/", build.Get).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
