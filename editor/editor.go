package editor

import (
	"log"
	"os"
	"os/exec"
)

type editor interface {
	EditFile(path string) bool
}

// EditFile launches the env variable defined "EDITOR" with the specified path
func EditFile(path string) bool {
	editorCommand := os.Getenv("EDITOR")

	switch editorCommand {
	case "vi":
		return launchCliEditor(editorCommand, path)
	case "code":
		log.Fatal("VS code editor not yet implemented")
		return false
	}
	return false
}

func launchCliEditor(command, path string) bool {
	cmd := exec.Command(command, path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()

	log.Printf("waiting for command to finish\n")
	err := cmd.Wait()
	log.Printf("Command Finished with  %s\n", err)

	return err == nil
}
