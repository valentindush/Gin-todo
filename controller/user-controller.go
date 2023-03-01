package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context)
}

type userController struct{
	service service.UserService
}

