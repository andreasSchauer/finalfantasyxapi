package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// ParseText checks, if a simple string parameter is valid, without any other checks
func (q QueryParam) ParseText(r *http.Request) (string, error) {
	text, err := q.CheckEmptyNoLower(r)
	if errExceptEmptyQuery(err) {
		return "", err
	}

	decodedText, err := url.QueryUnescape(text)
	if err != nil {
		return "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid query encoding for parameter '%s'", q.Name), err)
	}

	return decodedText, nil
}
