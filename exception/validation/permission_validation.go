package validation

import (
	"service-acl/exception"
	"service-acl/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func CreatePermissionValidate(request model.CreatePermissionRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required.When(request.Name == "").Error("Name is Required")),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UpdatePermissionValidate(request model.UpdatePermissionRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required.When(request.Name == "").Error("Name is Required")),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
