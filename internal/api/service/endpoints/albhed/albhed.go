// Package albhed ajsdoasdoij
package albhed

import (
	"net/http"
	"unicode"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
)

func Handle(params servicetypes.AlBhedParams) (servicetypes.AlBhedResponse, error) {
	response := servicetypes.AlBhedResponse{
		OriginalText: params.Text,
		Direction:    params.Direction,
	}

	charLookup := initCharLookup(params.Direction)
	translatedRunes := []rune{}
	translate := true
	var err error

	for _, char := range params.Text {
		if string(char) == "[" {
			translate, err = setTranslator(translate, false)
			if err != nil {
				return servicetypes.AlBhedResponse{}, err
			}
			continue
		}

		if string(char) == "]" {
			translate, err = setTranslator(translate, true)
			if err != nil {
				return servicetypes.AlBhedResponse{}, err
			}
			continue
		}

		if !translate {
			translatedRunes = append(translatedRunes, char)
			continue
		}

		var isUpper bool
		charLower := unicode.ToLower(char)
		if char != charLower {
			isUpper = true
		}

		newChar, ok := charLookup[charLower]
		if !ok {
			translatedRunes = append(translatedRunes, char)
			continue
		}

		if isUpper {
			newChar = unicode.ToUpper(newChar)
		}

		translatedRunes = append(translatedRunes, newChar)
	}

	if !translate {
		return servicetypes.AlBhedResponse{}, errs.NewHTTPError(http.StatusBadRequest, "wrong format. all brackets must be closed", nil)
	}

	response.TranslatedText = string(translatedRunes)

	return response, nil
}
