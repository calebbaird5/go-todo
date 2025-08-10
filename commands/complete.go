package commands

import (
	"fmt"
	"time"
	"todo/models"
	"todo/utils"

	"github.com/spf13/cobra"
)

func complete(cmd *cobra.Command, args []string) {
	taskDescription, err := utils.GetTaskName(args)
	if err != nil {
		return
	}

	_, _, err = utils.FindTaskByName(
		taskDescription,
		func(i int, task *models.Task, tasks []models.Task) {
			now := time.Now()
			tasks[i].CompletedAt = &now
			err = utils.SaveTasks(tasks)
			if err != nil {
				fmt.Printf("Failed to save updated tasks: %s\n", err)
				return
			}
		},
	)
	if err != nil {
		fmt.Printf("failed to mark task completed: %v\n", err)
		return
	}

	fmt.Println("Marked task completed:", taskDescription)
}

var CompleteCmd = &cobra.Command{
	Use:     "complete [task name]",
	Aliases: []string{"c"},
	Short:   "Mark a task complete",
	Run:     complete,
	Args:    cobra.MaximumNArgs(1),
	ValidArgsFunction: utils.MakeTaskNameCompletion(
		// Only show tasks that are not completed
		func(task models.Task) bool { return task.CompletedAt == nil },
	),
}
