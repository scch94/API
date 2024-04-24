package validation

import (
	"encoding/hex"
	"fmt"

	"github.com/scch94/API.git/structs"
)

func Validacion(p structs.Petition, r *structs.RequestPayload) error {
	var messagelenght uint16
	//validamos que el numero cumpla con la expresion regular
	err := validateRegex(p.Mobile())
	if err != nil {
		return fmt.Errorf("error al validar la expresion regular validateRegex(): %w", err)
	}
	//validamos el largo del mensaje
	messagelenght, err = validateMessageLength(p.Message())
	if err != nil {
		return fmt.Errorf("error in (#words of the message: %v) validatemessageLength () %w", messagelenght, err)
	}
	//validamos que la palabra no tenga tildes o ñ
	err = validateNumberRegex(p.Message())
	if err != nil {
		return fmt.Errorf("error al tratar de validar el mensaje validateNumberRegex(): %w", err)
	}
	//llenamos el request esto hira en otro modulo
	//vamos a la api de portabilidad para ver a que cliente debe atender
	r.TLVValue = "CLARO"
	r.ServiceType = ""
	r.OriginTon = 3
	r.OriginNpi = 0
	r.DestinationTon = 3
	r.DestinationNpi = 0
	r.ValidityPeriod = ""
	r.ScheduleDeliveryTime = ""
	r.ProtocolID = 0
	r.EsmeClass = 0
	r.PriorityFlag = 3
	r.RegisteredDelivery = 0
	r.ReplaceIfPresentFlag = 0
	//funcion para volver el texto jexadecimal
	r.Data = hex.EncodeToString([]byte(p.Message()))
	r.DataHeaderIndicator = 0
	r.DataCodingScheme = 0
	r.DataLength = messagelenght
	r.MessageType = 4
	r.TLVTag = 515
	r.TLVLength = 5
	return nil
}
