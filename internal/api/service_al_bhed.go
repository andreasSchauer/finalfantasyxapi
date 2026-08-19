package api

import (
	"net/http"
	"unicode"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type AlBhedResponse struct {
	URL            string     `json:"url"`
	TranslatedText string     `json:"translated_text"`
	OriginalText   string     `json:"original_text"`
	Direction      string 	  `json:"direction"`
}


// the serviceGet and servicePost functions look like two functions that can probably be replaced by a generic function.
// they basically take the verifyFn and the serviceFn as extras
// could also make them methods of AlBhedParams (*params.verify, params.execute) and create one or two interfaces
func serviceAlBhedGet(cfg *Config, r *http.Request, i handlerInputService[AlBhedResponse]) (AlBhedResponse, error) {
	var params AlBhedParams
	params, payloadMap, err := convertStateParam(r, i.queryLookup, params)
	if err != nil {
		return AlBhedResponse{}, err
	}

	url, err := convertToStateURL(cfg, string(i.endpoint), params)
	if err != nil {
		return AlBhedResponse{}, err
	}

	params, err = verifyAlBhedParams(cfg, params, payloadMap)
	if err != nil {
		return AlBhedResponse{}, err
	}

	return translateAlBhed(cfg, params, url)
}


func serviceAlBhedPost(cfg *Config, r *http.Request, i handlerInputService[AlBhedResponse]) (AlBhedResponse, error) {
	var params AlBhedParams
	params, payloadMap, err := decodeJsonBody(r, params)
	if err != nil {
		return AlBhedResponse{}, err
	}

	url, err := convertToStateURL(cfg, string(i.endpoint), params)
	if err != nil {
		return AlBhedResponse{}, err
	}

	params, err = verifyAlBhedParams(cfg, params, payloadMap)
	if err != nil {
		return AlBhedResponse{}, err
	}

	return translateAlBhed(cfg, params, url)
}



func translateAlBhed(cfg *Config, params AlBhedParams, url string) (AlBhedResponse, error) {
	response := AlBhedResponse{
		OriginalText: params.Text,
		Direction:    params.Direction,
		URL: 		  url,
	}
	
	charLookup := initCharLookup(cfg, params.Direction)
	translatedRunes := []rune{}
	translate := true
	var err error

	for _, char := range params.Text {
		if string(char) == "[" {
			translate, err = setTranslator(translate, false)
			if err != nil {
				return AlBhedResponse{}, err
			}
			continue
		}

		if string(char) == "]" {
			translate, err = setTranslator(translate, true)
			if err != nil {
				return AlBhedResponse{}, err
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
		return AlBhedResponse{}, newHTTPError(http.StatusBadRequest, "wrong format. all brackets must be closed.", nil)
	}

	response.TranslatedText = string(translatedRunes)

	return response, nil
}

func setTranslator(translate, newState bool) (bool, error) {
	// this essentially means, the same bracket came twice without the other in between, like '[['
	if translate == newState {
		return false, newHTTPError(http.StatusBadRequest, "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket.", nil)
	}
	return newState, nil
}

func initCharLookup(cfg *Config, direction string) map[rune]rune {
	charLookup := make(map[rune]rune)

	for _, primer := range cfg.l.Primers {
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
