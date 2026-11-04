package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type ListParams struct {
	Count    int     `json:"count"`
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
}

func createPaginatedList[T any](server *Server, r *http.Request, items []T) (ListParams, []T, error) {
	queryParamOffset := Q.defaultParams[ts.QpnOffset]
	queryParamLimit := Q.defaultParams[ts.QpnLimit]

	offset, err := queryParamOffset.parseInt(r)
	if err != nil {
		return ListParams{}, nil, err
	}

	limit, err := queryParamLimit.parseInt(r)
	if err != nil {
		return ListParams{}, nil, err
	}
	if limit == 0 {
		limit = *queryParamLimit.DefaultVal
	}

	size := len(items)

	listParams := ListParams{
		Count:    size,
		Previous: createPreviousURL(server, r, offset, limit),
		Next:     createNextURL(server, r, offset, limit, len(items)),
	}

	if size == 0 {
		return listParams, []T{}, nil
	}

	if offset >= size {
		return ListParams{}, nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("offset must be smaller than number of resources (%d)", size), err)
	}

	upperLimit := min(offset+limit, size)
	shownResources := items[offset:upperLimit]

	return listParams, shownResources, nil
}

func createPreviousURL(server *Server, r *http.Request, offset, limit int) *string {
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
	q.Set(string(ts.QpnLimit), strconv.Itoa(limit))
	q.Set(string(ts.QpnOffset), strconv.Itoa(previousOffset))

	previousURL := fmt.Sprintf("http://%s%s?%s", server.Cfg.Host, path, q.Encode())
	return &previousURL
}

func createNextURL(server *Server, r *http.Request, offset, limit, size int) *string {
	nextOffset := offset + limit

	if nextOffset >= size {
		return nil
	}

	path := strings.TrimSuffix(r.URL.Path, "/")
	q := r.URL.Query()
	q.Set(string(ts.QpnLimit), strconv.Itoa(limit))
	q.Set(string(ts.QpnOffset), strconv.Itoa(nextOffset))

	nextURL := fmt.Sprintf("http://%s%s?%s", server.Cfg.Host, path, q.Encode())
	return &nextURL
}
