package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TeamMembers struct {
	Id        uuid.UUID    `json:"id"`
	TeamId    uuid.UUID    `json:"team_id"`
	UserId    uuid.UUID    `json:"user_id"`
	Role      string       `json:"role"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

var TeamMembersColumnName = struct {
	Id        string
	TeamId    string
	UserId    string
	Role      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	"id",
	"team_id",
	"user_id",
	"role",
	"created_at",
	"updated_at",
	"deleted_at",
}

var TeamMembersColumnList = []string{
	"id",
	"team_id",
	"user_id",
	"role",
	"created_at",
	"updated_at",
	"deleted_at",
}
