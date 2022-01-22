package service

import (
	"service-acl/entity"
	"service-acl/exception/validation"
	"service-acl/model"
	"service-acl/repository"
)

type menuServiceImpl struct {
	repository repository.MenuRepository
}

func NewMenuService(menuRepository *repository.MenuRepository) MenuService {
	return &menuServiceImpl{*menuRepository}
}

func MapMenu(menu entity.Menu) model.GetAllMenuResponse {
	response := model.GetAllMenuResponse{}
	response.ID = menu.ID
	response.Position = menu.Position
	response.ParentID = menu.ParentID
	response.Name = menu.Name
	response.Url = menu.Url
	response.IsActive = menu.IsActive
	response.Icon = menu.Icon
	response.Description = menu.Description

	response_sub := []model.GetAllMenuResponse{}

	for _, sub_menu := range menu.SubMenu {
		response_sub = append(response_sub, MapMenu(sub_menu))
	}
	response.SubMenu = response_sub
	return response
}

func (service *menuServiceImpl) List() ([]model.GetAllMenuResponse, error) {
	//get data menu
	menus := service.repository.FindAll()

	// mapping response
	responses := []model.GetAllMenuResponse{}

	for _, menu := range menus {
		responses = append(responses, MapMenu(menu))
	}

	return responses, nil
}

func (service *menuServiceImpl) Detail(id int) (model.GetMenuResponse, error) {
	menu, err := service.repository.FindById(id)

	if err != nil {
		return model.GetMenuResponse{}, err
	}

	//mapping response
	response_permission := model.GetPermissionResponse{}
	response_permission.ID = menu.Permission.ID
	response_permission.Name = menu.Permission.Name
	response_permission.Description = menu.Permission.Description

	var response = model.GetMenuResponse{}
	response.ID = menu.ID
	response.Position = menu.Position
	response.ParentID = menu.ParentID
	response.Name = menu.Name
	response.Url = menu.Url
	response.IsActive = menu.IsActive
	response.Icon = menu.Icon
	response.Description = menu.Description
	response.Permission = response_permission

	return response, nil
}

func (service *menuServiceImpl) Create(input model.CreateMenuRequest) (model.CreateMenuResponse, error) {
	//validate input
	validation.CreateMenuValidate(input)

	menu := entity.Menu{}
	menu.Position = input.Position
	menu.ParentID = input.ParentID
	menu.PermissionID = &input.PermissionID
	menu.Name = input.Name
	menu.Url = input.Url
	menu.IsActive = input.IsActive
	menu.Icon = input.Icon
	menu.Description = input.Description

	//update menu
	service.repository.Create(menu)

	//mapping response
	var response = model.CreateMenuResponse{}
	response.Name = menu.Name
	response.Url = menu.Url

	return response, nil
}

func (service *menuServiceImpl) Edit(id int, input model.UpdateMenuRequest) (model.UpdateMenuResponse, error) {
	//validate input
	validation.UpdateMenuValidate(input)

	//check menu
	_, err := service.repository.FindById(id)

	if err != nil {
		return model.UpdateMenuResponse{}, err
	}

	menu := entity.Menu{}
	menu.Position = input.Position
	menu.ParentID = input.ParentID
	menu.PermissionID = &input.PermissionID
	menu.Name = input.Name
	menu.Url = input.Url
	menu.IsActive = input.IsActive
	menu.Icon = input.Icon
	menu.Description = input.Description
	//update menu
	service.repository.Update(id, menu)

	//mapping response
	var response = model.UpdateMenuResponse{}
	response.Name = menu.Name
	response.Url = menu.Url

	return response, nil
}

func (service *menuServiceImpl) Delete(id int) error {
	_, err := service.repository.FindById(id)

	if err != nil {
		return err
	}

	//delete menu
	service.repository.Delete(id)

	return nil
}
