package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)


// checks for default values, special values, validity, and range validity of an integer-based non-id query. if the query doesn't use defaults, special vals, or ranges, they are simply ignored. also checks for emptiness and replaces empty inputs with default values, if they exist.
func parseIntQuery(r *http.Request, queryParam QueryType) (int, error) {
	query := r.URL.Query().Get(queryParam.Name)

	val, err := checkQueryInt(queryParam, query)
	if errors.Is(err, errEmptyQuery) {
		return 0, errEmptyQuery
	}
	if err != nil {
		return 0, err
	}

	return val, nil
}

// checks if query-int is valid and within allowed range. replaces special inputs with their corresponding value and replaces empty inputs with default values.
func checkQueryInt(queryParam QueryType, queryVal string) (int, error) {
	defaultVal, err := checkIntQueryDefaultVal(queryParam, queryVal)
	if errors.Is(err, errEmptyQuery) {
		return 0, errEmptyQuery
	}
	if !errors.Is(err, errNoDefaultVal) {
		return defaultVal, nil
	}

	specialVal, err := checkIntQuerySpecialVals(queryParam, queryVal)
	if !errors.Is(err, errNoSpecialInput) {
		return specialVal, nil
	}

	val, err := checkIntQueryAllowedRange(queryParam, queryVal)
	if err != nil && !errors.Is(err, errNoIntRange) {
		return 0, err
	}

	return val, nil
}

// checks, if an int-queryParam uses a default value. if the query is empty, it returns the default val if it exists, and errEmptyQuery if not. returns errNoDefaultVal, if the query isn't empty.
func checkIntQueryDefaultVal(queryParam QueryType, query string) (int, error) {
	if query == "" {
		if queryParam.DefaultVal == nil {
			return 0, errEmptyQuery
		}
		return *queryParam.DefaultVal, nil
	}

	return 0, errNoDefaultVal
}

// checks, if an int-queryParam uses special inputs. returns errNoSpecialInput, if it doesn't, and replaces the query value with the special input value, if it does.
func checkIntQuerySpecialVals(queryParam QueryType, query string) (int, error) {
	if queryParam.SpecialInputs == nil {
		return 0, errNoSpecialInput
	}

	for _, input := range queryParam.SpecialInputs {
		if query == input.Key {
			return input.Val, nil
		}
	}

	return 0, errNoSpecialInput
}

// checks if an int-queryParam input is valid, and within the allowed range. returns errNoIntRange, if the parameter doesn't use an integer range.
func checkIntQueryAllowedRange(queryParam QueryType, query string) (int, error) {
	val, err := strconv.Atoi(query)
	if err != nil {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. usage: '%s'.", query, queryParam.Name, queryParam.Usage), err)
	}

	intRange := queryParam.AllowedIntRange
	if intRange == nil {
		return val, errNoIntRange
	}

	min := intRange[0]
	max := intRange[1]

	if val > max || val < min {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%d' used for parameter '%s'. value must be an integer ranging from %d to %d.", val, queryParam.Name, min, max), nil)
	}

	return val, nil
}