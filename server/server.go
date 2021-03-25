package main

import (
	"fmt"
	"net/http"

	"github.com/0Hidder/novelupdatesscrapperv1/endpoints"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the application...")

	router := mux.NewRouter()

	router.HandleFunc("/insertNovel", endpoints.InsertNewNovel).Methods("POST")
	router.HandleFunc("/getAllNovels", endpoints.GetAllNovels).Methods("GET")
	http.ListenAndServe(":12345", router)

	fmt.Println("Running...")

}
