// File: task.go
// Creation: Thu Sep  5 08:29:59 2024
// Time-stamp: <2024-09-26 14:31:00>
// Copyright (C): 2024 Pierre Lecocq

package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Task struct {
	ID           int64      `sql:"id" json:"id"`
	UserID       int64      `sql:"user_id" json:"user_id"`
	ProjectID    int64      `sql:"project_id" json:"project_id"`
	ParentTaskID int64      `sql:"parent_task_id" json:"parent_task_id"`
	Title        string     `sql:"title" json:"title"`
	State        string     `sql:"state" json:"state"`
	DueAt        *time.Time `sql:"due_at" json:"due_at"`
	CreatedAt    *time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `sql:"updated_at" json:"updated_at"`
	Position     int64      `sql:"position" json:"position"`
	Overdue      bool       `sql:"-" json:"overdue"`
}

func ListTasks(db *sql.DB, userID int64, page int64, limit int64) ([]Task, error) {
	var tasks []Task

	rows, err := db.Query(`
      SELECT id, user_id, COALESCE(project_id, 0), COALESCE(parent_task_id, 0), title, state, due_at, created_at, updated_at, position
        FROM task
        WHERE user_id = ?
        ORDER by (case state when 'todo' then 0 when 'done' then 1 end), position asc
        LIMIT ? OFFSET ?`,
		userID,
		limit,
		(page-1)*limit,
	)

	defer rows.Close()

	if err != nil {
		return tasks, err
	}

	err = rows.Err()

	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		t := Task{}

		err = rows.Scan(
			&t.ID,
			&t.UserID,
			&t.ProjectID,
			&t.ParentTaskID,
			&t.Title,
			&t.State,
			&t.DueAt,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.Position,
		)

		if err != nil {
			// Use logger here
			continue
		}

		t.Overdue = t.DueAt.Before(time.Now())

		tasks = append(tasks, t)
	}

	err = rows.Err()

	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func CountTasks(db *sql.DB, userID int64) (int, error) {
	var count int

	err := db.QueryRow("SELECT count(1) FROM task WHERE user_id = ?", userID).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func FetchTask(db *sql.DB, userID int64, ID int64) (Task, error) {
	t := Task{}

	err := db.QueryRow(
		"SELECT id, user_id, COALESCE(project_id, 0), COALESCE(parent_task_id, 0), title, state, due_at, created_at, updated_at, position FROM task WHERE id = ? AND user_id = ?",
		ID,
		userID,
	).Scan(
		&t.ID,
		&t.UserID,
		&t.ProjectID,
		&t.ParentTaskID,
		&t.Title,
		&t.State,
		&t.DueAt,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.Position,
	)

	if err != nil {
		return t, err
	}

	t.Overdue = t.DueAt.Before(time.Now())

	return t, nil
}

func CreateTask(db *sql.DB, userID int64, t Task) (Task, error) {
	t.UserID = userID

	stmt, err := db.Prepare(`INSERT INTO task (user_id, project_id, parent_task_id, title, state, due_at, position)
      SELECT ?, NULL, NULL, ?, 'todo', datetime('now', '+1 day', 'start of day'), COALESCE(MAX(position), 0) + 1 FROM task WHERE user_id = ?
      RETURNING id, user_id, COALESCE(project_id, 0), COALESCE(parent_task_id, 0), title, state, due_at, created_at, updated_at, position`)

	if err != nil {
		return t, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		t.UserID,
		t.Title,
		t.UserID,
	).Scan(
		&t.ID,
		&t.UserID,
		&t.ProjectID,
		&t.ParentTaskID,
		&t.Title,
		&t.State,
		&t.DueAt,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.Position,
	)

	if err != nil {
		return t, err
	}

	t.Overdue = t.DueAt.Before(time.Now())

	return t, nil
}

func UpdateTask(db *sql.DB, userID int64, ID int64, t Task) (Task, error) {
	tdb, err := FetchTask(db, userID, ID)

	if err != nil {
		return t, err
	}

	// @TODO clean that mess
	var cols []string
	var clauses []string
	var args []interface{}
	var vals []interface{}

	if t.Title != "" && t.Title != tdb.Title {
		cols = append(cols, "title")
		clauses = append(clauses, "title = ?")
		args = append(args, t.Title)
		vals = append(vals, &t.Title)
	}

	if t.State != "" && t.State != tdb.State {
		cols = append(cols, "state")
		clauses = append(clauses, "state = ?")
		args = append(args, t.State)
		vals = append(vals, &t.State)
	}

	if t.DueAt != nil && t.DueAt != tdb.DueAt {
		cols = append(cols, "due_at")
		clauses = append(clauses, "due_at = ?")
		args = append(args, t.DueAt)
		vals = append(vals, &t.DueAt)
	}

	if len(args) == 0 {
		return t, fmt.Errorf("No data to update")
	}

	args = append(args, ID)
	args = append(args, userID)

	stmt, err := db.Prepare(
		fmt.Sprintf(
			"UPDATE task SET %s WHERE id = ? AND user_id = ? RETURNING %s",
			strings.Join(clauses, ", "),
			strings.Join(cols, ", "),
		),
	)

	defer stmt.Close()

	if err != nil {
		return t, err
	}

	err = stmt.QueryRow(args...).Scan(vals...)

	if err != nil {
		return t, err
	}

	return t, nil
}

func DeleteTask(db *sql.DB, userID int64, ID int64) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM task WHERE id = ? AND user_id = ?")

	defer stmt.Close()

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(ID, userID)

	if err != nil {
		return 0, err
	}

	nb, _ := res.RowsAffected()

	return nb, nil
}

func DeleteTasksForProject(db *sql.DB, userID int64, projectID int64) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM task WHERE project_id = ? AND user_id = ?")

	defer stmt.Close()

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(projectID, userID)

	if err != nil {
		return 0, err
	}

	nb, _ := res.RowsAffected()

	return nb, nil
}
