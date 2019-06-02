package main

import (
	"fmt"
	"net/http"

	"github.com/fummicc1/diary_api_go/src/apis/diary"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/diary/findall", diary.FindAll).Methods("GET")
	router.HandleFunc("/api/diary/search/{sender}", diary.Search).Methods("GET")
	router.HandleFunc("/api/diary/insert", diary.Insert).Methods("POST")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
