package taskutils

import (
	"time"

	"gogid/model"
)

// AddNewTask adds a new uncompleted task to the task list.
func AddNewTask(taskList *model.TaskList, task string) {
	var newTask model.Task
	newTask.ID = int32(len(taskList.Task))
	newTask.Task = task
	newTask.Created = int32(time.Now().Unix())
	newTask.Updated = int32(time.Now().Unix())
	newTask.Complete = false
	newTask.NextAction = ""
	newTask.Notes = make([]string, 0)
	taskList.Task = append(taskList.Task, newTask)
}

// AddNewNote adds a new note to a task.
func AddNewNote(taskList *model.TaskList, task int32, note string) {
	taskList.Task[task].Notes = append(taskList.Task[task].Notes, note)
	taskList.Task[task].Updated = int32(time.Now().Unix())
}

// AddNextAction adds next action to task, no history, only one next action.
func AddNextAction(taskList *model.TaskList, task int32, nextAction string) {
	taskList.Task[task].NextAction = nextAction
	taskList.Task[task].Updated = int32(time.Now().Unix())
}

// TODO: Combine these, or leave as on always sets true and one always sets
// false

// CompleteTask marks a task as completed and updates update timestamp.
func CompleteTask(taskList *model.TaskList, task int32) {
	taskList.Task[task].Complete = true
	taskList.Task[task].Updated = int32(time.Now().Unix())
}

// InvertCompleteTask flips bool to oposite and updates update timestamp.
func InvertCompleteTask(taskList *model.TaskList, task int32) {
	taskList.Task[task].Complete = !taskList.Task[task].Complete
	taskList.Task[task].Updated = int32(time.Now().Unix())
}
