package models

type FileType int

const (
	JPG FileType = iota
	PNG
	PDF
	M4A
	MP3
	TXT
	XLSX
	NULL
)

func (f FileType) String() string {
	switch f {
	case JPG:
		return "jpg"
	case PNG:
		return "png"
	case PDF:
		return "pdf"
	case M4A:
		return "m4a"
	case MP3:
		return "mp3"
	case TXT:
		return "txt"
	case XLSX:
		return "xlsx"
	case NULL:
		return ""
	}
	return ""
}

type Media struct {
	Entity
	FileName           string            `json:"name"`
	AttributionsByType map[string]string `json:"attributions_by_type"`
	FileType           FileType          `json:"file_type"`
	FilePath           string            `json:"file_path"`
}

func NewMedia(fileName string, fileType FileType) *Media {
	return &Media{
		Entity:             Entity{},
		FileName:           fileName,
		AttributionsByType: nil,
		FileType:           fileType,
		FilePath:           "",
	}
}

const (
	DefaultImageId = "default_image"
)