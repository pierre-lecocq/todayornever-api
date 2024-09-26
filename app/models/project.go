// File: project.go
// Creation: Wed Sep 25 10:38:22 2024
// Time-stamp: <2024-09-26 14:00:57>
// Copyright (C): 2024 Pierre Lecocq

package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Project struct {
	ID          int64      `sql:"id" json:"id"`
	UserID      int64      `sql:"user_id" json:"user_id"`
	Name        string     `sql:"name" json:"name"`
	Description string     `sql:"description" json:"description"`
	CreatedAt   *time.Time `sql:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `sql:"updated_at" json:"updated_at"`
	Position    int64      `sql:"position" json:"position"`
}

func ListProjects(db *sql.DB, userID int64) ([]Project, error) {
	var projects []Project

	rows, err := db.Query(`
      SELECT id, user_id, name, COALESCE(description, ''), created_at, updated_at, position
        FROM project
        WHERE user_id = ?
        ORDER by position asc`,
		userID,
	)

	defer rows.Close()

	if err != nil {
		return projects, err
	}

	err = rows.Err()

	if err != nil {
		return projects, err
	}

	for rows.Next() {
		p := Project{}

		err = rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Name,
			&p.Description,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.Position,
		)

		if err != nil {
			// Use logger here
			panic(err)
			continue
		}

		projects = append(projects, p)
	}

	err = rows.Err()

	if err != nil {
		return projects, err
	}

	return projects, nil
}

func FetchProject(db *sql.DB, userID int64, ID int64) (Project, error) {
	p := Project{}

	err := db.QueryRow("SELECT id, user_id, name, COALESCE(description, ''), created_at, updated_at, position FROM project WHERE id = ? AND user_id = ?",
		ID,
		userID,
	).Scan(
		&p.ID,
		&p.UserID,
		&p.Name,
		&p.Description,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Position,
	)

	if err != nil {
		return p, err
	}

	return p, nil
}

func CreateProject(db *sql.DB, userID int64, p Project) (Project, error) {
	p.UserID = userID

	stmt, err := db.Prepare(`INSERT INTO project (user_id, name, description, position)
      SELECT ?, ?, NULL, COALESCE(MAX(position), 0) + 1 FROM project WHERE user_id = ?
      RETURNING id, user_id, name, COALESCE(description, ''), created_at, updated_at, position`)

	if err != nil {
		return p, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		p.UserID,
		p.Name,
		p.UserID,
	).Scan(
		&p.ID,
		&p.UserID,
		&p.Name,
		&p.Description,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Position,
	)

	if err != nil {
		return p, err
	}

	return p, nil
}

func UpdateProject(db *sql.DB, userID int64, ID int64, p Project) (Project, error) {
	pdb, err := FetchProject(db, userID, ID)

	if err != nil {
		return p, err
	}

	// @TODO clean that mess
	var cols []string
	var clauses []string
	var args []interface{}
	var vals []interface{}

	if p.Name != "" && p.Name != pdb.Name {
		cols = append(cols, "name")
		clauses = append(clauses, "name = ?")
		args = append(args, p.Name)
		vals = append(vals, &p.Name)
	}

	if len(args) == 0 {
		return p, fmt.Errorf("No data to update")
	}

	args = append(args, ID)
	args = append(args, userID)

	stmt, err := db.Prepare(
		fmt.Sprintf(
			"UPDATE project SET %s WHERE id = ? AND user_id = ? RETURNING %s",
			strings.Join(clauses, ", "),
			strings.Join(cols, ", "),
		),
	)

	defer stmt.Close()

	if err != nil {
		return p, err
	}

	err = stmt.QueryRow(args...).Scan(vals...)

	if err != nil {
		return p, err
	}

	return p, nil
}

func DeleteProject(db *sql.DB, userID int64, ID int64) (int64, error) {
	_, err := DeleteTasksForProject(db, userID, ID)

	if err != nil {
		return 0, err
	}

	stmt, err := db.Prepare("DELETE FROM project WHERE id = ? AND user_id = ?")

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
