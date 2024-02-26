package models

import "time"

type filetype int

const (
	JPG filetype = iota
	PNG
	PDF
	M4A
	MP3
	NULL
)

type Media struct {
	Name               string
	Uuid               string
	DateUploaded       time.Time
	CourseAttributions []Course
	FileType           string
}
