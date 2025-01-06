package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  ``,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := Task{
			ID:          currentID,
			Description: args[0],
			Status:      Todo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		tasks = append(tasks, task)
		currentID++
		err := saveTasks()
		if err != nil {
			fmt.Printf("Error saving task: %v\n", err)
			return
		}

		fmt.Printf("Added new task: %s (ID: %d)\n", task.Description, task.ID)
	},
}

func init() {}
