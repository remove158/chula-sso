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
	_ = h.authService.GetLogin(request.Service)

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
	fmt.Printf("request: %v\n", request)
	_ = h.authService.PostLogin("id")
	ctx.Redirect(http.StatusTemporaryRedirect, generatePath(request.Service, request.UID))
}

func generatePath(service string, ticket string) string {
	return fmt.Sprintf("%s?ticket=%s", service, ticket)
}

func (h *AuthHandler) ServiceValidation(ctx *gin.Context) {
	ticket := ""
	response := h.authService.ServiceValidation(ticket)
	ctx.JSON(200, gin.H{
		"message": response,
	})
}
