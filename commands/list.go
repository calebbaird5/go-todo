package commands

import (
	"fmt"
	"todo/utils"

	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, args []string) {
	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Printf("failed to load task to list: %v\n", err)
		return
	}
	for _, task := range tasks {
		fmt.Printf(" - %s\n", task.Description)
	}
}
