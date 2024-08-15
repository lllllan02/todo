package todo

import (
	"fmt"
	"os"
	"time"

	"github.com/lllllan02/todo/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Task struct {
	Id        int
	Done      bool
	Title     string
	StartTime time.Time // 开始时间
	DueDate   time.Time // 截止时间
	CreatedAt time.Time // 创建时间

	Sts string `gorm:"-" survey:"Start Time"`
	Dds string `gorm:"-" survey:"Due Date"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if t.Sts != "" {
		t.StartTime = cast.ToTimeInDefaultLocation(t.Sts, time.Local)
	}

	if t.Dds != "" {
		t.DueDate = cast.ToTimeInDefaultLocation(t.Dds, time.Local)
	}

	return nil
}

type TaskTable []*Task

func (t TaskTable) Render() {
	table := tablewriter.NewWriter(os.Stdout)
	for _, task := range t {
		var color []int
		if task.Done {
			color = []int{tablewriter.FgGreenColor}
		}

		table.Rich(
			[]string{lo.Ternary(task.Done, utils.CheckMark, utils.BlankSquareBox), fmt.Sprintf("%d", task.Id), task.Title, task.StartTime.Format(time.DateOnly), task.DueDate.Format(time.DateOnly), task.CreatedAt.Format(time.DateOnly)},
			[]tablewriter.Colors{color, color, color, color, color, color},
		)
	}

	table.SetHeader([]string{"", "id", "title", "start_time", "due_time", "created_at"})
	table.SetFooter([]string{"total", cast.ToString(len(t)), "", "", "", ""})
	table.SetBorder(false)
	table.Render()
}
