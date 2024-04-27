package models

import (
	"database/sql"
)

// Entity defines a database object, in other words,
// an entity. This is a fundamental database object. I found this scheme from:
// https://github.com/g8rswimmer/go-data-access-example/blob/master/pkg/model/entity.go
type Entity struct {
	ID        string       `json:"id"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
