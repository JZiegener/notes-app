package main

import (
	"log"
	"os/exec"
	"os"
	"github.com/JZiegener/notes-app/commands"
)

func launchVi (filename string){
	cmd := exec.Command("vi", filename)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()

	log.Printf("waiting for command to finish\n")
	err := cmd.Wait()
	log.Printf("Command Finished with  %s\n", err)
}




func main() {
	_ = commands.Execute()

	
}
