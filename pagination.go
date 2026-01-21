package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ListParams struct {
	Count          int         `json:"count"`
	Previous       *string     `json:"previous"`
	Next           *string     `json:"next"`
}

func createPaginatedList[T any](cfg *Config, r *http.Request, items []T) (ListParams, []T, error) {
	queryParamOffset := cfg.q.defaultParams["offset"]
	queryParamLimit := cfg.q.defaultParams["limit"]

	offset, err := parseIntQuery(r, queryParamOffset)
	if err != nil {
		return ListParams{}, nil, err
	}

	limit, err := parseIntQuery(r, queryParamLimit)
	if err != nil {
		return ListParams{}, nil, err
	}
	if limit == 0 {
		limit = *queryParamLimit.DefaultVal
	}

	size := len(items)

	listParams := ListParams{
		Count:    size,
		Previous: cfg.createPreviousURL(r, offset, limit),
		Next:     cfg.createNextURL(r, offset, limit, len(items)),
	}

	if size == 0 {
		return listParams, []T{}, nil
	}

	if offset >= size {
		return ListParams{}, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("offset must be smaller than number of resources (%d).", size), err)
	}

	upperLimit := min(offset+limit, size)
	shownResources := items[offset:upperLimit]

	return listParams, shownResources, nil
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

	path := strings.TrimSuffix(r.URL.Path, "/")
	q := r.URL.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(previousOffset))

	previousURL := fmt.Sprintf("http://%s%s?%s", cfg.host, path, q.Encode())
	return &previousURL
}

func (cfg *Config) createNextURL(r *http.Request, offset, limit, size int) *string {
	nextOffset := offset + limit

	if nextOffset >= size {
		return nil
	}

	path := strings.TrimSuffix(r.URL.Path, "/")
	q := r.URL.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("offset", strconv.Itoa(nextOffset))

	nextURL := fmt.Sprintf("http://%s%s?%s", cfg.host, path, q.Encode())
	return &nextURL
}
