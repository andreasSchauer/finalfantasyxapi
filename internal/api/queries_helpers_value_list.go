package api

import (
	"net/http"
)

func (q QueryParam) parseValueList(server *Server, r *http.Request) ([]string, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return nil, err
	}

	return q.sliceValues(server, query)
}

func (q QueryParam) sliceValues(server *Server, query string) ([]string, error) {
	values, err := queryListSplit(server, query)
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
