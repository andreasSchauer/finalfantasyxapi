package queryparam

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

// ParseText checks, if a simple string parameter is valid, without any other checks
func (q QueryParam) ParseText(r *http.Request) (string, error) {
	text, err := q.CheckEmptyNoLower(r)
	if errs.ErrExceptEmptyQuery(err) {
		return "", err
	}

	decodedText, err := url.QueryUnescape(text)
	if err != nil {
		return "", errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid query encoding for parameter '%s'", q.Name), err)
	}

	return decodedText, nil
}
