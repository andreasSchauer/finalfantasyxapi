package queryparam

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
)

// ParseBool validates bool-queryParam and checks emptiness.
func (q QueryParam) ParseBool(r *http.Request) (bool, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return false, err
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return false, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid boolean value '%s' used for parameter '%s'. usage: '%s'", query, q.Name, q.Usage), err)
	}

	return b, nil
}
