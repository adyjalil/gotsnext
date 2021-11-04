package models

import (
	"database/sql"
	"time"
)

type User struct {
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

var UserColumnName = struct {
	Username  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	"username",
	"email",
	"password",
	"created_at",
	"updated_at",
	"deleted_at",
}

var UserColumnList = []string{
	"username",
	"email",
	"password",
	"created_at",
	"updated_at",
	"deleted_at",
}
