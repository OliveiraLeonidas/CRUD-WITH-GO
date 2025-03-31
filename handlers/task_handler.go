package handlers

import "database/sql"

type TaskHandler struct {
	DB *sql.DB
}

// constructor de task handler
func newTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}
