package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/remove158/chula-sso/cmd/models"
	"github.com/remove158/chula-sso/internal/services"
)

var Set = wire.Bind(new(IAuthHandler), new(*AuthHandler))

//go:generate mockgen -source=./auth.go --package=handlers -destination=./mocks/auth.go
type IAuthHandler interface {
	GetLogin(ctx *gin.Context)
	PostLogin(ctx *gin.Context)
	ServiceValidation(ctx *gin.Context)
}

type AuthHandler struct {
	authService services.IAuthService
}

func NewAuthHandler(authService services.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) GetLogin(ctx *gin.Context) {
	var request models.GetLoginRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"service": request.Service,
	})
}

func (h *AuthHandler) PostLogin(ctx *gin.Context) {
	var request models.PostLoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := generateUserResponse(request)
	ticket := h.authService.PostLogin(user)
	ctx.Redirect(http.StatusTemporaryRedirect, generatePath(request.Service, ticket))
}

func generateUserResponse(request models.PostLoginRequest) models.UserResponse {
	username := fmt.Sprintf("%s-%s", request.FirstName, request.LastName)
	email := fmt.Sprintf("%s@student.chula.ac.th", request.UID)

	return models.UserResponse{
		UID:         request.UID,
		Username:    username,
		GECOS:       username,
		Disable:     false,
		Roles:       []string{request.Roles},
		FirstName:   request.FirstName,
		FirstNameTH: request.FirstName,
		LastName:    request.LastName,
		LastNameTH:  request.LastName,
		OUID:        request.UID,
		Email:       email,
	}
}

func generatePath(service string, ticket string) string {
	return fmt.Sprintf("%s?ticket=%s", service, ticket)
}

func (h *AuthHandler) ServiceValidation(ctx *gin.Context) {
	var request models.ServiceValidateRequest
	if err := ctx.ShouldBindHeader(&request); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	response, err := h.authService.ServiceValidation(request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}
