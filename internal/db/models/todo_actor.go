package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TodoActor struct {
	Id        uuid.UUID    `json:"id"`
	TodoId    uuid.UUID    `json:"todo_id"`
	UserId    uuid.UUID    `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

var TodoActorColumnName = struct {
	Id        string
	TodoId    string
	UserId    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	"id",
	"todo_id",
	"user_id",
	"created_at",
	"updated_at",
	"deleted_at",
}

var TodoActorColumnList = []string{
	"id",
	"todo_id",
	"user_id",
	"created_at",
	"updated_at",
	"deleted_at",
}
