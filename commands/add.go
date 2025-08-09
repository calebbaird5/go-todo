package commands

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"todo/models"
	"todo/utils"

	"github.com/spf13/cobra"
)

func Add(cmd *cobra.Command, args []string) {
	taskDescription := ""
	if len(args) > 0 {
		taskDescription = args[0]
	}
	if taskDescription == "" {
		fmt.Print("Enter task description: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("failed to read input: %v\n", err)
			return
		}
		taskDescription = input[:len(input)-1]
	}
	task := models.Task{
		Description: taskDescription,
		CreatedAt:   time.Now(),
	}
	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Printf("failed to load tasks: %v\n", err)
		return
	}
	// Append the new task
	existingTask := false
	for _, t := range tasks {
		if t.Description == task.Description {
			existingTask = true
			break
		}
	}

	// make sure the task does not already exist
	if existingTask {
		fmt.Println("Task already exists:", taskDescription)
		return
	}

	// Add the new task to the list and save the new task list
	tasks = append(tasks, task)
	utils.SaveTasks(tasks)

	fmt.Println("Added task:", taskDescription)
	// Write tasks back to file
	return
}
