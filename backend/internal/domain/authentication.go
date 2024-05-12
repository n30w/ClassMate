package domain

import (
	"time"

	"github.com/n30w/Darkspace/internal/models"
)

type AuthenticationStore interface {
	InsertToken(t *models.Token) error
	DeleteTokenFrom(netId, scope string) error
	GetNetIdFromHash(hash []byte) (string, error)
	GetTokenFromNetId(t *models.Token) (*models.Token, error)
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

func (as *AuthenticationService) RetrieveToken(netId string) (*models.Token, error) {
	t := &models.Token{
		NetID: netId,
	}

	token, err := as.store.GetTokenFromNetId(t)
	if err != nil {
		return nil, err
	}

	return token, err
}

func (as *AuthenticationService) GetNetIdFromToken(token string) (string, error) {
	hash := models.GenerateTokenHash(token)
	netid, err := as.store.GetNetIdFromHash(hash)
	if err != nil {
		return "", err
	}
	return netid, nil
}
