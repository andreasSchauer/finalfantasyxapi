package api

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


// validates typed-QueryParam and checks emptiness
func parseTypedQuery(r *http.Request, queryParam QueryType) (string, string, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return "", "", err
	}

	resType, unitStr, found := strings.Cut(query, ":")
	if !found {
		return "", "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'.", queryParam.Name, query, queryParam.Usage), nil)
	}

	if !slices.Contains(queryParam.AllowedResTypes, resType) {
		return "", "", newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resource type '%s' for parameter '%s'. supported resource types: %s.", resType, queryParam.Name, h.FormatStringSlice(queryParam.AllowedResTypes)), nil)
	}

	return resType, unitStr, nil
}
