package fileutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gogid/model"
)

// DataDirCheck checks to see if .gogid exsits in the users home directory.
func DataDirCheck() {
	home := os.Getenv("HOME")
	if home == "" {
		panic("can not get home dir")
	}
	path := filepath.Join(home, ".gogid")
	fmt.Println(path)
	stat, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(stat)
}

// LoadTasks reads the entire tasks JSON file and
// unmarshals it and returns it.
func LoadTasks(path string) model.TaskList {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var taskList model.TaskList
	json.Unmarshal(file, &taskList)
	return taskList
}

// WriteTasks writes JSON back to disk
func WriteTasks(path string, taskList *model.TaskList) {
	tList, err := json.Marshal(taskList)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	err = ioutil.WriteFile(path, tList, 0644)
}
