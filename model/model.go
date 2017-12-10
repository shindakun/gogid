package model

// TaskList is a struct with an array of Task structs.
type TaskList struct {
	Task []Task
}

// Task is a struct that corrisponds to a single task.
type Task struct {
	ID         int64    `json:"id"`
	Task       string   `json:"task"`
	Created    int64    `json:"created"`
	Updated    int64    `json:"updated"`
	Complete   bool     `json:"complete"`
	NextAction string   `json:"nextaction"`
	Notes      []string `json:"notes"`
}
