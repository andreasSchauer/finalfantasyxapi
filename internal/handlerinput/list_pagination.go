package handlerinput

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func createPaginatedList[T any](r *http.Request, items []T) (apiresource.ListParams, []T, error) {
	defaultParams := queryparam.GetDefaultParamMap()
	queryParamOffset := defaultParams[ts.QpnOffset]
	queryParamLimit := defaultParams[ts.QpnLimit]

	offset, err := queryParamOffset.ParseInt(r)
	if err != nil {
		return apiresource.ListParams{}, nil, err
	}

	limit, err := queryParamLimit.ParseInt(r)
	if err != nil {
		return apiresource.ListParams{}, nil, err
	}
	if limit == 0 {
		limit = *queryParamLimit.DefaultVal
	}

	size := len(items)

	listParams := apiresource.ListParams{
		Count:    size,
		Previous: CreatePreviousURL(r, offset, limit),
		Next:     CreateNextURL(r, offset, limit, len(items)),
	}

	if size == 0 {
		return listParams, []T{}, nil
	}

	if offset >= size {
		return apiresource.ListParams{}, nil, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("offset must be smaller than number of resources (%d)", size), err)
	}

	upperLimit := min(offset+limit, size)
	shownResources := items[offset:upperLimit]

	return listParams, shownResources, nil
}

func CreatePreviousURL(r *http.Request, offset, limit int) *string {
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

	previousURL := fmt.Sprintf("http://%s%s?%s", config.Cfg.Host, path, q.Encode())
	return &previousURL
}

func CreateNextURL(r *http.Request, offset, limit, size int) *string {
	nextOffset := offset + limit

	if nextOffset >= size {
		return nil
	}

	path := strings.TrimSuffix(r.URL.Path, "/")
	q := r.URL.Query()
	q.Set(string(ts.QpnLimit), strconv.Itoa(limit))
	q.Set(string(ts.QpnOffset), strconv.Itoa(nextOffset))

	nextURL := fmt.Sprintf("http://%s%s?%s", config.Cfg.Host, path, q.Encode())
	return &nextURL
}
