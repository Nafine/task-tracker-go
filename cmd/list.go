package cmd

import (
	"fmt"

	"github.com/Nafine/task-tracker/internal/model"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List all tasks with corresponding statuses",
	Long: `list is for listing all tasks with 1 of 3 statuses (in-progress, todo, done)
provide no argument for listing all tasks`,
	ValidArgs: []cobra.Completion{"in-progress", "done", "todo"},
	Args:      cobra.MatchAll(cobra.MaximumNArgs(3), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("All tasks:\n", taskService.List())
			return
		}

		fmt.Println(fmt.Sprintf("All tasks with status %v:\n%s", args, filterTasks(taskService.List(),
			func(task model.Task) bool {
				for _, statusStr := range args {
					status, _ := model.ParseStatus(statusStr)
					if task.Status == status {
						return true
					}
				}
				return false
			})))
	},
}

func filterTasks(tasks model.Tasks, filter func(model.Task) bool) model.Tasks {
	newTasks := make(model.Tasks, 0)
	for _, task := range tasks {
		if filter(task) {
			newTasks = append(newTasks, task)
		}
	}
	return newTasks
}

func init() {
	rootCmd.AddCommand(listCmd)
}
