package service

import "gilab.com/pragmaticreviews/golang-gin-poc/entity"

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
}

type userService struct{
	users []entity.User
}

func New() UserService{
	return &userService{}
}

func (service *userService) Save(entity.User) entity.User
func (service *userService) FindAll() []entity.User{
	return service.users
}