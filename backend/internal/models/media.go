package models

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
	Entity
	Name               string   `json:"name"`
	CourseAttributions []string `json:"course_attributions"`
	FileType           int      `json:"file_type"`
}
