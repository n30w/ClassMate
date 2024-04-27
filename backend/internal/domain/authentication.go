package domain

type AuthenticationStore interface{}

type AuthenticationService struct{ store AuthenticationStore }

func NewAuthenticationService(as AuthenticationStore) *AuthenticationService {
	return &AuthenticationService{store: as}
}
