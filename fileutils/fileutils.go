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
// It will create the directory and a blank tasks.json file if needed.
func DataDirCheck() string {
	home := os.Getenv("HOME")
	if home == "" {
		panic("can not get home dir")
	}
	path := filepath.Join(home, ".gogid")
	_, err := os.Stat(path)

	// directory does not exist we should create it.
	if err != nil {
		err = os.Mkdir(path, 0755)
		if err != nil {
			panic(err)
		}
	}

	tasksFile := filepath.Join(path, "tasks.json")

	// Check to see if tasks.json already exists on disk, create an empty one if
	// needed.
	if _, err := os.Stat(tasksFile); err != nil {

		// Write opening and closing JSON so next load doesn't result in crash.
		empty := []byte("{}")
		err = ioutil.WriteFile(tasksFile, empty, 0644)
		if err != nil {
			fmt.Printf("err %v", err)
		}
	}

	_, err = os.OpenFile(tasksFile, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	return tasksFile
}

// LoadTasks reads the entire tasks JSON file and
// unmarshals it and returns it.
func LoadTasks(path string) model.TaskList {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var taskList model.TaskList
	err = json.Unmarshal(file, &taskList)
	if err != nil {
		panic(err)
	}
	return taskList
}

// WriteTasks writes JSON back to disk.
func WriteTasks(path string, taskList *model.TaskList) {
	tList, err := json.Marshal(taskList)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	err = ioutil.WriteFile(path, tList, 0644)
	if err != nil {
		fmt.Printf("err %v", err)
	}
}
