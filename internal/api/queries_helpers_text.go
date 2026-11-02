package api

import (
	"fmt"
	"net/http"
	"net/url"
)

func (q QueryParam) parseText(r *http.Request) (string, error) {
	text, err := q.checkEmptyNoLower(r)
	if errExceptEmptyQuery(err) {
		return "", err
	}

	decodedText, err := url.QueryUnescape(text)
	if err != nil {
		return "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid query encoding for parameter '%s'.", q.Name), err)
	}

	return decodedText, nil
}
