package models

import (
	"database/sql"

	"github.com/fummicc1/diary_api_go/src/entitles"
)

type DiaryModel struct {
	DB *sql.DB
}

func (diaryModel DiaryModel) FindAll() (diary []entitles.Diary, err error) {
	rows, err := diaryModel.DB.Query("select * from diary")
	if err != nil {
		return nil, err
	} else {
		var diaries []entitles.Diary
		for rows.Next() {
			var id int64
			var sender string
			var title string
			var content string
			err2 := rows.Scan(&id, &sender, &title, &content)
			if err2 != nil {
				return nil, err2
			} else {
				diary := entitles.Diary{
					Id:      id,
					Sender:  sender,
					Title:   title,
					Content: content,
				}
				diaries = append(diaries, diary)
			}
		}
		return diaries, nil
	}
}