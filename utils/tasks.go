package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo/models"

	"github.com/spf13/cobra"
)

func FindTaskByName(
	name string,
	callback func(int, *models.Task, []models.Task),
) (int, *models.Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("failed to load task to list: %v\n", err)
		return -1, nil, err
	}

	for i, t := range tasks {
		if name == t.Name {
			callback(i, &t, tasks)
			return i, &t, nil
		}
	}

	fmt.Println("Task not found:", name)
	return -1, nil, nil
}

// Get the task name from the user first look in the args and take the first one
// if not found, prompt the user for the task name
func GetTaskName(args []string) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}

	fmt.Print("Enter task name: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("failed to read input: %v\n", err)
		return "", err
	}

	return input[:len(input)-1], nil
}

// GetStdinReader returns a bufio.Reader for stdin
func GetStdinReader() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

// TaskNameCompletionWithPredicate provides shell completion for task names
// It returns a list of task names that match the input prefix.
// It uses a predicate function to filter tasks based on custom logic.
func TaskNameCompletionWithPredicate(
	_ *cobra.Command,
	_ []string,
	_toComplete string,
	predicate func(task models.Task) bool,
) ([]string, cobra.ShellCompDirective) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	var taskNames []string
	for _, task := range tasks {
		if predicate(task) && (len(_toComplete) == 0 || strings.HasPrefix(
			strings.ToLower(task.Name),
			strings.ToLower(_toComplete),
		)) {
			taskNames = append(taskNames, task.Name)
		}
	}

	return taskNames, cobra.ShellCompDirectiveNoFileComp
}

// TaskNameCompletion provides shell completion for task names
// It returns a list of task names that match the input prefix.
// notice it uses TaskNameCompletionWithPredicate with a predicate
// that always returns true so all tasks are included if you want
// to filter tasks use TaskNameCompletionWithPredicate directly
// or use MakeTaskNameCompletion for a cleaner interface.
func TaskNameCompletion(
	cmd *cobra.Command,
	args []string,
	_toComplete string,
) ([]string, cobra.ShellCompDirective) {
	return TaskNameCompletionWithPredicate(cmd, args, _toComplete, func(models.Task) bool {
		return true // Include all tasks
	})
}

func MakeTaskNameCompletion(predicate func(task models.Task) bool) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return TaskNameCompletionWithPredicate(cmd, args, toComplete, predicate)
	}
}

func HasTag(task models.Task, tag string) bool {
	for _, t := range task.Tags {
		if strings.EqualFold(t, tag) {
			return true
		}
	}
	return false
}

// TagCompletionWithPredicate provides shell completion for task names
// It returns a list of task names that match the input prefix.
// It uses a predicate function to filter tasks based on custom logic.
func TagCompletionWithPredicate(
	_ *cobra.Command,
	_ []string,
	_toComplete string,
	predicate func(task models.Task) bool,
) ([]string, cobra.ShellCompDirective) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	var uniqueTags = make(map[string]struct{})
	for _, task := range tasks {
		// skip tasks that do not match the predicate
		if !predicate(task) {
			continue
		}

		for _, tag := range task.Tags {
			if strings.HasPrefix(strings.ToLower(tag), strings.ToLower(_toComplete)) {
				uniqueTags[tag] = struct{}{}
			}
		}
	}
	var tags []string
	for tag := range uniqueTags {
		tags = append(tags, tag)
	}

	return tags, cobra.ShellCompDirectiveNoFileComp
}

// TagCompletion provides shell completion for task names
// It returns a list of task names that match the input prefix.
// notice it uses TaskNameCompletionWithPredicate with a predicate
//
//	that always returns true so all tasks are included if you want
//
// to filter tasks use TaskNameCompletionWithPredicate directly
// or use MakeTaskNameCompletion for a cleaner interface.
func TagCompletion(
	cmd *cobra.Command,
	args []string,
	_toComplete string,
) ([]string, cobra.ShellCompDirective) {
	return TagCompletionWithPredicate(cmd, args, _toComplete, func(models.Task) bool {
		return true // Include all tasks
	})
}

func MakeTagCompletion(predicate func(task models.Task) bool) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return TagCompletionWithPredicate(cmd, args, toComplete, predicate)
	}
}
