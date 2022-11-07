package models

type GetLoginRequest struct {
	Service string `form:"service" binding:"required"`
}

type PostLoginRequest struct {
	Service   string `form:"service" binding:"required"`
	UID       string `form:"uid" binding:"required"`
	FirstName string `form:"fname" binding:"required"`
	LastName  string `form:"lname" binding:"required"`
}

type UserResponse struct {
	UID         string
	Username    string
	GECOS       string
	Disable     bool
	Roles       []string
	FirstName   string
	FirstNameTH string
	LastName    string
	LastNameTH  string
	OUID        string
	Email       string
}
