package services

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
	"github.com/remove158/chula-sso/cmd/models"
)

var Set = wire.Bind(new(IAuthService), new(*AuthService))

type AppSecret struct {
	DeeAppSecret string `envconfig:"DEE_APP_SECRET" default:"test"`
	DeeAppId     string `envconfig:"DEE_APP_ID" default:"test"`
}

//go:generate mockgen -source=./auth.go --package=services -destination=./mocks/auth.go
type IAuthService interface {
	PostLogin(user models.UserResponse) string
	ServiceValidation(models.ServiceValidateRequest) (models.UserResponse, error)
	GeneratePath(service string, ticket string) (result string, err error)
}

type AuthService struct {
	services map[string]int
	cache    map[string]models.UserResponse
	config   AppSecret
}

func NewAuthService() *AuthService {
	var config AppSecret
	envconfig.Process("", &config)
	return &AuthService{
		services: make(map[string]int),
		config:   config,
		cache:    make(map[string]models.UserResponse),
	}
}

func (h *AuthService) PostLogin(user models.UserResponse) string {
	ticket := uuid.New().String()
	h.cache[ticket] = user
	return ticket
}

func (h *AuthService) ServiceValidation(request models.ServiceValidateRequest) (models.UserResponse, error) {
	if request.DeeAppID != h.config.DeeAppId || request.DeeAppSecret != h.config.DeeAppSecret {
		return models.UserResponse{}, fmt.Errorf("invalid app secret")
	}
	user, ok := h.cache[request.Ticket]
	if !ok {
		return models.UserResponse{}, fmt.Errorf("ticket not found")
	}
	return user, nil
}

func (h *AuthService) GeneratePath(service string, ticket string) (result string, err error) {
	var u *url.URL

	if u, err = url.Parse(service); err != nil {
		return
	}

	values := u.Query()
	values.Add("ticket", ticket)

	u.RawQuery = values.Encode()
	result = u.String()

	return
}
