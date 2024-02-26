package models

import "time"

type Media struct {
	name               string
	uuid               string
	dateUploaded       time.Time
	courseAttributions []Course
	fileType           string
}
