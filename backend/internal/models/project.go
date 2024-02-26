package models

import "time"

type project struct {
	Name            string
	Uuid            string
	Deadline        time.Time
	MediaReferences []Media
	Members         []User
	Discussion      Discussion
}
