package taskutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nu7hatch/gouuid"

	"gogid/httputils"
	"gogid/model"
)

type TaskListToSend struct {
	IdUpdated []IdUpdated `json:"tasks"`
}

// Task is a struct that corrisponds to a single task.
type IdUpdated struct {
	UUID    string `json:"uuid"`
	Updated int64  `json:"updated"`
}

// SyncTasks attempts to sync with remote web server.
func SyncTasks(taskList *model.TaskList) {

	var tList TaskListToSend

	// Build our list of tasks and update timestampes.
	for c := 0; c < len(taskList.Task); c++ {
		var toSend IdUpdated
		toSend.UUID = taskList.Task[c].UUID
		toSend.Updated = taskList.Task[c].Updated
		tList.IdUpdated = append(tList.IdUpdated, toSend)
		// fmt.Printf("%s  %d", taskList.Task[c].UUID, taskList.Task[c].Updated)
		// fmt.Println("  " + string(httputils.HTTPRequest("GET", "http://localhost:3000/getbyuuid/"+taskList.Task[c].UUID, "", nil)))
		// task, _ := json.Marshal(taskList.Task[c])
		// fmt.Println("  " + string(httputils.HTTPRequest("POST", "http://localhost:3000/addtask/", "", bytes.NewReader(task))))

	}
	tasks, _ := json.Marshal(tList)

	// Send list of tasks to remote, recieve list of UUID's we need to send complete data for.
	resp := string(httputils.HTTPRequest("POST", "http://localhost:3000/addtask/", "", bytes.NewReader(tasks)))

	fmt.Println("  " + resp)

	//
}

// AddNewTask adds a new uncompleted task to the task list.
func AddNewTask(taskList *model.TaskList, task string) {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	var newTask model.Task
	newTask.ID = int64(len(taskList.Task))
	newTask.UUID = u.String()
	newTask.Task = task
	newTask.Created = int64(time.Now().Unix())
	newTask.Updated = int64(time.Now().Unix())
	newTask.Complete = false
	newTask.NextAction = ""
	newTask.Notes = make([]string, 0)
	taskList.Task = append(taskList.Task, newTask)
}

// AddNewNote adds a new note to a task.
func AddNewNote(taskList *model.TaskList, task int64, note string) {
	if int64(len(taskList.Task)) > task && int64(len(taskList.Task)) >= 0 {
		taskList.Task[task].Notes = append(taskList.Task[task].Notes, note)
		taskList.Task[task].Updated = int64(time.Now().Unix())
	}
}

// AddNextAction adds next action to task, no history, only one next action.
func AddNextAction(taskList *model.TaskList, task int64, nextAction string) {
	if int64(len(taskList.Task)) > task && int64(len(taskList.Task)) >= 0 {
		taskList.Task[task].NextAction = nextAction
		taskList.Task[task].Updated = int64(time.Now().Unix())
	}
}

// TODO: Combine these, or leave as on always sets true and one always sets
// false

// CompleteTask marks a task as completed and updates update timestamp.
func CompleteTask(taskList *model.TaskList, task int64) {
	if int64(len(taskList.Task)) > task && int64(len(taskList.Task)) >= 0 {
		taskList.Task[task].Complete = true
		taskList.Task[task].Updated = int64(time.Now().Unix())
	}
}

// InvertCompleteTask flips bool to oposite and updates update timestamp.
func InvertCompleteTask(taskList *model.TaskList, task int64) {
	if int64(len(taskList.Task)) > task && int64(len(taskList.Task)) >= 0 {
		taskList.Task[task].Complete = !taskList.Task[task].Complete
		taskList.Task[task].Updated = int64(time.Now().Unix())
	}
}
