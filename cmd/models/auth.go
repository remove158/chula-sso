package models

type GetLoginRequest struct {
	Service string `form:"service" binding:"required"`
}
