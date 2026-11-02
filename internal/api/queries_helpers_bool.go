package api

import (
	"fmt"
	"net/http"
	"strconv"
)

// validates bool-queryParam and checks emptiness.
func (q QueryParam) parseBool(r *http.Request) (bool, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid boolean value '%s' used for parameter '%s'. usage: '%s'", query, q.Name, q.Usage), err)
	}

	return b, nil
}
