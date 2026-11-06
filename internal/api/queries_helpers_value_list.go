package api

import (
	"net/http"
)

func (q QueryParam) ParseValueList(r *http.Request) ([]string, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return nil, err
	}

	return q.sliceValues(query)
}

func (q QueryParam) sliceValues(query string) ([]string, error) {
	values, err := QueryListSplit(query)
	if err != nil {
		return nil, err
	}

	for _, value := range values {
		err := q.CheckValue(value)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}
