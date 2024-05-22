package main

import (
	"log"
	"martext/editor"
	"martext/terminal"
	"os"
)

func setupTerm() {

	var fd int = int(os.Stdin.Fd())
	term, err := terminal.Init(fd)
	if err != nil {
		log.Fatal(err)
	}

	term, err = term.MakeRaw()
	defer term.Restore()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	setupTerm()

	// var editor editor.Editor = editor.Editor{}
	// editor.LoadFile
	editor, err := editor.LoadFile("./main.go")
	if err != nil {
		log.Fatal(err)
	}

	for {
		editor.ClearScreen()
		editor.Display()

		var input [1]byte
		os.Stdin.Read(input[:])

		if input[0] == 'q' {
			editor.ClearScreen()
			os.Exit(0)
		}
	}

}
