package stories

import (
	"encoding/json"
)

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type Story map[string]Chapter

func ParseStory(jsonByte []byte, story *Story) error {
	err := json.Unmarshal(jsonByte, story)
	if err != nil {
		return err
	}
	return nil
}
