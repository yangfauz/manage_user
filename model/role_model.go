package model

//request
type CreateRoleRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
}

type UpdateRoleRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
}

//response
type CreateRoleResponse struct {
	Name string `json:"name"`
}

type UpdateRoleResponse struct {
	Name string `json:"name"`
}

type GetRoleResponse struct {
	ID          uint                    `json:"id"`
	Name        string                  `json:"name"`
	Description *string                 `json:"description"`
	Permissions []GetPermissionResponse `json:"permissions"`
}
