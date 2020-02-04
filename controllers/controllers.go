package controllers

import (
	"api_gopher_library/domain"
	"api_gopher_library/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Status  int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func CreatingUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	user, err = services.CreateUser(user)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func GettingUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func GettingUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	user,err = services.GetUser(user)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func parseError(e error) ApiError {
	switch e {
	case services.ErrorNoName, services.ErrorNoSurname, services.ErrorInvalidID, services.ErrorUserExists:
		return ApiError{400, e.Error()}
	case services.ErrorUsersNotFound, services.ErrorUserNotFound:
		return ApiError{404, e.Error()}
	default:
		return ApiError{500, e.Error()}
	}
}
