package migrations

import "fmt"

type Task struct {
    TaskID         int
    TaskName       string
    TaskDescription string
    SolutionCode   string
    TestCases      []string
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks() ([]Task, error) {
    var tasks []Task
    rows, err := DB.Query("SELECT * FROM tasks")
    fmt.Println(rows)
    return tasks, err
}
