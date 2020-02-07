package controllers

import (
	"api_gopher_library/domain"
	"api_gopher_library/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

/***************************************************
******************CONTROLLERS LOANS*****************
***************************************************/

func CreatingLoan(c *gin.Context){
	var loan domain.Loan
	err := c.BindJSON(&loan)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error()) 
	}
	loan, err = services.CreateLoan(loan)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}
	c.JSON(http.StatusOK, loan)
}

func GettingLoans(c *gin.Context){
	loans, err := services.GetAllLoans()
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}
	c.JSON(http.StatusOK, loans)
}

func GettingLoan(c *gin.Context) {
	id := c.Param("id")

	loan, err := services.GetLoan(id)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, loan)
}

func UpdatingLoan(c *gin.Context){
	var loan domain.Loan

	err := c.BindJSON(&loan)
	if err != nil{
		c.String(http.StatusBadRequest, err.Error())
	}

	loan, err = services.UpdateLoan(loan)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, loan)
}

func DeletingLoan(c *gin.Context){
	id := c.Param("id")

	loan, err := services.DeleteLoan(id)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, loan)
}

/***************************************************
******************CONTROLLERS USERS*****************
***************************************************/

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

