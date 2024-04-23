package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/scch94/API.git/constants"
	"github.com/scch94/API.git/routes"
)

func main() {
	router := gin.Default()

	//configuramos las rutas
	routes.Routes(router)
	err := router.Run(constants.PORT)
	if err != nil {
		fmt.Println("cant connect to the server: ", err)
		return
	}
	fmt.Println("server is running on port ", constants.PORT)
}
