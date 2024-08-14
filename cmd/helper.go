package cmd

import (
	"github.com/fatih/color"
)

const (
	CheckMark = "\u2714"
	CrossMark = "\u2716"
)

func printE(message string) {
	color.Red("%s  %s", CrossMark, message)
}

func printS(message string) {
	color.Green("%s  %s", CheckMark, message)
}
