package models

import (
	"database/sql"
)

type DiaryModel struct {
	DB *sql.DB
}

func (diaryModel DiaryModel) FindAll() (diary []DiaryModel, err error) {
	rows, err := diaryModel.DB.Query("select * from diary")
}
