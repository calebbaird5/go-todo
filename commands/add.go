package commands

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
	"todo/models"
	"todo/utils"

	"github.com/urfave/cli/v3"
)

func Add(ctx context.Context, cmd *cli.Command) error {
	taskDescription := cmd.Args().Get(0)

	// if the task was not provided prompt for it now allow multi words
	if taskDescription == "" {
		fmt.Print("Enter task description: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}
		taskDescription = input[:len(input)-1] // Remove the newline character
	}

	task := models.Task{
		Description: taskDescription,
		CreatedAt:   time.Now(),
	}

	tasks, err := utils.LoadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w", err)
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
		return nil
	}

	// Add the new task to the list and save the new task list
	tasks = append(tasks, task)
	utils.SaveTasks(tasks)

	fmt.Println("Added task:", taskDescription)
	// Write tasks back to file
	return nil
}
