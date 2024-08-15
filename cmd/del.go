/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func delCmd() *cobra.Command {
	return &cobra.Command{
		Use: "del",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := query.Task.Order(query.Task.Id.Desc()).Find()
			if err != nil {
				printE(err.Error())
				return
			}

			if len(tasks) == 0 {
				printE("No tasks found")
				return
			}

			options := make([]string, len(tasks))
			for i, task := range tasks {
				options[i] = cast.ToString(task.Id)
			}

			var selects []string
			if err := survey.AskOne(&survey.MultiSelect{
				Message: "Select tasks to delete",
				Options: options,
				Description: func(value string, index int) string {
					return tasks[index].Title
				},
			}, &selects); err != nil {
				printE(err.Error())
				return
			}

			if _, err := query.Task.
				Where(query.Task.Id.In(cast.ToIntSlice(selects)...)).
				Delete(); err != nil {
				printE(err.Error())
				return
			}

			printS("")
		},
	}
}
