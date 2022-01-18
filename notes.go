package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
)

func launchVi (string filename) {
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
	
	
	//launchVi("testFile")
	
}
