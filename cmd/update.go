package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [description]",
	Short: "Update task with specified id",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		idx, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task id '%s': must be an integer", args[0])
		}

		description := args[1]

		found, err := taskService.Update(idx, description)
		if err != nil {
			return fmt.Errorf("could not update task: %w", err)
		}

		if !found {
			return fmt.Errorf("task with id %d not found", idx)
		}

		fmt.Printf("Task %d updated successfully\n", idx)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
