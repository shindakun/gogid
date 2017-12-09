package taskutils

import (
	"fmt"
	"gogid/model"
	"os"
	"text/template"

	_ "github.com/fatih/color"
)

// PrintAllTasks prints out the entire task list.
func PrintAllTasks(taskList *model.TaskList) {
	// ID:18 Task:test task Created:1512845256 Updated:1512845256 Complete:false Notes:[]
	tmpl, err := template.New("tasks").Parse("{{.ID}} {{.Task}} {{.Created}}  {{.Updated}}  {{.Complete}}  {{.Notes}}")
	if err != nil {
		panic(err)
	}

	for c := 0; c < len(taskList.Task); c++ {
		err = tmpl.Execute(os.Stdout, taskList.Task[c])
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n\n")
	}
}

// PrintTask prints out the specificed task.
func PrintTask(taskList *model.TaskList, task int32) {
	// ID:18 Task:test task Created:1512845256 Updated:1512845256 Complete:false Notes:[]
	tmpl, err := template.New("task").Parse("{{.ID}} {{.Task}} {{.Created}}  {{.Updated}}  {{.Complete}}  {{.Notes}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, taskList.Task[task])
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\n")
}
