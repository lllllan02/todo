/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func doneCmd() *cobra.Command {
	return &cobra.Command{
		Use: "done",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := query.Task.Where(query.Task.Done.Not()).Find()
			if err != nil {
				printE(err.Error())
				return
			}

			if len(tasks) == 0 {
				printS("No tasks to done")
				return
			}

			options := make([]string, len(tasks))
			for i, task := range tasks {
				options[i] = cast.ToString(task.Id)
			}

			var selects []string
			if err := survey.AskOne(&survey.MultiSelect{
				Message: "Select tasks to done",
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
				Update(query.Task.Done, true); err != nil {
				printE(err.Error())
				return
			}

			printS("")
		},
	}
}
