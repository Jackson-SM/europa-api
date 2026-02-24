package controllers

import (
	"net/http"

	"github.com/Jackson-SM/Europa/cmd/internal/entities"
	"github.com/Jackson-SM/Europa/cmd/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository interfaces.UserRepositoryInterface
}

func NewUserController(userRepository interfaces.UserRepositoryInterface) *UserController {
	return &UserController{userRepository: userRepository}
}

func (uc *UserController) Create(ctx *gin.Context) {
	var user *entities.User = entities.NewUser("Jackson", "example@gmail.com")

	ctx.JSON(http.StatusCreated, user)

	user, msg := uc.userRepository.Create(user)

	print(user, msg)
}