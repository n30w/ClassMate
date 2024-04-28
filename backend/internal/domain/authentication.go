package domain

import (
	"github.com/n30w/Darkspace/internal/models"
	"time"
)

type AuthenticationStore interface {
	InsertToken(t models.Token) error
	DeleteTokenFrom(netId, scope string) error
}

type AuthenticationService struct{ store AuthenticationStore }

func NewAuthenticationService(as AuthenticationStore) *AuthenticationService {
	return &AuthenticationService{store: as}
}

func (as *AuthenticationService) NewToken(netId string) (*models.Token, error) {
	token, err := models.GenerateToken(netId, 24*time.Hour, "authentication")
	if err != nil {
		return nil, err
	}

	return token, nil
}
