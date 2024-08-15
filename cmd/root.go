/*
Copyright © 2024 lllllan
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	limit  = -1
	offset = -1
)

func Execute(args []string) error {
	root := &cobra.Command{Use: "todo"}

	// PersistentFlags
	root.PersistentFlags().IntVar(&limit, "limit", -1, "limit specify the number of records to be retrieved")
	root.PersistentFlags().IntVar(&offset, "offset", -1, "offset specify the number of records to skip before starting to return the records")

	// Commands
	root.AddCommand(addCmd())
	root.AddCommand(delCmd())
	root.AddCommand(ListCmd())
	root.AddCommand(doneCmd())

	root.SetArgs(args)
	return root.Execute()
}
