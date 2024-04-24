package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scch94/API.git/constants"
	"github.com/scch94/API.git/routes"
)

func main() {
	//para evitar logeo de gin
	gin.SetMode(gin.ReleaseMode)
	//creas el servidor
	router := gin.Default()
	fmt.Println("starting micropagos gateway")
	//configuramos las rutas del servidor
	routes.Routes(router)
	//configuracion del servidor rutas y puerto
	serverConfig := &http.Server{
		Addr:    constants.PORT,
		Handler: router,
	}
	//arrancamos el servidor
	err := serverConfig.ListenAndServe()
	fmt.Println("estamos escuchando en el puerto", constants.PORT)
	if err != nil {
		fmt.Println("cant connect to the server: ", err)
		return
	}
}
