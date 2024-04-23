package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/API.git/constants"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/")
	routes.POST(constants.PATH, func(c *gin.Context) { c.String(http.StatusOK, "si recibimos la solicitud") })
	routes.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "¡Se recibió una solicitud GET en la ruta raíz!")
	})
	//ruta post para recibir mensajes

}
