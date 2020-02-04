package router

import (
	"api_gopher_library/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func MapRoutes() {
	router.GET("/prueba", controllers.Prueba)
}

func Run() {
	router.Run(":8080")
}

/* code
.
.
.
*/
