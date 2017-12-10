package main

import (
	"flag"
	"gogid/fileutils"
	"gogid/taskutils"
	"os"

	"github.com/fatih/color"
)

// main command line interface for gogid.
func main() {

	// Should return path for tasks.json.
	tasksFile := fileutils.DataDirCheck()

	// TODO: Maybe move flags up to thier own function?
	printAllPtr := flag.Bool("printtasks", false, "Print entire task list to console.")
	printDonePtr := flag.Bool("printdone", false, "Print completed tasks to console.")
	printNotDonePtr := flag.Bool("printnotdone", false, "Print completed tasks to console.")
	printTaskPtr := flag.Int("print", -1, "Print a single task.")
	completeTaskPtr := flag.Int("complete", -1, "Mark a single task as completed.")
	notCompleteTaskPtr := flag.Int("notcomplete", -1, "Mark a single task as not completed.")

	idPtr := flag.Int("id", -1, "ID of task in question, used to add notes")

	var new string
	flag.StringVar(&new, "new", "", "Add a new task, enclose in quotes.")

	var newNote string
	flag.StringVar(&newNote, "newnote", "", "Add a new note to a task, enclose in quotes.")

	var newNextAction string
	flag.StringVar(&newNextAction, "nextaction", "", "Add a new next action, enclose in quotes.")

	flagNoColor := flag.Bool("no-color", false, "Disable color output")

	flag.Parse()

	if *flagNoColor {
		color.NoColor = true // disables colorized output
	}

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// TODO: Error handling?
	switch os.Args[1] {
	case "-printtasks":
		if *printAllPtr {
			taskList := fileutils.LoadTasks(tasksFile)
			taskutils.PrintAllTasks(&taskList)
		}
	case "-printdone":
		if *printDonePtr {
			taskList := fileutils.LoadTasks(tasksFile)
			taskutils.BoolPrintTasks(&taskList, true)
		}
	case "-printnotdone":
		if *printNotDonePtr {
			taskList := fileutils.LoadTasks(tasksFile)
			taskutils.BoolPrintTasks(&taskList, false)
		}
	case "-print":
		if *printTaskPtr > -1 {
			taskList := fileutils.LoadTasks(tasksFile)
			taskutils.PrintTask(&taskList, int32(*printTaskPtr))
		}
	case "-complete":
		if *completeTaskPtr > -1 {
			taskList := fileutils.LoadTasks(tasksFile)
			taskutils.CompleteTask(&taskList, int32(*completeTaskPtr))
			fileutils.WriteTasks(tasksFile, &taskList)
			taskutils.PrintTask(&taskList, int32(*completeTaskPtr))
		}
	case "-notcomplete":
		if *notCompleteTaskPtr > -1 {
			taskList := fileutils.LoadTasks(tasksFile)
			taskutils.InvertCompleteTask(&taskList, int32(*notCompleteTaskPtr))
			fileutils.WriteTasks(tasksFile, &taskList)
			taskutils.PrintTask(&taskList, int32(*notCompleteTaskPtr))
		}
	case "-new":
		if new != "" {
			taskList := fileutils.LoadTasks(tasksFile)
			taskutils.AddNewTask(&taskList, new)
			fileutils.WriteTasks(tasksFile, &taskList)
			taskutils.PrintTask(&taskList, int32(len(taskList.Task)-1))
		}
	case "-newnote":

		// Requires -id, should expand command to use subcommand but this works
		if newNote != "" && *idPtr > -1 {
			taskList := fileutils.LoadTasks(tasksFile)
			taskutils.AddNewNote(&taskList, int32(*idPtr), newNote)
			fileutils.WriteTasks(tasksFile, &taskList)
			taskutils.PrintTask(&taskList, int32(*idPtr))
		}
	case "-nextaction":

		// Requires -id, should expand command to use subcommand but this works
		if newNextAction != "" && *idPtr > -1 {
			taskList := fileutils.LoadTasks(tasksFile)
			taskutils.AddNextAction(&taskList, int32(*idPtr), newNextAction)
			fileutils.WriteTasks(tasksFile, &taskList)
			taskutils.PrintTask(&taskList, int32(*idPtr))
		}
	default:
		flag.PrintDefaults()
	}
}
