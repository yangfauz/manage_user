package model

//request
type CreatePermissionRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
}

type UpdatePermissionRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
}

//response
type CreatePermissionResponse struct {
	Name string `json:"name"`
}

type UpdatePermissionResponse struct {
	Name string `json:"name"`
}

type GetPermissionResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
