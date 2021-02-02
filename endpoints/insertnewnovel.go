package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/0Hidder/novelupdatesscrapperv1/connection"
	"github.com/0Hidder/novelupdatesscrapperv1/mongodb"
)

//InsertNewNovel inserts new novel in the database
func InsertNewNovel(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	novel := connection.ReturnNewNovelStructure()
	_ = json.NewDecoder(request.Body).Decode(&novel)
	client, _ := mongodb.GetMongoClient()
	collection := client.Database(mongodb.DB).Collection(mongodb.NOVELS)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection.InsertOne(ctx, novel)
	response.Write([]byte(`{ "message": "New novel (` + novel.Name + `) inserted successfully" }`))
	json.NewEncoder(response)

}
