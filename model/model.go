package model

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int          ` db:"id" json:"id"`
	Description string       ` db:"description" json:"description"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at"`
	CompletedAt sql.NullTime `db:"completed_at" json:"completed_at"`
}

type UpdateListInput struct {
	Description *string `json:"description"`
}