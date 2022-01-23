package main

import (
	"notes-app/commands"
	"notes-app/editor"
)

func main() {
	commands.Execute()

	vi := new(editor.EditorVi)
	editor.EditFile(vi, "testPath")
}
