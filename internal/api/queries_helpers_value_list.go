package api

import (
	"net/http"
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
		err := checkValue(queryParam, value)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}