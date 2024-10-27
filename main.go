package main

import (
	"fmt"
	"github.com/felipeoli7eira/note-go/note"
	"bufio"
	"os"
	"strings"
)

func main() {
	title, content := getNote()

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	saveErr := newNote.Save()

	if saveErr != nil {
		fmt.Println("saving failed!")
		return
	}

	fmt.Println("saving succeeded!")
}

func getNote() (string, string) {
	var title string = getPrompt("note: ")
	var content string = getPrompt("content: ")

	return title, content
}

func getPrompt(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}
