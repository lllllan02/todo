/*
Copyright © 2024 lllllan
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lllllan02/todo/cmd"
)

func main() {
	cmd.Execute([]string{"help"})
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\ntodo> ")
		if scanner.Scan() {
			command := scanner.Text()
			if err := scanner.Err(); err != nil {
				fmt.Println("输入有误，请重新输入。")
				continue
			}

			if command == "exit" {
				fmt.Println("Bye")
				break
			}

			cmd.Execute(strings.Split(command, " "))
		}
	}
}


