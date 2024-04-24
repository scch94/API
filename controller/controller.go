package controller

import (
	"fmt"
	"io"
	"net/http"

	"encoding/xml"

	"github.com/gin-gonic/gin"
	"github.com/scch94/API.git/structs"
	"github.com/scch94/API.git/validation"
)

var requestToGateway structs.RequestPayload

func SendMessageService(c *gin.Context) {

	fmt.Println("starting to get the xml")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error al leer el cuerpo de la solicitud:", err)
		c.String(http.StatusBadRequest, "bad body request")
		return
	}
	var envelope structs.Envelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		fmt.Println("Error al decodificar XML:", err)
		c.String(http.StatusBadRequest, "error decoding XML")
		return
	}
	//convertimos los datos de la peticion en una structura para faciliar su manejo y mostramos la informacion;
	petition := structs.NewPetition(envelope.Body.Send.Mobile, envelope.Body.Send.Message, envelope.Body.Send.UseOriginName)
	fmt.Println(petition.ToString())
	//vamos a validar la solicitud del envio de mensajes
	err = validation.Validacion(*petition, &requestToGateway)
	if err != nil {
		fmt.Println("error en el momento de validar los datos de la peticion validation.Validacion()")
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	requestToGateway.DestinationNumber = petition.Mobile()
	requestToGateway.OriginNumber = "0911234567"
	//aqui vamos al backend de envio de mensajes

	c.JSON(http.StatusOK, requestToGateway)
}
