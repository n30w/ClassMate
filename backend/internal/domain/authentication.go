package domain

import (
	"time"

	"github.com/n30w/Darkspace/internal/models"
)

type AuthenticationStore interface {
	InsertToken(t *models.Token) error
	DeleteTokenFrom(netId, scope string) error
	GetNetIdFromToken(token string) (string, error)
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
	err = as.store.InsertToken(token)
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (as *AuthenticationService) GetNetIdFromToken(token string) (string, error) {
	netid, err := as.store.GetNetIdFromToken(token)
	if err != nil {
		return "", err
	}
	return netid, nil
}
