package router

import (
	"api_gopher_library/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func MapRoutes() {

	// CRUD users
	router.POST("/create_user", controllers.CreatingUser)
	router.GET("/get_user/:id", controllers.GettingUser)
	router.GET("/get_all_users", controllers.GettingUsers)
	router.PUT("/update_user", controllers.UpdatingUser)
	router.DELETE("/delete_user/:id", controllers.DeletingUser)

	// Read book
	router.GET("/get_book", controllers.GettingBook)

	// CRUD loans
	router.POST("/create_loan", controllers.CreatingLoan)
	router.GET("/get_all_loans", controllers.GettingLoans)
	router.GET("/get_loan/:id", controllers.GettingLoan)
	router.PUT("/update_loan", controllers.UpdatingLoan)
	router.DELETE("/delete_loan/:id", controllers.DeletingLoan)

}

func Run() {
	router.Run(":8080")
}
