package validation

import (
	"fmt"

	"github.com/scch94/API.git/structs"
)

func Validacion(p structs.Petition) (string, error) {
	var messagelenght uint16
	//validamos que el numero cumpla con la expresion regular
	err := validateRegex(p.Mobile())
	if err != nil {
		return "", fmt.Errorf("error al validar la expresion regular validateRegex(): %w", err)
	}
	//validamos el largo del mensaje
	messagelenght, err = validateMessageLength(p.Message())
	if err != nil {
		return "", fmt.Errorf("error in (#words of the message: %v) validatemessageLength () %w", messagelenght, err)
	}
	//validamos que la palabra no tenga tildes o ñ
	err = validateNumberRegex(p.Message())
	if err != nil {
		return "", fmt.Errorf("error al tratar de validar el mensaje validateNumberRegex(): %w", err)
	}
	//vamos a la api de portabilidad para ver a que cliente debe atender
	return "CLARO", nil
}
