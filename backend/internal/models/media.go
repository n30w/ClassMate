package models

import (
	"time"
)

type filetype int

const (
	JPG filetype = iota
	PNG
	PDF
	M4A
	MP3
	TXT
	NULL
)

type Media struct {
	Name               string    `json:"name"`
	MediaId            string    `json:"uuid"`
	DateUploaded       time.Time `json:"date_uploaded"`
	CourseAttributions []string  `json:"course_attributions"`
	FileType           int       `json:"file_type"`
}
