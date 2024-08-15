/*
Copyright © 2024 lllllan
*/
package main

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/lllllan02/todo/cmd"
	"github.com/lllllan02/todo/utils"
)

func main() {
	cmd.Execute([]string{"help"})

	rl, err := readline.New("todo> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Println("输入有误，请重新输入。")
			continue
		}

		if line == "exit" {
			fmt.Println("Bye")
			break
		} else if line == "clear" {
			utils.ClearScreen()
			continue
		}

		cmd.Execute(strings.Split(line, " "))
		fmt.Println()
	}
}
