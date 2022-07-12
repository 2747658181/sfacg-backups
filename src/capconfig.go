package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ChapterJson struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Index    int    `json:"index"`
	IsVip    bool   `json:"is_vip"`
	VolumeID string `json:"volume_id"`
	Content  string `json:"content"`
}
type TestChapterConfig struct {
	ChapterInfo []ChapterJson
	BookName    string
}

func TestInit(bookName string) *TestChapterConfig {
	var AutoGenerateds []ChapterJson
	if data, ok := ioutil.ReadFile("./config/" + bookName + ".json"); ok == nil {
		if err := json.Unmarshal(data, &AutoGenerateds); err != nil {
			fmt.Println(err)
		}
	}
	return &TestChapterConfig{ChapterInfo: AutoGenerateds, BookName: bookName}
}

func (is *TestChapterConfig) In(ChapID string) bool {
	for _, s := range is.ChapterInfo {
		if s.ID == ChapID {
			return true
		}
	}
	return false
}