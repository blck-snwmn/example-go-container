// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"database/sql"
)

type User struct {
	ID   string
	Name string
	Bio  sql.NullString
}
