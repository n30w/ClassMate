package models

type FileType int

const (
	JPG FileType = iota
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
	FileType           FileType          `json:"file_type"`
}
