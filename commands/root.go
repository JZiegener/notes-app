package commands

import (
	"fmt"
	"notes-app/notebook"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "notes",
	Short: "A command line application for creating, searching, and managing notes",
	Long:  "A Command line application for creating, searching and, managing notes",
	//	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		backpack, err := notebook.InitializeBackpack()
		if err != nil {
			fmt.Println("error creating notebook")
			return
		}
		notebook, err := backpack.CreateNotebook("TestNotebook")
		if err != nil {
			return
		}
		fmt.Println(notebook.GetName())

		notebooks, err := backpack.GetAllNoteBooks()
		if err != nil {
			fmt.Println("error loading notebook")
			return
		}

		for index, element := range notebooks {
			fmt.Println("Notebook", index, "Name is", element.GetName(), "create Time", element.GetCreateTime())
		}
	},
}

// Execute is the main entry point for command parsing and execution
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
