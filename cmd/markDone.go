package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var MarkDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Updates the status of a task as done",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 8)
		if err != nil {
			fmt.Printf("Invalid task ID: %s\n", args[0])
			return
		}

		taskFound := false
		for i := range tasks {
			if tasks[i].ID == uint8(id) {
				tasks[i].Status = Done

				taskFound = true

				err := saveTasks()
				if err != nil {
					fmt.Printf("Error saving task: %v\n", err)
					return
				}
				fmt.Printf("Marked task %d as done\n", id)
				break
			}
		}

		if !taskFound {
			fmt.Printf("No task found with ID: %d\n", id)
		}
	},
}

func init() {}