package main

import (
	"log"
	"net/http"

	api "github.com/viraj20/Joeys/pkg/api"
	v1 "github.com/viraj20/Joeys/pkg/api/v1"
	build "github.com/viraj20/Joeys/pkg/api/v1/build"
	commit "github.com/viraj20/Joeys/pkg/api/v1/commit"

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
	log.Fatal(http.ListenAndServe(":8080", r))
}
