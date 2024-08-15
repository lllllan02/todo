package cmd

import (
	"github.com/fatih/color"
	"github.com/lllllan02/todo/utils"
)

func printE(message string) {
	color.Red("%s  %s", utils.CrossMark, message)
}

func printS(message string) {
	color.Green("%s  %s", utils.CheckMark, message)
}
