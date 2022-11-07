package services

import (
	"github.com/google/wire"
)

var Set = wire.Bind(new(IAuthService), new(*AuthService))

//go:generate mockgen -source=./auth.go --package=services -destination=./mocks/auth.go
type IAuthService interface {
	GetLogin(service string) error
	PostLogin(uid string) string
	ServiceValidation(ticket string) string
}

type AuthService struct {
	services map[string]int
}

func NewAuthService() *AuthService {
	return &AuthService{services: make(map[string]int)}
}

func (h *AuthService) GetLogin(service string) error {
	h.services[service] += 1
	return nil
}

func (h *AuthService) PostLogin(uid string) string {
	return "Hello"
}

func (h *AuthService) ServiceValidation(ticket string) string {
	return "Hello"
}
