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
	var TLV string
	TLV, err = validation.Validacion(*petition)
	if err != nil {
		fmt.Println("error en el momento de validar los datos de la peticion validation.Validacion()")
		c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println("this is the TLV", TLV)
	//aqui vamos al gateway de carlos para enviar el mensaje

	c.String(http.StatusOK, "XML parsed successfully")
}
