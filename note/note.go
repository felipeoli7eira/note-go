package note

import (
	"time"
	"errors"
	"fmt"
	"strings"
	"encoding/json"
	"os"
)

type Note struct {
	Note       string    `json:"note"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
}

func New(note string, content string) (Note, error) {
	if note == "" || content == "" {
		return Note{}, errors.New("invalid note data")
	}

	return Note{
		Note: note,
		Content: content,
		Created_at: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("note: %v. content: %v", note.Note, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Note, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName + ".json", json, 0644)
}
