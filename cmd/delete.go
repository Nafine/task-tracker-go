package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete specific task",
	Long:  "deletes task by id if task with specified id was found",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idx, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task id '%s': must be an integer", args[0])
		}

		found, err := taskService.Delete(idx)
		if err != nil {
			return fmt.Errorf("could not delete task: %w", err)
		}

		if !found {
			return fmt.Errorf("task with id %d not found", idx)
		}

		fmt.Printf("Task %d deleted successfully\n", idx)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
