package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

// query uses a list of names or ids as database input to filter for resources. alternatively, "none" can be used as input instead.
func (i handlerInput[T, R, A]) queryNameIDListNul[P seedtypes.Lookupable](cfg *Config, r *http.Request, inputIDs []int32, queryName QueryParamName, pResType ResTypeSingle, pLookup map[string]P, dbQuery DBQueryIntList) IDFilter {
	return func(ctx context.Context) ([]int32, error) {
		queryParam := i.queryLookup[queryName]
		if queryParam.replacingParamsPresent(r, i.queryLookup) {
			return inputIDs, nil
		}

		queryIDs, err := queryParam.parseNameIDList(cfg, r, pResType, pLookup)
		if queryIsEmpty(err) {
			return inputIDs, nil
		}
		if err != nil {
			return nil, err
		}

		dbIDs, err := dbQuery(ctx, queryIDs)
		if err != nil {
			return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParam, err)
		}

		return dbIDs, nil
	}
}
