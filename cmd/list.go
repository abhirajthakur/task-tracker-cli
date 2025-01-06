package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var filteredTasks []Task

		if len(args) > 0 {
			status := TaskStatus(strings.ToLower(args[0]))

			switch status {
			case Todo, InProgress, Done:
				for _, task := range tasks {
					if task.Status == status {
						filteredTasks = append(filteredTasks, task)
					}
				}
			default:
				fmt.Printf("Invalid status: %s\nValid statuses are: todo, in-progress, done\n", args[0])
				return
			}
		} else {
			filteredTasks = tasks
		}

		if len(filteredTasks) == 0 {
			if len(args) > 0 {
				fmt.Printf("No tasks found with status: %s\n", args[0])
			} else {
				fmt.Println("No tasks found.")
			}
			return
		}

		if len(args) > 0 {
			fmt.Printf("Tasks with status '%s':\n\n", args[0])
		} else {
			fmt.Printf("All tasks:\n\n")
		}

		fmt.Println("ID | Status      | Description")
		fmt.Println("----------------------------")
		for _, task := range filteredTasks {
			fmt.Printf("%2d | %-11s | %s\n", task.ID, task.Status, task.Description)
		}
	},
}

func init() {}
