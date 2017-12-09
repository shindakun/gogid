package main

import (
	"flag"
	"fmt"
	"gogid/fileutils"
	"gogid/taskutils"
)

func main() {
	// TODO: Maybe move flags up to thier own function?
	printAllPtr := flag.Bool("printtasks", false, "Print entire task list to console.")
	printTaskPtr := flag.Int("print", -1, "Print a single task.")
	completeTaskPtr := flag.Int("complete", -1, "Mark a single task as completed.")
	notCompleteTaskPtr := flag.Int("notcomplete", -1, "Mark a single task as not completed.")

	idPtr := flag.Int("id", -1, "ID of task in question, used to add notes")

	var new string
	flag.StringVar(&new, "new", "", "Add a new task, enclose in quotes.")

	var newNote string
	flag.StringVar(&newNote, "newnote", "", "Add a new note to a task, enclose in quotes.")

	flag.Parse()

	// TODO: Switch statement?

	if *printAllPtr {
		taskList := fileutils.LoadTasks("example.json")
		taskutils.PrintAllTasks(&taskList)
	} else if *printTaskPtr > -1 {
		taskList := fileutils.LoadTasks("example.json")
		taskutils.PrintTask(&taskList, int32(*printTaskPtr))
	} else if *completeTaskPtr > -1 {
		taskList := fileutils.LoadTasks("example.json")
		taskutils.CompleteTask(&taskList, int32(*completeTaskPtr))
		fileutils.WriteTasks("example.json", &taskList)
		taskutils.PrintTask(&taskList, int32(*completeTaskPtr))
	} else if *notCompleteTaskPtr > -1 {
		taskList := fileutils.LoadTasks("example.json")
		taskutils.InvertCompleteTask(&taskList, int32(*notCompleteTaskPtr))
		fileutils.WriteTasks("example.json", &taskList)
		taskutils.PrintTask(&taskList, int32(*notCompleteTaskPtr))
	} else if new != "" {
		taskList := fileutils.LoadTasks("example.json")
		taskutils.AddNewTask(&taskList, new)
		fileutils.WriteTasks("example.json", &taskList)
		taskutils.PrintTask(&taskList, int32(len(taskList.Task)-1))
	} else if newNote != "" && *idPtr > -1 {
		taskList := fileutils.LoadTasks("example.json")
		taskutils.AddNewNote(&taskList, int32(*idPtr), newNote)
		fileutils.WriteTasks("example.json", &taskList)
		taskutils.PrintTask(&taskList, int32(*idPtr))
	} else {
		fmt.Println("\n\n--------------------")
		fmt.Println("\n\nPrint useful stuff here.")
		fileutils.DataDirCheck()
		fmt.Println("\n\n--------------------")
	}
}
