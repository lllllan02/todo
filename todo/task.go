package todo

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Task struct {
	Id        int
	Done      bool
	Title     string
	CreatedAt time.Time
}

type TaskTable []*Task

func (t TaskTable) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	for _, task := range t {
		table.Append([]string{
			fmt.Sprintf("%d", task.Id),
			task.Title,
			fmt.Sprintf("%t", task.Done),
			task.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	table.SetHeader([]string{"id", "title", "done", "created_at"})
	table.Render()
}
