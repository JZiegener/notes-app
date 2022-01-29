package main

import (
	"notes-app/commands"
	"os"
)

func main() {
	os.Exit(commands.Execute(os.Args))
}
