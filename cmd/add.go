package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "add a task",
	Long:  `adds task with specified description, update and creation time will be set to current`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idx, err := taskService.Add(args[0])
		if err != nil {
			return fmt.Errorf("could not delete task: %w", err)
		}

		fmt.Printf("Task %d added successfully\n", idx)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
