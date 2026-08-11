package api

import (
	"net/http"
	"unicode"
)

type AlBhedResponse struct {
	TranslatedText string 	`json:"translated_text"`
	OriginalText   string 	`json:"original_text"`
	FromLanguage   Language `json:"from_language"`
	ToLanguage     Language `json:"to_language"`
}

type Language string

const (
	langAlBhed		Language = "al bhed"
	langEnglish		Language = "english"
)

func serviceAlBhed(cfg *Config, r *http.Request, i handlerInputService[AlBhedResponse]) (AlBhedResponse, error) {
	err := verifyQueryParams[any](r, i.endpoint, i.queryLookup, nil, nil)
	if err != nil {
		return AlBhedResponse{}, err
	}

	response, err := initAlBhedResponse(r, i.queryLookup)
	if err != nil {
		return AlBhedResponse{}, err
	}

	return translateAlBhed(cfg, response)
}

func initAlBhedResponse(r *http.Request, queryLookup map[QueryParamName]QueryParam) (AlBhedResponse, error) {
	response := AlBhedResponse{
		FromLanguage: langAlBhed,
		ToLanguage:   langEnglish,
	}

	en, err := parseBooleanQuery(r, queryLookup[qpnEn])
	if errExceptEmptyQuery(err) {
		return AlBhedResponse{}, err
	}

	if en {
		response.FromLanguage = langEnglish
		response.ToLanguage = langAlBhed
	}

	text, err := parseTextQuery(r, queryLookup[qpnText])
	if errExceptEmptyQuery(err) {
		return AlBhedResponse{}, err
	}

	response.OriginalText = text

	return response, nil
}

func translateAlBhed(cfg *Config, response AlBhedResponse) (AlBhedResponse, error) {
	charLookup := initCharLookup(cfg, response.ToLanguage)
	translatedRunes := []rune{}
	translate := true
	var err error

	for _, char := range response.OriginalText {
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

func initCharLookup(cfg *Config, toLanguage Language) map[rune]rune {
	charLookup := make(map[rune]rune)
	
	for _, primer := range cfg.l.Primers {
		alBhedChar := charToRune(primer.AlBhedLetter)
		englishChar := charToRune(primer.EnglishLetter)

		if toLanguage == langAlBhed {
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