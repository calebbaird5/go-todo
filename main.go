package main

import (
	"fmt"
	"os"
	"todo/commands"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "A CLI TODO tracker",
		Long: "A CLI TODO tracker.\n\nSupports:\n" +
			" - Adding tasks\n" +
			" - Listing tasks\n" +
			" - Marking tasks completed\n" +
			" - Deleting tasks\n" +
			" - Tagging tasks",
	}

	rootCmd.AddCommand(
		commands.AddCmd,
		commands.ListCmd,
		commands.CompleteCmd,
		commands.DeleteCmd,
		commands.TagCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
