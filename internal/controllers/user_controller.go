package controllers

import (
	"net/http"

	"github.com/Jackson-SM/Europa/internal/dtos"
	"github.com/Jackson-SM/Europa/internal/entities"
	"github.com/Jackson-SM/Europa/internal/interfaces"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository interfaces.UserRepositoryInterface
}

func NewUserController(userRepository interfaces.UserRepositoryInterface) *UserController {
	return &UserController{userRepository: userRepository}
}

func (uc *UserController) Create(ctx *gin.Context) {
	var body dtos.CreateUserDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user entities.User = *entities.NewUser(body.Name, body.Email)

	uc.userRepository.Create(&user)

	ctx.JSON(http.StatusCreated, user)
}

func (uc *UserController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := uc.userRepository.FindById(id)

	ctx.JSON(http.StatusOK, user)
}
