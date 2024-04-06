package domain

import (
	"github.com/n30w/Darkspace/internal/models"
)

// Announcement and Discussion services
type MessageStore interface {
	InsertMessage(m *models.Message) (*models.Message, error)
	GetMessageById(id string) (*models.Message, error)
	DeleteMessage(id string) error
	EditMessage(id string) (*models.Message, error)
}

type MessageService struct {
	store MessageStore
}

func NewMessageService(ms MessageStore) *MessageService {
	return &MessageService{store: ms}
}

func (ms *MessageService) CreateMessage(m *models.Message) (
	*models.Message,
	error,
) {
	return nil, nil
}
