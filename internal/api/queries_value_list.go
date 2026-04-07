package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func filterByValues[T h.HasID, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], query string, queryParam QueryParam, dbQueryMap map[string]DbQueryNoInput) ([]int32, error) {
	values := querySplit(query, ",")
	valueMap := make(map[string]bool)

	filteredLists := []filteredIdList{}

	for _, value := range values {
		dbQuery, ok := dbQueryMap[value]
		if !ok {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for parameter '%s'. allowed values: %s.", query, queryParam.Name, h.FormatStringSlice(queryParam.AllowedValues)), nil)
		}

		if valueMap[value] {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("duplicate use of value '%s' for parameter '%s'. each value can only be used once.", value, queryParam.Name), nil)
		}

		dbIDs, err := dbQuery(r.Context())
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by value '%s'.", i.resourceType, value), err)
		}
		filteredList := fidl(dbIDs, err)
		filteredLists = append(filteredLists, filteredList)

		valueMap[value] = true
	}

	return filterIdSlices(filteredLists)
}