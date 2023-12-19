package controller

import (
	"goback/src/model"
	"goback/src/model/dto"
	"goback/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	user, err := uc.userService.GetUserByID(userID)
	data := map[string]interface{}{
		"user": user,
	}

	if err == nil {
		c.IndentedJSON(http.StatusOK, dto.JsonSuccess{
			Data: data,
		})
	} else {
		errorMessage := "Failed to get user: " + err.Error()
		c.IndentedJSON(http.StatusOK, dto.JsonError{
			Message: errorMessage,
		})
	}
}

func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	data := map[string]interface{}{
		"users": users,
	}

	if err == nil {
		c.IndentedJSON(http.StatusOK, dto.JsonSuccess{
			Data: data,
		})

	} else {
		errorMessage := "Failed to get users: " + err.Error()
		c.IndentedJSON(http.StatusOK, dto.JsonError{
			Message: errorMessage,
		})
	}
}

func (uc *UserController) AddUser(c *gin.Context) {
	var newUser model.User
	err := c.ShouldBindJSON(&newUser)

	if err != nil {
		errorMessage := "Failed to parse user data: " + err.Error()
		c.IndentedJSON(http.StatusBadRequest, dto.JsonError{
			Message: errorMessage,
		})
		return
	}

	_, err = uc.userService.AddUser(newUser)

	if err == nil {
		c.IndentedJSON(http.StatusCreated, dto.JsonSuccess{
			Data: "User added successfully",
		})
	} else {
		errorMessage := "Failed to add user: " + err.Error()
		c.IndentedJSON(http.StatusInternalServerError, dto.JsonError{
			Message: errorMessage,
		})
	}
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var updateUser model.User
	err := c.ShouldBindJSON(&updateUser)

	if err != nil {
		errorMessage := "Failed to parse user data: " + err.Error()
		c.IndentedJSON(http.StatusBadRequest, dto.JsonError{
			Message: errorMessage,
		})
		return
	}

	isUpdated, updateError := uc.userService.UpdateUser(updateUser)

	if updateError == nil && isUpdated {
		c.IndentedJSON(http.StatusCreated, dto.JsonSuccess{
			Data: "User Updated successfully",
		})
	} else {
		errorMessage := "Failed to Update user: " + err.Error()
		c.IndentedJSON(http.StatusInternalServerError, dto.JsonError{
			Message: errorMessage,
		})
	}
}
