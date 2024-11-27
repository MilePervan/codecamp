package infrastructure

import (
	"database/sql"
	"fmt"
	"time"
	"todo-cc/port"
)

type SqliteAdapter struct {
	dbClient *sql.DB
}

func NewPersistenceAdapter(dbClient *sql.DB) *SqliteAdapter {
	return &SqliteAdapter{
		dbClient: dbClient,
	}
}

func (a *SqliteAdapter) GetTask(id int) (*port.TaskDTO, error) {
	findTaskSqlStatement := `
  SELECT title, description, deadline, completed, deleted FROM task WHERE id = ?;
`
	statement, err := a.dbClient.Prepare(findTaskSqlStatement)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare query: %v", err.Error())
	}
	defer statement.Close()

	var TaskDTO port.TaskDTO
	err = statement.
		QueryRow(id).
		Scan(&TaskDTO.Title, &TaskDTO.Description, &TaskDTO.Deadline, &TaskDTO.Completed, &TaskDTO.Deleted)
	if err != nil {
		return nil, fmt.Errorf("unable to set ID into statement: %v", err.Error())
	}

	return &TaskDTO, nil
}

func (a *SqliteAdapter) NewTask(title, description string, deadline time.Time, completed bool) error {
	createTaskSql := `INSERT INTO task(title, description, deadline, completed) values(?, ?, ?, ?)`

	stmt, err := a.dbClient.Prepare(createTaskSql)
	if err != nil {
		return fmt.Errorf("unable to prepare insert statement: %v", err.Error())
	}

	_, err = stmt.Exec(title, description, deadline, completed)
	if err != nil {
		return fmt.Errorf("unable to execute insert statement: %v", err.Error())
	}
	return nil
}

// DeleteTask GetAllTasks CompleteTask TODO Think about parameters and return type(s)
func (a *SqliteAdapter) GetAllTasks() {
	func (a *SqliteAdapter) GetAllTasks() ([]port.TaskDTO, error) {
		getAllTasksSql := `SELECT id, title, description, deadline, completed, deleted FROM task`
	
		rows, err := a.dbClient.Query(getAllTasksSql)
		if err != nil {
			return nil, fmt.Errorf("unable to query tasks: %v", err.Error())
		}
		defer rows.Close()
	
		var tasks []port.TaskDTO
		for rows.Next() {
			var task port.TaskDTO
			err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Deadline, &task.Completed, &task.Deleted)
			if err != nil {
				return nil, fmt.Errorf("unable to scan row: %v", err.Error())
			}
			tasks = append(tasks, task)
		}
	
		if err = rows.Err(); err != nil {
			return nil, fmt.Errorf("error while iterating rows: %v", err.Error())
		}
	
		return tasks, nil
	}
}

func (a *SqliteAdapter) DeleteTask() {
}

func (a *SqliteAdapter) CompleteTask() {
}
