package queryparam

import (
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (q QueryParam) BasicChecks(r *http.Request, queryLookup map[ts.QueryParamName]QueryParam) (string, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return "", err
	}

	if q.ReplacingParamsPresent(r, queryLookup) {
		return "", errs.ErrQueryRedirect
	}

	return query, nil
}

// CheckEmpty checks, if a queryParam is empty and returns errEmptyQuery, if it is
func (q QueryParam) CheckEmpty(r *http.Request) (string, error) {
	query := r.URL.Query().Get(string(q.Name))
	if query == "" {
		return "", errs.ErrEmptyQuery
	}

	return strings.ToLower(query), nil
}

func (q QueryParam) CheckEmptyNoLower(r *http.Request) (string, error) {
	query := r.URL.Query().Get(string(q.Name))
	if query == "" {
		return "", errs.ErrEmptyQuery
	}

	return query, nil
}

func (q QueryParam) ReplacingParamsPresent(r *http.Request, queryLookup map[ts.QueryParamName]QueryParam) bool {
	for _, param := range q.ReplacedBy {
		p := queryLookup[param]
		_, err := p.CheckEmpty(r)
		if !errs.QueryIsEmpty(err) {
			return true
		}
	}
	return false
}
