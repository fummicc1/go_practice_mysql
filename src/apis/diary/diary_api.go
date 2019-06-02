package diary

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/fummicc1/diary_api_go/src/config"
	"github.com/fummicc1/diary_api_go/src/entitles"
	"github.com/fummicc1/diary_api_go/src/models"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		diaryModel := models.DiaryModel{
			DB: db,
		}
		diaries, err2 := diaryModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, diaries)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sender := vars["sender"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		diaryModel := models.DiaryModel{
			DB: db,
		}
		diaries, err2 := diaryModel.Search(sender)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, diaries)
		}
	}
}

func Insert(response http.ResponseWriter, request *http.Request) {
	body, cannnotGetBodyError := request.GetBody()
	if cannnotGetBodyError != nil {
		respondWithError(response, http.StatusBadRequest, cannnotGetBodyError.Error())
		return
	}
	var diary entitles.Diary
	if err := json.NewDecoder(body).Decode(&diary); err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	db, err1 := config.GetDB()
	if err1 != nil {
		respondWithError(response, http.StatusBadRequest, err1.Error())
	} else {
		diaryModel := models.DiaryModel{
			DB: db,
		}
		err2 := diaryModel.Insert(diary.Sender, diary.Title, diary.Content)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, diary)
		}
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
