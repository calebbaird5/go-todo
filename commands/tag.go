package commands

import (
	"bufio"
	"fmt"
	"os"
	"todo/models"
	"todo/utils"

	"github.com/spf13/cobra"
)

func tag(cmd *cobra.Command, args []string) {
	taskName, err := utils.GetTaskName(args)
	if err != nil {
		return
	}

	// tags, err := utils.LoadTags()

	tags := []string{}
	if len(args) > 1 {
		tags = args[1:]
	} else {
		fmt.Print("Enter the tag to apply to the task: ")
		reader := bufio.NewReader(os.Stdin)
		tag, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("failed to read input: %v\n", err)
			return
		}
		tags = []string{tag}
	}

	_, _, err = utils.FindTaskByName(
		taskName,
		func(i int, task *models.Task, tasks []models.Task) {
			if cmd.Flag("remove").Changed {
				for _, tagToDelete := range tags {
					for tagI, existingTag := range tasks[i].Tags {
						if existingTag == tagToDelete {
							tasks[i].Tags = append(tasks[i].Tags[:tagI], tasks[i].Tags[tagI+1:]...)
							break
						}
					}
				}
			} else if cmd.Flag("replace").Changed {
				tasks[i].Tags = tags
			} else {
				tasks[i].Tags = append(tasks[i].Tags, tags...)
			}
			err = utils.SaveTasks(tasks)
			if err != nil {
				fmt.Printf("Failed to save updated tasks: %s\n", err)
				return
			}
		},
	)
	if err != nil {
		fmt.Printf("failed to tag the task: %v\n", err)
		return
	}

	fmt.Println("Tagged the task:", taskName)
}

var TagCmd = &cobra.Command{
	Use:     "tag [Task Name] [Tag1] [Tag2] ...",
	Aliases: []string{"t"},
	Short:   "Adds Tag(s) to a task",
	Run:     tag,
	Args:    cobra.MinimumNArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			return utils.TaskNameCompletion(cmd, args, toComplete)
		}
		if cmd.Flag("remove").Changed {
			return utils.TagCompletionWithPredicate(
				cmd, args, toComplete,
				func(task models.Task) bool {
					return task.Name == args[0]
				},
			)
		} else {
			return utils.TagCompletion(cmd, args, toComplete)
		}
	},
}

func init() {
	TagCmd.Flags().BoolP("remove", "r", false,
		"Removes tags from a task rather than adding them",
	)
	TagCmd.Flags().BoolP("replace", "R", false,
		"Replace any existing tags with the new ones",
	)
}
