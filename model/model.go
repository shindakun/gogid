package model

// TaskList is a struct with an array of Task structs.
type TaskList struct {
	Task []Task `json:"Task"`
}

// Task is a struct that corrisponds to a single task.
type Task struct {
	UUID       string   `json:"uuid"`
	ID         int64    `json:"id"`
	Task       string   `json:"task"`
	Created    int64    `json:"created"`
	Updated    int64    `json:"updated"`
	Complete   bool     `json:"complete"`
	NextAction string   `json:"nextaction"`
	Notes      []string `json:"notes"`
}

// TasksIdandUpdated a struct of []IDUpdated
type TasksIdandUpdated struct {
	IDUpdated []IDUpdated `json:"tasks"`
}

// IDUpdated a struct of UUID and Updated timestamp.
type IDUpdated struct {
	UUID    string `json:"uuid"`
	Updated int64  `json:"updated"`
}
