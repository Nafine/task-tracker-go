package cmd

import (
	"github.com/Nafine/task-tracker/internal/tasks"
	"github.com/spf13/cobra"
)

var taskService *tasks.Service

var rootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     "task-cli",
	Short:   "task-cli is a manager for simple scheduling of your tasks ",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		srv, err := tasks.NewService()
		if err != nil {
			return err
		}
		taskService = srv
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
