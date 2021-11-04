package psql

import (
	"context"
	"database/sql"
	"fmt"
	"gotsnext/internal/db/models"
	"strconv"
	"strings"
)

func AddUser(ctx context.Context, db *sql.DB, m *models.User) error {
	params := []string{}

	for i := 0; i < len(models.UserColumnList); i++ {
		is := strconv.Itoa(i + 1)
		params = append(params, "$"+is)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO users (%s) VALUES (%s)",
		strings.Join(models.UserColumnList, ","),
		strings.Join(params, ","),
	)

}
