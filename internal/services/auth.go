package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/remove158/chula-sso/cmd/models"
)

var Set = wire.Bind(new(IAuthService), new(*AuthService))

//go:generate mockgen -source=./auth.go --package=services -destination=./mocks/auth.go
type IAuthService interface {
	PostLogin(user models.UserResponse) string
	ServiceValidation(ticket string) (models.UserResponse, error)
}

type AuthService struct {
	services map[string]int
	cache    map[string]models.UserResponse
}

func NewAuthService() *AuthService {
	return &AuthService{
		services: make(map[string]int),
		cache:    make(map[string]models.UserResponse),
	}
}

func (h *AuthService) PostLogin(user models.UserResponse) string {
	ticket := uuid.New().String()
	h.cache[ticket] = user
	return ticket
}

func (h *AuthService) ServiceValidation(ticket string) (models.UserResponse, error) {
	user, ok := h.cache[ticket]
	if !ok {
		return models.UserResponse{}, fmt.Errorf("ticket not found")
	}
	return user, nil
}
