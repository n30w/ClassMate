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
	FileName           string            `json:"name"`
	AttributionsByType map[string]string `json:"attributions_by_type"`
	FileType           int               `json:"file_type"`
}
