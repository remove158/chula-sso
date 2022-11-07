package services

import (
	"github.com/google/wire"
	"github.com/remove158/chula-sso/cmd/models"
)

var Set = wire.Bind(new(IAuthService), new(*AuthService))

//go:generate mockgen -source=./auth.go --package=services -destination=./mocks/auth.go
type IAuthService interface {
	UpdateUser(service string) error
	PostLogin(user models.UserResponse) string
	ServiceValidation(ticket string) string
}

type AuthService struct {
	services map[string]int
}

func NewAuthService() *AuthService {
	return &AuthService{services: make(map[string]int)}
}

func (h *AuthService) UpdateUser(service string) error {
	h.services[service] += 1
	return nil
}

func (h *AuthService) PostLogin(user models.UserResponse) string {
	return "Hello"
}

func (h *AuthService) ServiceValidation(ticket string) string {
	return "Hello"
}
