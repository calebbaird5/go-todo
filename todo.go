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
			" - Marking tasks as done\n" +
			" - Deleting tasks\n" +
			" - Tagging tasks",
	}

	var addCmd = &cobra.Command{
		Use:     "add",
		Aliases: []string{"a"},
		Short:   "Add a new task",
		Run:     commands.Add,
	}

	var listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"l", "ls"},
		Short:   "List all tasks",
		Run:     commands.List,
	}

	var doneCmd = &cobra.Command{
		Use:     "done",
		Aliases: []string{"d"},
		Short:   "Mark a task as done",
		Run:     func(cmd *cobra.Command, args []string) { fmt.Println("done command placeholder") },
	}

	var deleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del"},
		Short:   "Delete a task",
		Run:     func(cmd *cobra.Command, args []string) { fmt.Println("delete command placeholder") },
	}

	var tagCmd = &cobra.Command{
		Use:     "tag",
		Aliases: []string{"t"},
		Short:   "Tag a task",
		Run:     func(cmd *cobra.Command, args []string) { fmt.Println("tag command placeholder") },
	}

	rootCmd.AddCommand(addCmd, listCmd, doneCmd, deleteCmd, tagCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
