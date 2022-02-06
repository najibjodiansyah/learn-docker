package common

type LoginRequestFormat struct {
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginResponseFormat struct {
	Token string `json:"token"`
}