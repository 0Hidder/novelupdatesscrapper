package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/0Hidder/novelupdatesscrapperv1/connection"
	"github.com/0Hidder/novelupdatesscrapperv1/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllNovels gets all novels from the database
func GetAllNovels(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	novels := connection.ReturnNewNovelStructureArray()
	client, _ := mongodb.GetMongoClient()
	collection := client.Database(mongodb.DB).Collection(mongodb.NOVELS)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		novel := connection.ReturnNewNovelStructure()
		cursor.Decode(&novel)
		novels = append(novels, novel)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(novels)
}
