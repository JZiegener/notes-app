package editor

import (
	"log"
	"os"
	"os/exec"
)

type editor interface {
	editFile(path string) bool
}

type EditorVi struct {
}

func EditFile(vi *EditorVi, path string) bool {
	cmd := exec.Command("vi", path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()

	log.Printf("waiting for command to finish\n")
	err := cmd.Wait()
	log.Printf("Command Finished with  %s\n", err)
	if err != nil {
		return false
	}
	return true
}
