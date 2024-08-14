package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/lllllan02/todo/gen"
	"github.com/lllllan02/todo/todo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	query  *gen.Query
	dbpath string
)

func init() {
	flag.StringVar(&dbpath, "db", "", "db file (default is $HOME/.todo.db)")
	flag.Parse()

	// Default to $HOME/.todo.db
	if dbpath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		dbpath = fmt.Sprintf("%s/.todo.db", homeDir)
	}

	// Create the db file
	file, err := os.OpenFile(dbpath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Open the db
	if db, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{}); err != nil {
		panic("failed to connect ~/.todo.db")
	}
	query = gen.Use(db)

	// Migrate the schema
	db.AutoMigrate(&todo.Task{})
}
