package entitles

import "fmt"

type Diary struct {
	Id      int64
	Sender  string
	Title   string
	Content string
}

func (diary Diary) ToString() string {
	return fmt.Sprintf("id: %d\nsender: %s\ntitle: %s\ncontent:%s", diary.Id, diary.Sender, diary.Title, diary.Content)
}
