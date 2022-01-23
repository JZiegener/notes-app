package commands

import (
	"fmt"
	"notes-app/editor"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "notes",
	Short: "A command line application for creating, searching, and managing notes",
	Long:  "A Command line application for creating, searching and, managing notes",
	//	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yep running some code")

		vi := new(editor.EditorVi)
		editor.EditFile(vi, "testPath")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
└── project
    ├── go.mod
    ├── main.go
	├── commands
	|   └── root.go
	└── editor
	    └── editor.go

Want to reference editor from commands
*/
