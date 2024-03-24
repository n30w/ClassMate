package domain

// announcement and discussion services
type MessageStore interface {
	InsertMessage()
	GetMessageById()
	DeleteMessage()
	ChangeMessage()
}
