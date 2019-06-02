package models

import (
	"database/sql"

	"github.com/fummicc1/diary_api_go/src/entitles"
)

type DiaryModel struct {
	DB *sql.DB
}

func (diaryModel DiaryModel) FindAll() (diaries []entitles.Diary, err error) {
	rows, err := diaryModel.DB.Query("select * from diary")
	if err != nil {
		return nil, err
	} else {
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

func (diaryModel DiaryModel) Search(sender string) (diaries []entitles.Diary, err error) {
	rows, err := diaryModel.DB.Query("select * from diary where sender like ?", "%"+sender+"%")
	if err != nil {
		return nil, err
	} else {
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

func (diaryModel DiaryModel) Insert(sender string, title string, content string) (diary entitles.Diary, err error) {
	insert, err := diaryModel.DB.Prepare("INSERT INTO diary(sender, title, content) VALUES(?, ?, ?, ?)")
	if err != nil {
		return entitles.Diary.CreateErrorDiary(), err
	}
	result, err2 := insert.Exec(sender, title, content)
	if err2 != nil {
		return nil, err2
	}
	id, err3 := result.LastInsertId()
	if err3 != nil {
		return nil, err3
	}
	diary = entitles.Diary{
		Id:      id,
		Sender:  sender,
		Title:   title,
		Content: content,
	}
	return diary, nil
}
