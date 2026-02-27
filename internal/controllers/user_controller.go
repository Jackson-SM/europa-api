package controllers

import (
	"net/http"

	"github.com/Jackson-SM/Europa/internal/domain"
	"github.com/Jackson-SM/Europa/internal/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryInterface interface {
	Create(user *domain.User) *domain.User
	FindById(id string) *domain.User
}

type UserController struct {
	userRepository UserRepositoryInterface
}

func NewUserController(userRepository UserRepositoryInterface) *UserController {
	return &UserController{userRepository: userRepository}
}

func (uc *UserController) Create(ctx *gin.Context) {

	var body dto.CreateUserRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return
	}

	var user domain.User = domain.User{
		Name: body.Name,
		Email: body.Email,
		Password: string(hashPassword),
	}

	createdUser := uc.userRepository.Create(&user)
	if createdUser == nil {
		ctx.JSON(http.StatusConflict, gin.H{"message": "It was not possible to insert the user in the database"})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := uc.userRepository.FindById(id)

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
