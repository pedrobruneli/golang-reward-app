package users

import (
	"rewards/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) RegisterUserRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/users", uc.GetUsers)
	router.POST("/users", uc.CreateUser)
	router.GET("/users/:username", uc.GetUserByUsername)
	router.DELETE("/users/:username", uc.DeleteUser)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func (uc *UserController) GetUsers(ctx *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}

// CreateUser godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserDTO true "User object"
// @Success 201 {object} UserDTO
// @Router /users [post]
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user UserDTO
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = uc.UserService.CreateUser(&User{Username: user.Username})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, user)
}

// GetUserByUsername godoc
// @Summary Get a user by username
// @Description Get a user by username
// @Tags users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} User
// @Router /users/{username} [get]
func (uc *UserController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := uc.UserService.GetUserByUsername(username)
	if err != nil {
		statusCode, errorMessage := utils.MapError(err)
		ctx.JSON(statusCode, gin.H{"error": errorMessage})
		return
	}
	ctx.JSON(200, user)
}

// DeleteUser godoc
// @Summary Delete a user by username
// @Description Delete a user by username
// @Tags users
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 204
// @Router /users/{username} [delete]
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("username")
	err := uc.UserService.DeleteUser(username)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(204)
}
