package models

import "time"

type project struct {
	name            string
	uuid            string
	deadline        time.Time
	mediaReferences []Media
	members         []User
	discussion      Discussion
}
