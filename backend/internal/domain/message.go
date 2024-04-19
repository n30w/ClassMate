package domain

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/n30w/Darkspace/internal/models"
)

// announcement and discussion services
type MessageStore interface {
	InsertMessage(m *models.Message, courseid models.CourseId) error
	GetMessageById(messageid models.MessageId) (*models.Message, error)
	DeleteMessage(m *models.Message) error
	ChangeMessage(m *models.Message) (*models.Message, error)
}

type MessageService struct {
	store MessageStore
}

func NewMessageService(m MessageStore) *MessageService { return &MessageService{store: m} }

func (ms *MessageService) ValidateID(id models.MessageId) bool {
	return true
}

func (ms *MessageService) CreateMessage(m *models.Message, courseid models.CourseId) (*models.Message, error) {
	newUUID := uuid.New()
	m.ID = models.MessageId(newUUID)

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

func (ms *MessageService) UpdateMessage(messageid models.MessageId, action string, updatedField string) (*models.Message, error) {
	if !ms.ValidateID(messageid) {
		return nil, fmt.Errorf("invalid message ID: %s", messageid)
	}

	msg, err := ms.store.GetMessageById(messageid)
	if err != nil {
		return nil, err
	}
	if action == "title" {
		msg.Post.Title = updatedField
		msg, err = ms.store.ChangeMessage(msg)
		if err != nil {
			return nil, err
		}
	} else if action == "body" {
		msg.Post.Description = updatedField
		msg, err = ms.store.ChangeMessage(msg)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("%s is an invalid action", action)
	}
	return msg, nil
}

func (ms *MessageService) DeleteMessage(messageid models.MessageId) error {
	if !ms.ValidateID(messageid) {
		return fmt.Errorf("invalid message ID: %s", messageid)
	}

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
