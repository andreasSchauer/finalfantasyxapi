package api

import (
	"fmt"
	"net/http"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func parseValueListQuery(cfg *Config, r *http.Request, queryParam QueryParam) ([]string, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	return queryValuesToSlice(cfg, query, queryParam)
}

func queryValuesToSlice(cfg *Config, query string, queryParam QueryParam) ([]string, error) {
	values, err := queryListSplit(cfg, query)
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		if !slices.Contains(queryParam.AllowedValues, value) {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s.", query, queryParam.Name, h.FormatStringSlice(queryParam.AllowedValues)), nil)
		}
	}

	return values, nil
}