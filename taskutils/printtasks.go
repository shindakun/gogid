package taskutils

import (
	"fmt"
	"gogid/model"
	"os"
	"text/template"
	"time"

	"github.com/fatih/color"
)

// Setup for coloring within the output template
var (
	red           = color.New(color.FgRed).SprintFunc()
	green         = color.New(color.FgHiGreen).SprintFunc()
	boldcyan      = color.New(color.FgCyan, color.Bold).SprintFunc()
	boldred       = color.New(color.FgRed, color.Bold).SprintFunc()
	boldblue      = color.New(color.FgHiBlue, color.Bold).SprintFunc()
	boldyellow    = color.New(color.FgYellow, color.Bold).SprintFunc()
	boldwhite     = color.New(color.FgHiWhite, color.Bold).SprintFunc()
	boldgreen     = color.New(color.FgHiGreen, color.Bold).SprintFunc()
	italicmagenta = color.New(color.FgHiMagenta, color.Italic).SprintFunc()
	italicblue    = color.New(color.FgBlue, color.Italic).SprintFunc()
)

// funcMap contains a map of functions to apply to the template output.
var funcMap = template.FuncMap{
	"id":         red,
	"task":       boldcyan,
	"complete":   italicblue,
	"nextaction": boldyellow,
	"time":       time.Unix,
}

const outputTemplate = `{{id .ID }} | {{task .Task}} | {{time .Created 0}} | {{time .Updated 0 }} | {{complete .Complete}} | {{nextaction .NextAction}} | {{.Notes}}`

// TODO: Should probably combine these all together and use switch from output type.

// PrintAllTasks prints out the entire task list.
func PrintAllTasks(taskList *model.TaskList) {
	tmpl, err := template.New("tasks").Funcs(funcMap).Parse(outputTemplate)
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

// BoolPrintTasks prints out the task list based on if it's completed or not.
func BoolPrintTasks(taskList *model.TaskList, complete bool) {
	tmpl, err := template.New("tasks").Funcs(funcMap).Parse(outputTemplate)
	if err != nil {
		panic(err)
	}

	for c := 0; c < len(taskList.Task); c++ {
		if taskList.Task[c].Complete == complete {
			err = tmpl.Execute(os.Stdout, taskList.Task[c])
			if err != nil {
				panic(err)
			}
			fmt.Printf("\n\n")
		}
	}
}

// PrintTask prints out the specificed task.
func PrintTask(taskList *model.TaskList, task int64) {
	if int64(len(taskList.Task)) > task && int64(len(taskList.Task)) >= 0 {
		tmpl, err := template.New("task").Funcs(funcMap).Parse(outputTemplate)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(os.Stdout, taskList.Task[task])
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n\n")
	} else {
		fmt.Printf("Task #%d does not exist.", task)
	}
}
