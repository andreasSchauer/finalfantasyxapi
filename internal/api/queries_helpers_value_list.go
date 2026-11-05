package api

import (
	"net/http"
)

func (q QueryParam) parseValueList(r *http.Request) ([]string, error) {
	query, err := q.checkEmpty(r)
	if err != nil {
		return nil, err
	}

	return q.sliceValues(query)
}

func (q QueryParam) sliceValues(query string) ([]string, error) {
	values, err := queryListSplit(query)
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
