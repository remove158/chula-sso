package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	if _, err := url.Parse(request.Service); err != nil {
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

	redirectURL, err := h.authService.GeneratePath(request.Service, ticket)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Redirect(http.StatusFound, redirectURL)
}

func generateUserResponse(request models.PostLoginRequest) models.UserResponse {
	uid := uuid.New().String()
	email := fmt.Sprintf("%s@student.chula.ac.th", request.UID)
	gecos := fmt.Sprintf("%s %s, ENG", request.FirstName, request.LastName)

	return models.UserResponse{
		UID:         uid,
		Username:    request.UID,
		GECOS:       gecos,
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
	ctx.JSON(http.StatusOK, response)
}
