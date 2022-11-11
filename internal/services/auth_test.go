package services_test

import (
	"net/url"
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

func (s *AuthServiceTest) TestGenerateRedirectURLSuccess() {
	service := "https://www.google.com"
	ticket := "this-is-a-ticket"
	result, err := s.authService.GeneratePath(service, ticket)
	s.NoError(err)

	resultURL, err := url.Parse(result)
	s.NoError(err)
	expected, err := url.Parse("https://www.google.com?ticket=this-is-a-ticket")
	s.NoError(err)

	s.Equal(resultURL.Host, expected.Host)
	s.Equal(resultURL.Path, expected.Path)
	s.Equal(resultURL.RawQuery, expected.RawQuery)
	s.Equal(resultURL.Scheme, expected.Scheme)
}

func (s *AuthServiceTest) TestGenerateRedirectURLSuccessWithConservedURL() {
	service := "https://www.google.com/?redirect=%2Fhome&test=1"
	ticket := "this-is-a-ticket"
	result, err := s.authService.GeneratePath(service, ticket)
	s.NoError(err)

	resultURL, err := url.Parse(result)
	s.NoError(err)
	expected, err := url.Parse("https://www.google.com/?redirect=%2Fhome&test=1&ticket=this-is-a-ticket")
	s.NoError(err)

	s.Equal(resultURL.Host, expected.Host)
	s.Equal(resultURL.Path, expected.Path)
	s.Equal(resultURL.RawQuery, expected.RawQuery)
	s.Equal(resultURL.Scheme, expected.Scheme)
}

func (s *AuthServiceTest) TestGenerateRedirectURLFailWithCantPraseURLService() {
	service := "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require"
	ticket := "this-is-a-ticket"
	result, err := s.authService.GeneratePath(service, ticket)
	s.Error(err)
	s.Empty(result)
}
