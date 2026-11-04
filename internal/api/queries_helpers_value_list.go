package api

import (
	"net/http"
)

func (q QueryParam) parseValueList(srv *Server, r *http.Request) ([]string, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return nil, err
	}

	return q.sliceValues(srv, query)
}

func (q QueryParam) sliceValues(srv *Server, query string) ([]string, error) {
	values, err := queryListSplit(srv, query)
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		err := q.checkValue(value)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}
