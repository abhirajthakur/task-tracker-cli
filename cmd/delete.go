package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Delete a task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 8)
		if err != nil {
			fmt.Printf("Invalid task ID: %s\n", args[0])
			return
		}

		taskFound := false
		for i := range tasks {
			if tasks[i].ID == uint8(id) {

				tasks = append(tasks[:i], tasks[i+1:]...)

				taskFound = true
				err := saveTasks()
				if err != nil {
					fmt.Printf("Error saving task: %v\n", err)
					return
				}
				fmt.Printf("Deleted task %d\n", id)
				break
			}
		}

		if !taskFound {
			fmt.Printf("No task found with ID: %d\n", id)
		}
	},
}

func init() {}
