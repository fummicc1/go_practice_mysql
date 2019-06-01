package product

import (
	"encoding/json"
	"net/http"

	"github.com/fummicc1/diary_api_go/src/config"
	"github.com/fummicc1/diary_api_go/src/models"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		diaryModel := models.DiaryModel
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
