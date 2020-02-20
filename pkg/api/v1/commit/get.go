package commit

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Joeys/pkg/api/v1/model"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

//Get all the commits tags for the passed repo
func Get(responseWriter http.ResponseWriter, request *http.Request) {
	repo, _ := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/viraj20/golang-practice",
	})
	ref, _ := repo.Head()
	cIter, _ := repo.Log(&git.LogOptions{From: ref.Hash()})
	fmt.Println(cIter)
	var response []model.CommitResponse
	cIter.ForEach(func(c *object.Commit) error {
		commitResponse := model.CommitResponse{c.ID().String(), c.Message, "", c.Author.String()}
		response = append(response, commitResponse)
		return nil
	})
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	jsonString, _ := json.Marshal(response)
	responseWriter.Write([]byte(jsonString))
}
