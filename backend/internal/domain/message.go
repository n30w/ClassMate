package domain

import (
	"fmt"
	"time"

	"github.com/n30w/Darkspace/internal/models"
)

// announcement and discussion services
type MessageStore interface {
	InsertMessage(m *models.Message, courseid string) error
	GetMessageById(messageid string) (*models.Message, error)
	DeleteMessageByID(messageid string) error
	ChangeMessageTitle(m *models.Message) (*models.Message, error)
	ChangeMessageBody(m *models.Message) (*models.Message, error)
	GetMessagesByCourse(courseid string) ([]string, error)
}

type MessageService struct {
	store MessageStore
}

func NewMessageService(m MessageStore) *MessageService { return &MessageService{store: m} }

// CreateAnnouncement inserts an announcement into the database
// using method parameters.
func (ms *MessageService) CreateAnnouncement(
	title, description,
	owner, courseId string,
) (*models.Message, error) {
	msg := models.NewMessage(title, description, owner, true)

	msg.CreatedAt = time.Now()

	err := ms.store.InsertMessage(msg, courseId)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (ms *MessageService) CreateMessage(m *models.Message, courseid string) (*models.Message, error) {
	err := ms.store.InsertMessage(m, courseid)
	if err != nil {
		return nil, err
	}
	return m, nil
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

	err := ms.store.DeleteMessageByID(messageid)
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
func (ms *MessageService) RetrieveMessages(courseid string) ([]string, error) {

	msgids, err := ms.store.GetMessagesByCourse(courseid)
	if err != nil {
		return nil, err
	}
	return msgids, err
}
