package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update [taskID] [new description]",
	Short: "update a task's description",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 8)
		if err != nil {
			fmt.Printf("Invalid task ID: %s\n", args[0])
			return
		}

		taskFound := false
		for i := range tasks {
			if tasks[i].ID == uint8(id) {
				oldDescription := tasks[i].Description
				tasks[i].Description = args[1]
				tasks[i].UpdatedAt = time.Now()
				taskFound = true

				err := saveTasks()
				if err != nil {
					fmt.Printf("Error saving task: %v\n", err)
					return
				}

				fmt.Printf("Updated task %d:\n", id)
				fmt.Printf("Old description: %s\n", oldDescription)
				fmt.Printf("New description: %s\n", tasks[i].Description)
				break
			}
		}

		if !taskFound {
			fmt.Printf("No task found with ID: %d\n", id)
		}
	},
}

func init() {}
