package albhed

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func setTranslator(translate, newState bool) (bool, error) {
	// this essentially means, the same bracket came twice without the other in between, like '[['
	if translate == newState {
		return false, errs.NewHTTPError(http.StatusBadRequest, "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket", nil)
	}
	return newState, nil
}

func initCharLookup(direction string) map[rune]rune {
	charLookup := make(map[rune]rune)

	for _, primer := range lookup.L.Primers {
		alBhedChar := charToRune(primer.AlBhedLetter)
		englishChar := charToRune(primer.EnglishLetter)

		if direction == string(database.TranslationDirectionToAlBhed) {
			charLookup[englishChar] = alBhedChar
			continue
		}

		charLookup[alBhedChar] = englishChar
	}

	return charLookup
}

func charToRune(char string) rune {
	runes := []rune(char)
	return runes[0]
}
