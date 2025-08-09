package commands

import (
	"context"
	"fmt"
	"todo/utils"

	"github.com/urfave/cli/v3"
)

func List(ctx context.Context, cmd *cli.Command) error {
	tasks, err := utils.LoadTasks()
	if err != nil {
		return fmt.Errorf("failed to load task to list: %w", err)
	}

	for _, task := range tasks {
		fmt.Printf(" - %s\n", task.Description)
	}

	return nil
}
