package main

import (
	"github.com/lllllan02/todo/todo"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./gen",
		Mode:    gen.WithoutContext, // generate mode
	})
	g.ApplyBasic(todo.Task{})
	g.Execute()
}
