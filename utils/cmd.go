package utils

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	var cmd *exec.Cmd
	switch os := os.Getenv("GOOS"); os {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
