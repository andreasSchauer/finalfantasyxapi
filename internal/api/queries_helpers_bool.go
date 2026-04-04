package api

import (
	"fmt"
	"net/http"
	"strconv"
)

// validates bool-queryParam and checks emptiness.
func parseBooleanQuery(r *http.Request, queryParam QueryType) (bool, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid boolean value '%s' used for parameter '%s'. usage: '%s'.", query, queryParam.Name, queryParam.Usage), err)
	}

	return b, nil
}