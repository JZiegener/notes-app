package commands 

import (
	"github.com/spf13/cobra"
	"fmt"
)


var (

rootCmd = &cobra.Command{
	Use:	"notes",
	Short: "A command line application for creating, searching, and managing notes",
	Long: "A Command line application for creating, searching and, managing notes",
//	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yep running some code")
	},
}

versionCmd = &cobra.Command{
	Use: "version",
	Short: "Prints the application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0.1")
	},

}

)

func Execute() error {
	return rootCmd.Execute()
}


