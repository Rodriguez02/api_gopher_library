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
	user, err := services.GetAllUsers()
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func GettingUser(c *gin.Context) {
	id := c.Param("id")

	user, err := services.GetUser(id)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdatingUser(c *gin.Context) {
	var user domain.User

	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	user, err = services.UpdateUser(user)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeletingUser(c *gin.Context) {
	id := c.Param("id")

	user, err := services.DeleteUser(id)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func GettingBook(c *gin.Context) {
	var book domain.Book
	var result []domain.Information

	err := c.BindJSON(&book)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	result, err = services.GetBook(book)
	if err != nil {
		apiErr := parseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
