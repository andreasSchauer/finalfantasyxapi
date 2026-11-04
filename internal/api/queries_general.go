package api

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

// used for query filters that can't really be generalized. this one simply checks, if it's empty and then calls the wrapperFn
func (i handlerInput[T, R, A]) queryBasicWrapper(server *Server, r *http.Request, inputIDs []int32, queryName ts.QueryParamName, wrapperFn QueryWrapBasic) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		query, err := queryParam.basicChecks(r, i.queryLookup)
		if err != nil {
			return inputIDs, nil
		}

		dbIDs, err := wrapperFn(server, r, ctx, query, queryParam)
		if errors.Is(err, errs.ErrQueryRedirect) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		return dbIDs, nil
	}
}

func (q QueryParam) basicChecks(r *http.Request, queryLookup map[ts.QueryParamName]QueryParam) (string, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return "", err
	}

	if q.replacingParamsPresent(r, queryLookup) {
		return "", errs.ErrQueryRedirect
	}

	return query, nil
}

// checks, if a queryParam is empty and returns errEmptyQuery, if it is
func (q QueryParam) checkEmpty(r *http.Request) (string, error) {
	query := r.URL.Query().Get(string(q.Name))
	if query == "" {
		return "", errs.ErrEmptyQuery
	}

	return strings.ToLower(query), nil
}

func (q QueryParam) checkEmptyNoLower(r *http.Request) (string, error) {
	query := r.URL.Query().Get(string(q.Name))
	if query == "" {
		return "", errs.ErrEmptyQuery
	}

	return query, nil
}

// checks, if "none" was used as input and returns errQueryNone, if it was
func checkNoneQuery(query string) error {
	if query == "none" {
		return errs.ErrQueryNone
	}

	return nil
}

func (q QueryParam) replacingParamsPresent(r *http.Request, queryLookup map[ts.QueryParamName]QueryParam) bool {
	for _, param := range q.ReplacedBy {
		p := queryLookup[param]
		_, err := p.checkEmpty(r)
		if !queryIsEmpty(err) {
			return true
		}
	}
	return false
}
