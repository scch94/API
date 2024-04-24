package validation

import (
	"fmt"
	"regexp"

	"github.com/scch94/API.git/constants"
)

func validateRegex(number string) error {
	regex, err := regexp.Compile(constants.MOBIL_REGEX)
	if err != nil {
		return fmt.Errorf("error al compilar la expression regular porfavor verificala e inteta de neuvo regexp.Compile(): %w", err)
	}
	if !regex.MatchString(number) {
		return fmt.Errorf("el numero no machea con la expresion regular regex.MatchString(): %w", err)
	}
	return nil
}
