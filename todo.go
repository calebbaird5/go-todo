package main

import (
	"context"
	"log"
	"os"
	"todo/commands"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := (&cli.Command{
		Name: "TODO",
		Description: "A CLI TODO tracker.\n\nSupports:\n" +
			" - Adding tasks\n" +
			" - Listing tasks\n" +
			" - Marking tasks as done\n" +
			" - Deleting tasks\n" +
			" - Tagging tasks\n",
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a new task",
				Action:  commands.Add,
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls"},
				Usage:   "List all tasks",
				Action:  commands.List,
			},
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "Mark a task as done",
				Action:  func(context.Context, *cli.Command) error { return nil }, // Placeholder for add action
			},
			{
				Name:    "delete",
				Aliases: []string{"del"},
				Usage:   "Delete a task",
				Action:  func(context.Context, *cli.Command) error { return nil }, // Placeholder for add action
			},
			{
				Name:    "tag",
				Aliases: []string{"t"},
				Usage:   "Tag a task",
				Action:  func(context.Context, *cli.Command) error { return nil }, // Placeholder for add action
			},
		},
	})
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
