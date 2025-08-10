package commands

import (
	"fmt"
	"todo/models"
	"todo/utils"

	"github.com/spf13/cobra"
)

func delete(cmd *cobra.Command, args []string) {
	taskDescription, err := utils.GetTaskName(args)
	if err != nil {
		return
	}

	var tasks []models.Task
	utils.FindTaskByName(
		taskDescription,
		func(i int, task *models.Task, _tasks []models.Task) {
			tasks = append(_tasks[:i], _tasks[i+1:]...)
		},
	)

	err = utils.SaveTasks(tasks)
	if err != nil {
		fmt.Printf("failed to delete the task: %v\n", err)
		return
	}

	fmt.Println("Deleted task:", taskDescription)

}

var DeleteCmd = &cobra.Command{
	Use:               "delete [task name]",
	Aliases:           []string{"d"},
	Short:             "Delete a task",
	Run:               delete,
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: utils.TaskNameCompletion,
}
