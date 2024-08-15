/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/lllllan02/todo/todo"
	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	return &cobra.Command{
		Use: "add",
		Run: func(cmd *cobra.Command, args []string) {
			task := todo.Task{}

			if err := survey.Ask([]*survey.Question{
				{
					Name: "Title",
					Prompt: &survey.Input{
						Message: "Title",
						Default: "Unnamed",
					},
				},
				{
					Name: "Start Time",
					Prompt: &survey.Input{
						Message: "StartTime",
						Default: time.Now().Format(time.DateTime),
					},
				},
				{
					Name: "Due Date",
					Prompt: &survey.Input{
						Message: "DueDate",
						Default: time.Now().Add(24 * time.Hour).Format(time.DateTime),
					},
				},
			}, &task); err != nil {
				printE(err.Error())
				return
			}

			if err := query.Task.Create(&task); err != nil {
				printE(err.Error())
				return
			}
			printS("")
		},
	}
}
