package main

import(
	"fmt"
	"net/http"
)


type ListParams struct {
	Count		int					`json:"count"`
	Next		*string				`json:"next"`
	Previous	*string				`json:"previous"`
}


func createPaginatedList[T any](cfg *apiConfig, r *http.Request, resources []T) (ListParams, []T, error) {
	const defaultOffset = 0
	const defaultLimit = 20
	
	offset, err := queryStrToInt(r.URL.Query().Get("offset"), defaultOffset)
	if err != nil {
		return ListParams{}, nil, newHTTPError(http.StatusBadRequest, "invalid offset provided", err)
	}

	limit, err := queryStrToInt(r.URL.Query().Get("limit"), defaultLimit)
	if err != nil {
		return ListParams{}, nil, newHTTPError(http.StatusBadRequest, "invalid limit provided", err)
	}

	if limit == 0 {
		limit = defaultLimit
	}

	path := r.URL.Path
	size := len(resources)
	
	if offset >= size {
		return ListParams{}, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("offset must be smaller than number of resources (%d)", size), err)
	}
	
	upperLimit := min(offset + limit, size)
	shownResources := resources[offset:upperLimit]

	listParams := ListParams{
		Count:   	size,
		Next: 		cfg.createNextURL(path, offset, limit, len(resources)),
		Previous: 	cfg.createPreviousURL(path, offset, limit),
	}

	return listParams, shownResources, nil
}


func (cfg *apiConfig) createNextURL(path string, offset, limit, size int) *string {
	nextOffset := offset + limit

	if nextOffset >= size {
		return nil
	}

	nextURL := fmt.Sprintf("http://%s%s?offset=%d&limit=%d", cfg.host, path, nextOffset, limit)
	return &nextURL
}


func (cfg *apiConfig) createPreviousURL(path string, offset, limit int) *string {
	if offset == 0 {
		return nil
	}

	previousOffset := offset - limit
	if previousOffset < 0 {
		previousOffset = 0
		limit = offset
	}

	previousURL := fmt.Sprintf("http://%s%s?offset=%d&limit=%d", cfg.host, path, previousOffset, limit)
	return &previousURL
}