package models

type GetLoginRequest struct {
	Service string `form:"service" binding:"required"`
}

type ServiceValidateRequest struct {
	Ticket       string `header:"DeeTicket" binding:"required"`
	DeeAppID     string `header:"DeeAppId" binding:"required"`
	DeeAppSecret string `header:"DeeAppSecret" binding:"required"`
}

type PostLoginRequest struct {
	Service   string `form:"service" binding:"required"`
	UID       string `form:"uid" binding:"required"`
	FirstName string `form:"fname" binding:"required"`
	LastName  string `form:"lname" binding:"required"`
	Roles     string `form:"roles" binding:"required"`
}

type UserResponse struct {
	UID         string   `json:"uid"`
	Username    string   `json:"username"`
	GECOS       string   `json:"gecos"`
	Disable     bool     `json:"disable"`
	Roles       []string `json:"roles"`
	FirstName   string   `json:"firstname"`
	FirstNameTH string   `json:"firstnameth"`
	LastName    string   `json:"lastname"`
	LastNameTH  string   `json:"lastnameth"`
	OUID        string   `json:"ouid"`
	Email       string   `json:"email"`
}
