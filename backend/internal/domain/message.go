package domain

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/n30w/Darkspace/internal/models"
)

// announcement and discussion services
type MessageStore interface {
	InsertMessage(m *models.Message, courseid string) error
	GetMessageById(messageid string) (*models.Message, error)
	DeleteMessage(m *models.Message) error
	ChangeMessageTitle(m *models.Message) (*models.Message, error)
	ChangeMessageBody(m *models.Message) (*models.Message, error)
}

type MessageService struct {
	store MessageStore
}

func NewMessageService(m MessageStore) *MessageService { return &MessageService{store: m} }

func (ms *MessageService) CreateMessage(m *models.Message, courseid string) (*models.Message, error) {
	m.ID = uuid.New().String()
	err := ms.store.InsertMessage(m, courseid)
	if err != nil {
		return nil, err
	}
	msg, err := ms.store.GetMessageById(m.ID)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (ms *MessageService) UpdateMessage(messageid string, action string, updatedField string) (*models.Message, error) {
	msg, err := ms.store.GetMessageById(messageid)
	if err != nil {
		return nil, err
	}
	if action == "title" {
		msg.Post.Title = updatedField
		msg, err = ms.store.ChangeMessageTitle(msg)
		if err != nil {
			return nil, err
		}
	} else if action == "body" {
		msg.Post.Description = updatedField
		msg, err = ms.store.ChangeMessageBody(msg)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("%s is an invalid action", action)
	}
	return msg, nil
}

func (ms *MessageService) DeleteMessage(messageid string) error {

	msg, err := ms.store.GetMessageById(messageid)
	if err != nil {
		return err
	}
	err = ms.store.DeleteMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MessageService) ReadMessage(messageid string) (*models.Message, error) {
	msg, err := ms.store.GetMessageById(messageid)
	if err != nil {
		return nil, err
	}
	return msg, err
}
