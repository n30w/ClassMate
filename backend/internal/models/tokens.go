package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"time"
)

const (
	ScopeActivation     = "activation"
	ScopeAuthentication = "authentication"
)

// Token is a stateful authentication tool to validate a user's identity.
// It implements the Credential interface.
type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    int64     `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}

// generateToken generates a new user token.
// This token will be stored in the database.
func generateToken(userID int64, ttl time.Duration, scope string) (
	*Token,
	error,
) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return token, nil
}

func (t Token) String() string {
	return t.Plaintext
}

func (t Token) Valid() error {
	if t.Plaintext == "" {
		return errors.New("token must be provided")
	}

	if len(t.Plaintext) != 26 {
		return errors.New("token must be 26 bytes long")
	}

	return nil
}
