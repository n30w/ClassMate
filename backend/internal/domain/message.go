package domain

import "github.com/n30w/Darkspace/internal/models"

// announcement and discussion services
type MessageStore interface {
	InsertMessage(m *models.Message) error
	GetMessageById(messageid int64) (*models.Message, error)
	DeleteMessage(m *models.Message) error
	ChangeMessage(m *models.Message) (*models.Message, error)
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

func (ms *MessageService) UpdateMessage(messageid int64, action string, updatedField string) (*models.Message, error) {
	// check for message
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
		return nil, error // need to format error
	}
	return msg, nil
}

func (ms *MessageService) DeleteMessage(msgid int64) error {

	msg, err := ms.store.GetMessageById(msgid)
	if err != nil {
		return err
	}
	err = ms.store.DeleteMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
