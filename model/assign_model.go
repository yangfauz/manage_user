package model

//request
type UserRoleRequest struct {
	UserID int   `json:"user_id" binding:"required"`
	RoleID []int `json:"role_id" binding:"required"`
}

type RolePermissionRequest struct {
	RoleID       int   `json:"role_id" binding:"required"`
	PermissionID []int `json:"permission_id" binding:"required"`
}

//response
type UserRoleResponse struct {
	UserID int `json:"user_id"`
}

type RolePermissionResponse struct {
	RoleID int `json:"role_id"`
}
