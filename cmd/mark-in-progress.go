package cmd

import (
	"fmt"
	"strconv"

	"github.com/Nafine/task-tracker/internal/model"
	"github.com/spf13/cobra"
)

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress [id]",
	Short: "Mark specified task as in-progress",
	Long:  "marks specified task by id as in-progress if such task was found",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idx, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task id '%s': must be an integer", args[0])
		}

		found, err := taskService.Mark(idx, model.StatusInProgress)
		if err != nil {
			return fmt.Errorf("could not update task: %w", err)
		}

		if !found {
			return fmt.Errorf("task with id %d not found", idx)
		}

		fmt.Printf("Task %d marked successfully\n", idx)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)
}
