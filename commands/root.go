package commands

import (
	//"github.com/spf13/cobra"
	"context"
	"fmt"
	"notes-app/notebook"
	"os"

	"github.com/DavidGamba/go-getoptions"
)

// var rootCmd = &cobra.Command{
// 	Use:   "notes",
// 	Short: "A command line application for creating, searching, and managing notes",
// 	Long:  "A Command line application for creating, searching and, managing notes",
// 	//	Args: cobra.MinimumNArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {

// 		backpack, err := notebook.InitializeBackpack()
// 		if err != nil {
// 			fmt.Println("error creating notebook")
// 			return
// 		}
// 		notebook, err := backpack.CreateNotebook("TestNotebook")
// 		if err != nil {
// 			return
// 		}
// 		fmt.Println(notebook.GetName())

// 		notebooks, err := backpack.GetAllNoteBooks()
// 		if err != nil {
// 			fmt.Println("error loading notebook")
// 			return
// 		}

// 		for index, element := range notebooks {
// 			fmt.Println("Notebook", index, "Name is", element.GetName(), "create Time", element.GetCreateTime())
// 		}
// 	},
// }

func NewCommand(parent *getoptions.GetOpt) *getoptions.GetOpt {
	create := parent.NewCommand("create", "Create stuff")

	createNote := create.NewCommand("note", "create a new note ")
	createNote = createNote.SetCommandFn(RunNoteCreate)
	createNote.String("title", "", createNote.Required("Note Title must be specified to create a note"))

	createNoteBook := create.NewCommand("notebook", "Creates a new notebook")
	createNoteBook.SetCommandFn(RunNoteBookCreate)

	createNoteBook.String("name", "", createNote.Required("Notebook must have a name"))

	return create
}

func RunNoteCreate(ctx context.Context, opt *getoptions.GetOpt, args []string) error {
	fmt.Printf("Note Create %s!\n", opt.Value("title"))
	return nil
}

func RunNoteBookCreate(ctx context.Context, opt *getoptions.GetOpt, args []string) error {
	fmt.Printf("Notebook %s!\n", opt.Value("name"))

	backpack, err := notebook.InitializeBackpack()
	if err != nil {
		fmt.Println("error creating notebook")
		return nil
	}
	notebook, err := backpack.CreateNotebook("TestNotebook")
	if err != nil {
		return err
	}
	fmt.Println(notebook.GetName())

	notebooks, err := backpack.GetAllNoteBooks()
	if err != nil {
		fmt.Println("error loading notebook")
		return err
	}

	for index, element := range notebooks {
		fmt.Println("Notebook", index, "Name is", element.GetName(), "create Time", element.GetCreateTime())
	}

	return nil
}

// Execute is the main entry point for command parsing and execution
func Execute(args []string) int {
	opt := getoptions.New()
	opt.SetUnknownMode(getoptions.Fail)
	NewCommand(opt)

	remaining, err := opt.Parse(args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}

	fmt.Printf("Remaning cli args: %v\n", remaining)

	//Listen for system interrupts
	ctx, cancel, done := getoptions.InterruptContext()
	defer func() { cancel(); <-done }()

	opt.Dispatch(ctx, remaining)

	//getoptions.InterruptContext()
	//defer func() { cancel(); <- done }()

	return 0
	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}
