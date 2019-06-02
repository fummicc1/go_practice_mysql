package entitles

import "fmt"

type Diary struct {
	Id      int64  `json:"id"`
	Sender  string `json:"sender"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateErrorDiary() Diary {
	return Diary{
		Id:      0,
		Sender:  "",
		Title:   "",
		Content: "",
	}
}

func (diary Diary) ToString() string {
	return fmt.Sprintf("id: %d\nsender: %s\ntitle: %s\ncontent:%s", diary.Id, diary.Sender, diary.Title, diary.Content)
}
