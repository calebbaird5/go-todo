package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/models"
)

var (
	HomeDir   string
	ConfigDir string
	TasksFile string
)

func init() {
	var err error

	HomeDir, err = os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("failed to get home directory: %w", err))
	}
	ConfigDir = HomeDir + "/.config/todo"
	EnsureConfigDir()
	if err = os.MkdirAll(ConfigDir, 0755); err != nil {
		panic(fmt.Errorf("failed to create config directory: %w", err))
	}
	TasksFile = ConfigDir + "/tasks.json"
	EnsureTasksFile()
}

func EnsureTasksFile() error {
	if _, err := os.Stat(TasksFile); os.IsNotExist(err) {
		_, err := os.Create(TasksFile)
		if err != nil {
			return fmt.Errorf("failed to create tasks file: %w", err)
		}
	}

	return nil
}

func EnsureConfigDir() error {
	if _, err := os.Stat(ConfigDir); os.IsNotExist(err) {
		if err := os.MkdirAll(ConfigDir, 0755); err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	return nil
}

func SaveTasks(tasks []models.Task) error {

	file, err := os.Create(TasksFile)
	if err != nil {
		return fmt.Errorf("failed to create tasks file: %w", err)
	}
	defer file.Close()
	if err := json.NewEncoder(file).Encode(tasks); err != nil {
		return fmt.Errorf("failed to encode tasks: %w", err)
	}

	return nil
}

func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task
	// Read existing tasks if file exists
	if file, err := os.Open(TasksFile); err == nil {
		defer file.Close()
		if err := json.NewDecoder(file).Decode(&tasks); err != nil && err.Error() != "EOF" {
			return nil, fmt.Errorf("failed to decode tasks: %w", err)
		}
	}
	return tasks, nil
}
