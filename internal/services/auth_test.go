package services_test

import (
	"testing"

	"github.com/remove158/chula-sso/cmd/models"
	"github.com/remove158/chula-sso/internal/services"
	"github.com/stretchr/testify/suite"
)

type AuthServiceTest struct {
	suite.Suite
	authService *services.AuthService
	test_user   models.UserResponse
}

func TestAuthService(t *testing.T) {
	authService := services.NewAuthService()
	test_user := models.UserResponse{
		UID:         "test",
		Username:    "test",
		GECOS:       "test",
		Disable:     false,
		Roles:       []string{"test@test.com"},
		FirstName:   "test",
		FirstNameTH: "test",
		LastName:    "test",
		LastNameTH:  "test",
		OUID:        "test",
		Email:       "test",
	}

	suite.Run(t, &AuthServiceTest{
		authService: authService,
		test_user:   test_user,
	})
}

func (s *AuthServiceTest) TestPostLoginSuccess() {

	ticket := s.authService.PostLogin(s.test_user)
	s.NotEmpty(ticket)
}

func (s *AuthServiceTest) TestServiceValidationSuccess() {
	ticket := s.authService.PostLogin(s.test_user)
	request := models.ServiceValidateRequest{
		Ticket:       ticket,
		DeeAppID:     "test",
		DeeAppSecret: "test",
	}
	result, err := s.authService.ServiceValidation(request)
	s.Equal(result.UID, s.test_user.UID)
	s.NoError(err)
}

func (s *AuthServiceTest) TestServiceValidationFailOnDeeAppIDInValid() {
	ticket := s.authService.PostLogin(s.test_user)
	request := models.ServiceValidateRequest{
		Ticket:       ticket,
		DeeAppID:     "",
		DeeAppSecret: "test",
	}
	result, err := s.authService.ServiceValidation(request)
	s.Error(err)
	s.Empty(result)
}

func (s *AuthServiceTest) TestServiceValidationFailOnDeeAppSecretInValid() {
	ticket := s.authService.PostLogin(s.test_user)
	request := models.ServiceValidateRequest{
		Ticket:       ticket,
		DeeAppID:     "test",
		DeeAppSecret: "",
	}
	result, err := s.authService.ServiceValidation(request)
	s.Error(err)
	s.Empty(result)
}

func (s *AuthServiceTest) TestServiceValidationFailOnTicketInvalid() {
	ticket := "invalid-" + s.authService.PostLogin(s.test_user)
	request := models.ServiceValidateRequest{
		Ticket:       ticket,
		DeeAppID:     "test",
		DeeAppSecret: "test",
	}
	result, err := s.authService.ServiceValidation(request)
	s.Error(err)
	s.Empty(result)
}
