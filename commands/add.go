package commands

import (
	"fmt"
	"strings"
	"time"
	"todo/models"
	"todo/utils"

	"github.com/spf13/cobra"
)

func add(cmd *cobra.Command, args []string) {
	taskName, err := utils.GetTaskName(args)
	if err != nil {
		return
	}

	task := models.Task{
		Name:      taskName,
		CreatedAt: time.Now(),
	}

	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Printf("failed to load tasks: %v\n", err)
		return
	}
	// Append the new task
	existingTask := false
	for _, t := range tasks {
		if t.Name == task.Name {
			existingTask = true
			break
		}
	}

	// make sure the task does not already exist
	if existingTask {
		fmt.Println("Task already exists:", taskName)
		return
	}

	if cmd.Flags().Changed("description") && cmd.Flags().Changed("interactive-description") {
		fmt.Println("Cannot use both --description and --interactive-description flags together.")
		return
	}

	if cmd.Flags().Changed("description") {
		description, _ := cmd.Flags().GetString("description")
		task.Description = description
	} else if cmd.Flags().Changed("interactive-description") {
		fmt.Print("Enter task description: ")
		reader := utils.GetStdinReader()
		description, _ := reader.ReadString('\n')
		task.Description = strings.TrimSpace(description)
	}

	// Add the new task to the list and save the new task list
	tasks = append(tasks, task)
	utils.SaveTasks(tasks)

	fmt.Println("Added task:", taskName)
}

var AddCmd = &cobra.Command{
	Use:     "add [Task Name]",
	Aliases: []string{"a"},
	Short:   "Add a new task",
	Run:     add,
	Args:    cobra.MaximumNArgs(1),
}

func init() {
	AddCmd.Flags().StringArrayP("tags", "t", nil, "Tags to apply to the task")
	AddCmd.RegisterFlagCompletionFunc("tag", utils.TagCompletion)
	AddCmd.Flags().StringP("description", "d", "", "Description of the task")
	AddCmd.Flags().BoolP("interactive-description", "D", false, "Enter a description interactively")
}
