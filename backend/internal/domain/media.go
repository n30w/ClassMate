package domain

// announcement and discussion services
type MediaStore interface {
	InsertMediaReference()
	GetMediaReferenceById()
}

type MediaService struct {
	store MediaStore
}

func NewMediaService(m MessageStore) *MediaService { return &MediaService{store: med} }
