/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := query.Task.Where(query.Task.Done.Not()).Order(query.Task.Id.Desc()).Find()
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

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
