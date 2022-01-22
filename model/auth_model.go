package model

//request
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//response
type RegisterResponse struct {
	Username string `json:"username"`
}

type LoginResponse struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
