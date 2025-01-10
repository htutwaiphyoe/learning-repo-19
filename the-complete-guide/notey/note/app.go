package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title should not be empty")
	}

	if content == "" {
		return nil, errors.New("content should not be empty")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("\nThe note title '%v' has the following content: \n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := fmt.Sprintf("%v.json", strings.ToLower(strings.ReplaceAll(note.Title, " ", "-")))

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
