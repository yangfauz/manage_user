package model

//request
type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

//response
type UpdateUserResponse struct {
	Name string `json:"name"`
}

type GetUserResponse struct {
	ID       uint              `json:"id"`
	Name     string            `json:"name"`
	Username string            `json:"username"`
	Roles    []GetRoleResponse `json:"roles"`
}
