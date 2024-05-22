package editor

import (
	"fmt"
	"os"
	"strings"
)

type Editor struct {
	Content       [][]string
	stringContent string
}

const (
	CLEARSCREEN = "\x1b[3J\x1b[2J"
	RESETCURSOR = "\x1b[H"
)

func (editor *Editor) ResetCursor() {
	fmt.Print(RESETCURSOR)
}

func (editor *Editor) ClearScreen() {
	fmt.Print(CLEARSCREEN)
	fmt.Print(RESETCURSOR)
}

func (editor *Editor) Display() {
	fmt.Print(editor.stringContent)
	// for i := 0; i < len(editor.Content); i++ {
	// 	for j := 0; j < len(editor.Content[i]); j++ {
	// 		fmt.Printf("%s", editor.Content[i][j])
	// 	}
	// }
}

func LoadFile(filePath string) (*Editor, error) {

	file, err := openFile(filePath)
	if err != nil {
		return nil, err
	}

	fileStats, err := getFileStats(file)
	if err != nil {
		return nil, err
	}

	fileContent, err := getFileContents(file, fileStats)
	if err != nil {
		return nil, err
	}

	var editor *Editor = &Editor{}
	editor.Content = parseFileContents(fileContent)
	editor.stringContent = strings.Replace(string(fileContent), "\n", "\r\n", -1)

	return editor, nil
}

func getFileContents(file *os.File, fileStats os.FileInfo) (string, error) {
	var fileBtyes []byte = make([]byte, fileStats.Size())
	if _, err := file.Read(fileBtyes); err != nil {
		return "", err
	}

	return string(fileBtyes), nil
}

func getFileStats(file *os.File) (os.FileInfo, error) {
	fileStats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return fileStats, nil
}

func openFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	file.Stat()
	if err != nil {
		return nil, err
	}

	return file, nil
}

func parseFileContents(fileContents string) [][]string {
	lines := strings.Split(fileContents, "\n")

	contents := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		chars := strings.Split(lines[i], "")
		chars = append(chars, "\r\n")
		contents[i] = chars
	}

	return contents
}
