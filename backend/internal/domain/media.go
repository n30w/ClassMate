package domain

import "time"

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
	Uuid               string    `json:"uuid"`
	DateUploaded       time.Time `json:"date_uploaded"`
	CourseAttributions []Course  `json:"course_attributions"`
	FileType           string    `json:"file_type"`
}
