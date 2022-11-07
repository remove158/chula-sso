package services_test

import (
	"testing"

	"github.com/remove158/chula-sso/internal/services"
	"github.com/stretchr/testify/suite"
)

type AuthServiceTest struct {
	suite.Suite
	authService *services.AuthService
}

func TestAuthService(t *testing.T) {
	authService := services.NewAuthService()
	suite.Run(t, &AuthServiceTest{authService: authService})
}

func (s *AuthServiceTest) TestGetLoginSuccess() {
	s.Equal(nil, s.authService.GetLogin("test"))
}

func (s *AuthServiceTest) TestPostLoginSuccess() {
	s.Equal("Hello", s.authService.PostLogin("test"))
}

func (s *AuthServiceTest) TestServiceValidationSuccess() {
	s.Equal("Hello", s.authService.ServiceValidation("test"))
}
