package migrations

// import "fmt"

type Task struct {
    TaskID         int
    TaskName       string
    TaskDescription string
    SolutionCode   string
    TestCases      interface{}
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks() ([]Task, error) {
    var tasks []Task
    rows, err := DB.Query("SELECT * FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var task Task
        err := rows.Scan(&task.TaskID, &task.TaskName, &task.TaskDescription, &task.SolutionCode, &task.TestCases)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}