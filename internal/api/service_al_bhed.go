package api

import (
	"net/http"
	"unicode"
)

type AlBhedParams struct {
	Text      string     `json:"text"`
	Direction QueryValue `json:"direction"`
}

type AlBhedResponse struct {
	URL            string     `json:"url"`
	TranslatedText string     `json:"translated_text"`
	OriginalText   string     `json:"original_text"`
	Direction      QueryValue `json:"direction"`
}

const (
	qvToAlBhed  QueryValue = "to-al-bhed"
	qvToEnglish QueryValue = "to-english"
)

func serviceAlBhedGet(cfg *Config, r *http.Request, i handlerInputService[AlBhedResponse]) (AlBhedResponse, error) {
	params, err := initAlBhedParams(r, i.queryLookup)
	if err != nil {
		return AlBhedResponse{}, err
	}

	return translateAlBhed(cfg, i, params)
}

func serviceAlBhedPost(cfg *Config, r *http.Request, i handlerInputService[AlBhedResponse]) (AlBhedResponse, error) {
	var params AlBhedParams
	params, err := decodeJsonBody(r, params)
	if err != nil {
		return AlBhedResponse{}, err
	}

	return translateAlBhed(cfg, i, params)
}


func initAlBhedParams(r *http.Request, queryLookup map[QueryParamName]QueryParam) (AlBhedParams, error) {
	var response AlBhedParams

	response, err := convertStateParam(r, queryLookup, response)
	if errExceptEmptyQuery(err) {
		return AlBhedParams{}, err
	}
	if !queryIsEmpty(err) {
		return response, nil
	}

	response.Direction = qvToAlBhed

	direction, err := parseValueQuery(r, queryLookup[qpnDirection])
	if errExceptEmptyQuery(err) {
		return AlBhedParams{}, err
	}
	if direction == string(qvToEnglish) {
		response.Direction = qvToEnglish
	}

	text, err := parseTextQuery(r, queryLookup[qpnText])
	if errExceptEmptyQuery(err) {
		return AlBhedParams{}, err
	}
	response.Text = text

	return response, nil
}

func translateAlBhed(cfg *Config, i handlerInputService[AlBhedResponse], params AlBhedParams) (AlBhedResponse, error) {
	charLookup := initCharLookup(cfg, params.Direction)
	response := AlBhedResponse{
		OriginalText: params.Text,
		Direction:    params.Direction,
	}
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
	response.URL, err = convertToStateURL(cfg, string(i.endpoint), params)
	if err != nil {
		return AlBhedResponse{}, err
	}

	return response, nil
}

func setTranslator(translate, newState bool) (bool, error) {
	// this essentially means, the same bracket came twice without the other in between, like '[['
	if translate == newState {
		return false, newHTTPError(http.StatusBadRequest, "wrong format. nested brackets are not allowed, and the first bracket must be an opening bracket.", nil)
	}
	return newState, nil
}

func initCharLookup(cfg *Config, direction QueryValue) map[rune]rune {
	charLookup := make(map[rune]rune)

	for _, primer := range cfg.l.Primers {
		alBhedChar := charToRune(primer.AlBhedLetter)
		englishChar := charToRune(primer.EnglishLetter)

		if direction == qvToAlBhed {
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
