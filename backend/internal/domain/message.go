package domain

import "github.com/n30w/Darkspace/internal/models"

// announcement and discussion services
type MessageStore interface {
	InsertMessage(m *models.Message) error
	GetMessageById(messageid int64) (*models.Message, error)
	DeleteMessage(m *models.Message) error
	ChangeMessage(m *models.Message, msg string) (*models.Message, error)
}

type MessageService struct {
	store MessageStore
}

func NewMessageService(m MessageStore) *MessageService { return &MessageService{store: m} }

func (ms *MessageService) CreateMessage(m *models.Message) (*models.Message, error) {
	// create id here
	err := ms.store.InsertMessage(m)
	if err != nil {
		return nil, err
	}
	msg, err := ms.store.GetMessageById(int64(m.ID))
	if err != nil {
		return nil, err
	}
	return msg, nil
}
