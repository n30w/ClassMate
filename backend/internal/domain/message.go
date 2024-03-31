package domain

import (
	"github.com/n30w/Darkspace/internal/models"
)

// announcement and discussion services
type MessageStore interface {
	InsertMessage()
	GetMessageById()
	DeleteMessage()
	EditMessage()
}

type MessageService struct {
	store MessageStore
}

func NewMessageService(ms MessageStore) *MessageService {
	return &MessageService{store: ms}
}

func (ms *MessageService) CreateMessage(m *models.Message) (*models.Message, error) {
	return nil, nil
}
