package build

import (
	"Joeys/pkg/api/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type NameDesc struct {
	Name string
}

func Get(responseWriter http.ResponseWriter, request *http.Request) {
	var results []*NameDesc

	client := utils.ConnectMongo()
	cur, err := utils.GetMongoCollections(client, "myCollection").Find(context.TODO(), bson.D{{}})
	if err != nil {
		fmt.Print("1" + err.Error())
		//log.Fatal(err)
	}
	if cur != nil {
		//log.Fatal(err)

		for cur.Next(context.TODO()) {
			var result NameDesc
			err := cur.Decode(&result)

			if err != nil {
				fmt.Print("3" + err.Error())
				log.Fatal(err)
			}
			results = append(results, &result)

		}
		cur.Close(context.TODO())
	}
	utils.DisconnectMongo(client)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.WriteHeader(http.StatusOK)
	jsonResponse, _ := json.Marshal(results)
	responseWriter.Write([]byte(jsonResponse))
}
