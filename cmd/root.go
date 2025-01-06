package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type TaskStatus string

const (
	Todo       TaskStatus = "todo"
	InProgress TaskStatus = "in-progress"
	Done       TaskStatus = "done"
)

const DataFile string = "task-tracker.json"

type Task struct {
	ID          uint8
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var (
	currentID uint8 = 1
	tasks     []Task
)

func loadTasks() {
	data, err := os.ReadFile(DataFile)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error reading tasks file:", err)
		}
		return
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error parsing tasks:", err)
		return
	}

	for _, task := range tasks {
		if task.ID >= currentID {
			currentID = task.ID + 1
		}
	}
}

func saveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding tasks: %w", err)
	}

	err = os.WriteFile(DataFile, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing tasks file: %w", err)
	}

	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-tracker-cli",
	Short: "An app to track your tasks and manage your to-do list.",
	Long: `Add, Update, and Delete tasks
Mark a task as in progress or done
List all tasks
List all tasks that are done
List all tasks that are not done
List all tasks that are in progress
  `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		loadTasks()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(UpdateCmd)
	rootCmd.AddCommand(MarkDoneCmd)
	rootCmd.AddCommand(MarkInProgressCmd)
	rootCmd.AddCommand(DeleteCmd)
}
