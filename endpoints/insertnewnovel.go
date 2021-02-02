package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/0Hidder/novelupdatesscrapperv1/connection"
)

//InsertNewNovel inserts new novel in the database
func InsertNewNovel(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	novel := connection.ReturnNewNovelStructure()
	_ = json.NewDecoder(request.Body).Decode(&novel)

}
