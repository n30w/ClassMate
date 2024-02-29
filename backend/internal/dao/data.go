package dao

// UUIDs are only generated upon successful validation.

type Database interface {
	Insert() error
	Read() error
	Update() error
	Delete() error
}

type SQLAdapter struct {
	url      string
	password string
}

func NewSQLAdapter(url, password string) *SQLAdapter {
	return &SQLAdapter{
		url:      url,
		password: password,
	}
}
