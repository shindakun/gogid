package taskutils

import (
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/shindakun/gogid/model"
)

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
	if task >= 0 && task <= GetNumberOfTasks(taskList) {
		taskList.Task[task].Notes = append(taskList.Task[task].Notes, note)
		taskList.Task[task].Updated = int64(time.Now().Unix())
	}
}

// AddNextAction adds next action to task, no history, only one next action.
func AddNextAction(taskList *model.TaskList, task int64, nextAction string) {
	if task >= 0 && task <= GetNumberOfTasks(taskList) {
		taskList.Task[task].NextAction = nextAction
		taskList.Task[task].Updated = int64(time.Now().Unix())
	}
}

// TODO: Combine these, or leave as on always sets true and one always sets
// false

// CompleteTask marks a task as completed and updates update timestamp.
func CompleteTask(taskList *model.TaskList, task int64) {
	if task >= 0 && task <= GetNumberOfTasks(taskList) {
		taskList.Task[task].Complete = true
		taskList.Task[task].Updated = int64(time.Now().Unix())
	}
}

// InvertCompleteTask flips bool to oposite and updates update timestamp.
func InvertCompleteTask(taskList *model.TaskList, task int64) {
	if task >= 0 && task <= GetNumberOfTasks(taskList) {
		if taskList.Task[task].Complete == true {
			// task already completed.
		}
		taskList.Task[task].Complete = !taskList.Task[task].Complete
		taskList.Task[task].Updated = int64(time.Now().Unix())
	}
}

// GetNumberOfTasks returns the length
func GetNumberOfTasks(taskList *model.TaskList) int64 {
	return int64(len(taskList.Task) - 1)
}
