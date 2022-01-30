package commands

import (
	"fmt"
	"notes-app/notebook"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "notes-app",
	Short: "A command line application for creating, searching, and managing notes",
	Long:  "A Command line application for creating, searching and, managing notes",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create stuff",
	Long:  "Used for creating new notes, notebooks, or storage locations",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var createNotebookCmd = &cobra.Command{
	Use:   "notebook <Notebook Name> <Storage Location>",
	Short: "Create a notebook",
	Long:  "Used for creating new notes, notebooks, or storage locations",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		backpack, err := notebook.InitializeBackpack()
		if err != nil {
			fmt.Println("error creating notebook")
			return
		}
		notebook, err := backpack.CreateNotebook(args[0], args[1])
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

var createNoteCmd = &cobra.Command{
	Use:   "note <Note title> [Flags]",
	Short: "Create a new note ",
	Long:  "Used for creating new notes, notebooks, or storage locations",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: Not yet implemented")
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Used to set application default behavior",
	Long: `Configure defaults for note behavior
\tDefault Notebook
\tDefault Template\n`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: Not yet implemented")
	},
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Used to search for different types",
	Long:  `Find notes, tags, or notebooks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: Not yet implemented")
	},
}

var notebookName string

// Execute is the main entry point for command parsing and execution
func Execute() {
	createCmd.AddCommand(createNotebookCmd)
	createNoteCmd.Flags().StringVarP(&notebookName, "notebook name", "b", "", "Name of the notebook to use")
	createCmd.AddCommand(createNoteCmd)

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(findCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
