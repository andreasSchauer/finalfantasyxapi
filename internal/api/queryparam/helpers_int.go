package queryparam

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
)

// ParseInt checks for default values, special values, validity, and range validity of an integer-based non-id query. if the query doesn't use defaults, special vals, or ranges, they are simply ignored. also checks for emptiness and replaces empty inputs with default values, if they exist.
func (q QueryParam) ParseInt(r *http.Request) (int, error) {
	query := r.URL.Query().Get(string(q.Name))

	val, err := q.CheckInt(query)
	if errs.QueryIsEmpty(err) {
		return 0, errs.ErrEmptyQuery
	}
	if err != nil {
		return 0, err
	}

	return val, nil
}

// CheckInt checks if query-int is valid and within allowed range. replaces special inputs with their corresponding value and replaces empty inputs with default values.
func (q QueryParam) CheckInt(queryVal string) (int, error) {
	defaultVal, err := q.checkIntDefaultVal(queryVal)
	if errs.QueryIsEmpty(err) {
		return 0, errs.ErrEmptyQuery
	}
	if !errors.Is(err, errs.ErrNoDefaultVal) {
		return defaultVal, nil
	}

	specialVal, err := q.checkIntSpecialVals(queryVal)
	if !errors.Is(err, errs.ErrNoSpecialInput) {
		return specialVal, nil
	}

	val, err := q.checkIntAllowedRange(queryVal)
	if err != nil && !errors.Is(err, errs.ErrNoIntRange) {
		return 0, err
	}

	return val, nil
}

// checks, if an int-queryParam uses a default value. if the query is empty, it returns the default val if it exists, and errEmptyQuery if not. returns errNoDefaultVal, if the query isn't empty.
func (q QueryParam) checkIntDefaultVal(query string) (int, error) {
	if query == "" {
		if q.DefaultVal == nil {
			return 0, errs.ErrEmptyQuery
		}
		return *q.DefaultVal, nil
	}

	return 0, errs.ErrNoDefaultVal
}

// checks, if an int-queryParam uses special inputs. returns errNoSpecialInput, if it doesn't, and replaces the query value with the special input value, if it does.
func (q QueryParam) checkIntSpecialVals(query string) (int, error) {
	if q.SpecialInputs == nil {
		return 0, errs.ErrNoSpecialInput
	}

	for _, input := range q.SpecialInputs {
		if strings.ToLower(query) == string(input.Key) {
			return input.Val, nil
		}
	}

	return 0, errs.ErrNoSpecialInput
}

// checks if an int-queryParam input is valid, and within the allowed range. returns errNoIntRange, if the parameter doesn't use an integer range.
func (q QueryParam) checkIntAllowedRange(query string) (int, error) {
	val, err := strconv.Atoi(query)
	if err != nil {
		return 0, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. usage: '%s'", query, q.Name, q.Usage), err)
	}

	intRange := q.AllowedIntRange
	if intRange == nil {
		return val, errs.ErrNoIntRange
	}

	min := intRange[0]
	max := intRange[1]

	if val > max || val < min {
		return 0, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%d' used for parameter '%s'. value must be an integer ranging from %d to %d", val, q.Name, min, max), nil)
	}

	return val, nil
}
