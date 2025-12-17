package api

import (
	"fmt"
	"net/http"
	"strconv"
)

type ListParams struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}

func createPaginatedList[T any](cfg *Config, r *http.Request, resources []T) (ListParams, []T, error) {
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

	size := len(resources)

	listParams := ListParams{
		Count:    size,
		Next:     cfg.createNextURL(r, offset, limit, len(resources)),
		Previous: cfg.createPreviousURL(r, offset, limit),
	}

	if size == 0 {
		return listParams, []T{}, nil
	}

	if offset >= size {
		return ListParams{}, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("offset must be smaller than number of resources (%d)", size), err)
	}

	upperLimit := min(offset+limit, size)
	shownResources := resources[offset:upperLimit]

	return listParams, shownResources, nil
}

func (cfg *Config) createNextURL(r *http.Request, offset, limit, size int) *string {
	nextOffset := offset + limit

	if nextOffset >= size {
		return nil
	}

	path := r.URL.Path
	q := r.URL.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(nextOffset))

	nextURL := fmt.Sprintf("http://%s%s?%s", cfg.host, path, q.Encode())
	return &nextURL
}

func (cfg *Config) createPreviousURL(r *http.Request, offset, limit int) *string {
	if offset == 0 {
		return nil
	}

	previousOffset := offset - limit
	if previousOffset < 0 {
		previousOffset = 0
		limit = offset
	}

	path := r.URL.Path
	q := r.URL.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(previousOffset))

	previousURL := fmt.Sprintf("http://%s%s?%s", cfg.host, path, q.Encode())
	return &previousURL
}
