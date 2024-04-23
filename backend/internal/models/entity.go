package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Entity defines a database object, in other words,
// an entity. This is a fundamental database object. I found this scheme from:
// https://github.com/g8rswimmer/go-data-access-example/blob/master/pkg/model/entity.go
type Entity struct {
	ID        string       `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

// ID describes an identifier. This identifier is either a NetID
// or a Net ID, which would not have a UUID. ID implements stringer
// interface.
type ID struct {
	UUID       *uuid.UUID
	Serialized string
}

func (i ID) String() string {
	if i.UUID != nil {
		return i.UUID.String()
	}
	return i.Serialized
}

func (i ID) CreateUUID() {
	*i.UUID = uuid.New()
}
