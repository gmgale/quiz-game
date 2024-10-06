package questions

import (
	"encoding/json"
	"os"
)

type Question struct {
	ID        string   `json:"id"`
	Text      string   `json:"text"`
	Options   []string `json:"options"`
	TimeLimit int      `json:"timeLimit"`
}

func LoadQuestions(filename string) ([]Question, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var questions []Question
	err = json.Unmarshal(data, &questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}
