package dal

import (
	"errors"
)

var (
	ERR_RECORD_NOT_FOUND = errors.New("record not found")
	ERR_INVALID_BY       = errors.New("invalid get type received")
)
