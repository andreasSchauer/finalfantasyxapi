package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// validates typed-QueryParam and checks emptiness
func (q QueryParam) parseTyped(r *http.Request) (string, string, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return "", "", err
	}

	resType, unitStr, found := strings.Cut(query, ":")
	if !found {
		return "", "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'", q.Name, query, q.Usage), nil)
	}

	if !slices.Contains(q.AllowedResTypes, resType) {
		return "", "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resource type '%s' for parameter '%s'. supported resource types: %s", resType, q.Name, h.FormatStringSlice(q.AllowedResTypes)), nil)
	}

	return resType, unitStr, nil
}
