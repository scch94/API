package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/API.git/constants"
	"github.com/scch94/API.git/controller"
)

func Routes(router *gin.Engine) {
	//creas un grupo de rutas
	routes := router.Group("/")
	//aqui el path es la ruta que tienen actualmente para el envio de mensajes el controller en la funcion que correria cuando entran
	routes.POST(constants.PATH, controller.SendMessageService)
	//esta solo es una ruta de prueba
	routes.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "¡Se recibió una solicitud GET en la ruta raíz!")
	})
	//ruta post para recibir mensajes

}
