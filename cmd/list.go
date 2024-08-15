/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"github.com/lllllan02/todo/todo"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, _, err := query.Task.FindByPage(offset, limit)
			if err != nil {
				printE(err.Error())
			}

			todo.TaskTable(tasks).Render()
		},
	}
}
