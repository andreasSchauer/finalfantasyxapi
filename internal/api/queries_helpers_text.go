package api

import (
	"fmt"
	"net/http"
	"net/url"
)

func parseTextQuery(r *http.Request, queryParam QueryParam) (string, error) {
	text, err := checkEmptyQueryNoLower(r, queryParam)
	if errExceptEmptyQuery(err) {
		return "", err
	}

	decodedText, err := url.QueryUnescape(text)
	if err != nil {
		return "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid query encoding for parameter '%s'.", queryParam.Name), err)
	}

	return decodedText, nil
}
