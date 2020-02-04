package controllers

import (
	"api_gopher_library/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Prueba(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetSaludo())
}

/* code
.
.
.
*/
