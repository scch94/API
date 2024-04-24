package validation

import (
	"fmt"
	"regexp"

	"github.com/scch94/API.git/constants"
)

func validateMessageLength(message string) (uint16, error) {
	messagelenght := len(message)
	if messagelenght > constants.MAX_MESSAGE_LENGTH {
		return uint16(messagelenght), fmt.Errorf("el texto del mensaje supera los %v el texto tiene %w", constants.MAX_MESSAGE_LENGTH, messagelenght)
	}
	return uint16(messagelenght), nil
}
func validateNumberRegex(number string) error {
	regex, err := regexp.Compile(constants.MESSAGE_REGEX)
	if err != nil {
		return fmt.Errorf("error al compilar la expression regular porfavor verificala e inteta de neuvo regexp.Compile(): %w", err)
	}
	if !regex.MatchString(number) {
		return fmt.Errorf("el mensaje tiene tilde o una ñ y por eso no es permitido regex.MatchString(): %w", err)
	}
	return nil
}
