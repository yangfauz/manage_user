package repository

import (
	"service-acl/entity"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{database}
}

func (repository *userRepositoryImpl) FindAll() (users []entity.User) {

	err := repository.database.Preload("Role.Permission").Find(&users).Error

	if err != nil {
		panic(err)
	}

	return users
}

func (repository *userRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	err := repository.database.Preload("Role.Permission").Where("username = ?", username).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) FindById(id int) (entity.User, error) {
	var user entity.User
	err := repository.database.Preload("Role.Permission").Where("id = ?", id).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *userRepositoryImpl) Create(user entity.User) {
	err := repository.database.Create(&user).Error

	if err != nil {
		panic(err)
	}
}

func (repository *userRepositoryImpl) Update(id int, user entity.User) {
	err := repository.database.Model(&user).Where("id = ?", id).
		Updates(map[string]interface{}{"name": user.Name}).Error

	if err != nil {
		panic(err)
	}
}

func (repository *userRepositoryImpl) UserRole(user entity.User, role_id []int) {

	roles := make([]*entity.Role, len(role_id))

	for i, id := range role_id {
		roles[i] = &entity.Role{ID: uint(id)}
	}

	err := repository.database.Model(&user).Association("Role").Replace(roles)

	if err != nil {
		panic(err)
	}
}
