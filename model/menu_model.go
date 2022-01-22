package model

//request
type CreateMenuRequest struct {
	Position     int     `json:"position" binding:"required"`
	ParentID     *uint   `json:"parent_id" binding:"required"`
	PermissionID uint    `json:"permission_id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Url          string  `json:"url" binding:"required"`
	IsActive     bool    `json:"is_active" binding:"required"`
	Icon         string  `json:"icon" binding:"required"`
	Description  *string `json:"description" binding:"required"`
}

type UpdateMenuRequest struct {
	Position     int     `json:"position" binding:"required"`
	ParentID     *uint   `json:"parent_id" binding:"required"`
	PermissionID uint    `json:"permission_id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Url          string  `json:"url" binding:"required"`
	IsActive     bool    `json:"is_active" binding:"required"`
	Icon         string  `json:"icon" binding:"required"`
	Description  *string `json:"description" binding:"required"`
}

//response
type CreateMenuResponse struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type UpdateMenuResponse struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GetMenuResponse struct {
	ID          uint                  `json:"id"`
	Position    int                   `json:"position"`
	ParentID    *uint                 `json:"parent_id"`
	Name        string                `json:"name"`
	Url         string                `json:"url"`
	IsActive    bool                  `json:"is_active"`
	Icon        string                `json:"icon"`
	Description *string               `json:"description"`
	Permission  GetPermissionResponse `json:"permission"`
	SubMenu     []GetMenuResponse     `json:"sub_menu"`
}

type GetAllMenuResponse struct {
	ID          uint                 `json:"id"`
	Position    int                  `json:"position"`
	ParentID    *uint                `json:"parent_id"`
	Name        string               `json:"name"`
	Url         string               `json:"url"`
	IsActive    bool                 `json:"is_active"`
	Icon        string               `json:"icon"`
	Description *string              `json:"description"`
	SubMenu     []GetAllMenuResponse `json:"sub_menu"`
}
