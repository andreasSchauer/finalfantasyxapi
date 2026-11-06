package queryparam

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// ParseTyped validates typed-QueryParam and checks emptiness
func (q QueryParam) ParseTyped(r *http.Request) (string, string, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return "", "", err
	}

	resType, unitStr, found := strings.Cut(query, ":")
	if !found {
		return "", "", errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': '%s'. usage: '%s'", q.Name, query, q.Usage), nil)
	}

	if !slices.Contains(q.AllowedResTypes, resType) {
		return "", "", errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resource type '%s' for parameter '%s'. supported resource types: %s", resType, q.Name, h.FormatStringSlice(q.AllowedResTypes)), nil)
	}

	return resType, unitStr, nil
}
