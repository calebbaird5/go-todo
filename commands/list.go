package commands

import (
	"fmt"
	"strings"
	"todo/utils"

	"github.com/spf13/cobra"
)

func columnValue(value string, width int) string {
	if len(value) == 0 {
		return "--"
	} else if len(value) <= width {
		return value
	}
	return value[:width-3] + "..."
}

func list(cmd *cobra.Command, args []string) {
	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Printf("failed to load task to list: %v\n", err)
		return
	}

	if cmd.Flag("all").Changed && cmd.Flag("completed").Changed {
		fmt.Println("Cannot use --all and --completed flags together")
		return
	}

	if cmd.Flag("tag").Changed && cmd.Flag("untagged").Changed {
		fmt.Println("Cannot use --tag and --untagged flags together")
		return
	}

	if cmd.Flag("untagged").Changed && cmd.Flag("tagged").Changed {
		fmt.Println("Cannot use --untagged and --tagged flags together")
		return
	}

	// if the long flag is set, print the column header
	columnFormat := "%-10s %-12s %-20s %-12s\n"
	if cmd.Flag("long").Changed {
		if cmd.Flag("all").Changed {
			columnFormat = "%-10s %-12s %-20s %-12s %-12s\n"
			fmt.Printf(columnFormat, "Name", "Tags", "Description", "Created At", "Completed At")
		} else {
			fmt.Printf(columnFormat, "Name", "Tags", "Description", "Created At")
		}
	}
	for _, task := range tasks {
		// skip completed tasks unless the --all or --completed flag is set
		if !cmd.Flag("all").Changed && !cmd.Flag("completed").Changed && task.CompletedAt != nil {
			continue
		}

		// if the completed flag is set, skip tasks that are not completed
		if cmd.Flag("completed").Changed && task.CompletedAt == nil {
			continue
		}

		// if tag flag is set, skip tasks that do not match any of the specified tags
		if cmd.Flag("tag").Changed {
			tagValues, _ := cmd.Flags().GetStringArray("tag")
			hasAnyTag := false
			for _, tag := range tagValues {
				if utils.HasTag(task, tag) {
					hasAnyTag = true
					break
				}
			}
			if len(tagValues) > 0 && !hasAnyTag {
				continue
			}
		}

		// if tagged flag is set, skip tasks that do not have any tags
		if cmd.Flag("tagged").Changed && len(task.Tags) == 0 {
			continue
		}

		// if untagged flag is set, skip tasks that have tags
		if cmd.Flag("untagged").Changed && len(task.Tags) > 0 {
			continue
		}

		if cmd.Flag("long").Changed {
			dateFormat := "2006/01/02"
			var formatedCompletedAt string
			if task.CompletedAt != nil {
				formatedCompletedAt = task.CompletedAt.Format(dateFormat)
			} else {
				formatedCompletedAt = ""
			}

			if cmd.Flag("all").Changed {
				fmt.Printf(
					columnFormat,
					columnValue(task.Name, 10),
					columnValue(strings.Join(task.Tags, ","), 12),
					columnValue(task.Description, 20),
					columnValue(task.CreatedAt.Format(dateFormat), 12),
					columnValue(formatedCompletedAt, 12),
				)
			} else {
				fmt.Printf(
					columnFormat,
					columnValue(task.Name, 10),
					columnValue(strings.Join(task.Tags, ","), 12),
					columnValue(task.Description, 20),
					columnValue(task.CreatedAt.Format(dateFormat), 12),
				)
			}
		} else {
			if task.CompletedAt != nil {
				fmt.Printf(" ✓ %s\n", task.Name)
			} else {
				fmt.Printf(" - %s\n", task.Name)
			}
		}
	}
}

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "List all tasks",
	Run:     list,
	// provide tab completion for the tag flag
	ValidArgsFunction: utils.TagCompletion,
	Args:              cobra.NoArgs,
}

func init() {
	ListCmd.Flags().BoolP("all", "a", false, "List all tasks, including completed ones")
	ListCmd.Flags().BoolP("completed", "c", false, "List only completed tasks")
	ListCmd.Flags().BoolP("tagged", "T", false, "List only tagged tasks")
	ListCmd.Flags().BoolP("untagged", "U", false, "List only untagged tasks")
	ListCmd.Flags().BoolP("long", "l", false, "List tasks in multi column format")
	ListCmd.Flags().StringArrayP("tag", "t", nil, "List tasks matching any of the specified tags (e.g. --tag work --tag urgent)")
	ListCmd.RegisterFlagCompletionFunc("tag", utils.TagCompletion)
}
