package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// router.HandleFunc("/api/diary/findall", diary)

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
