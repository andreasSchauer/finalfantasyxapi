package api

import (
	"fmt"
	"net/http"
	"net/url"
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

	result, err := initAlBhedResponse(r, i.queryLookup)
	if err != nil {
		return AlBhedResponse{}, err
	}

	// assemble character map from al bhed primers (depending on from- and to-language)
	// then do the translation
	// add translated text to result

	return result, nil
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

// move to different file once everything is working
func parseTextQuery(r *http.Request, queryParam QueryParam) (string, error) {
	text, err := checkEmptyQuery(r, queryParam)
	if errExceptEmptyQuery(err) {
		return "", err
	}

	decodedText, err := url.QueryUnescape(text)
	if err != nil {
		return "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid query encoding for parameter '%s'.", queryParam.Name), err)
	}

	return decodedText, nil
}
