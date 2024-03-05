package dao

import "database/sql"

type SQLAdapter struct {
	url      string
	password string
	db       *sql.DB
}

func (s SQLAdapter) Query(q Query) ([]Query, error) {
	rows, err := s.db.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

}
