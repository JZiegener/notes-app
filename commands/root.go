package commands 

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)



var rootCmd = &cobra.Command{
	Use:	"notes",
	Short: "A command line application for creating, searching, and managing notes",
	Long: "A Command line application for creating, searching and, managing notes",
//	Args: cobra.MinimumNArgs(1),
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("yep running some code")
//	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}


