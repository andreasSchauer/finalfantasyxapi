package handlerinput

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func (i HandlerInput[T, R, A]) VerifyQueriesGet(r *http.Request, id int32) (T, error) {
	var zeroType T

	err := VerifyQueryParams(r, i.Endpoint, i.QueryLookup, &id, nil)
	if err != nil {
		return zeroType, err
	}

	resource, err := lookup.GetResourceByID(id, i.ObjLookupID)
	if err != nil {
		return zeroType, errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s with id '%d' doesn't exist", i.ResTypeSingle, id), err)
	}

	return resource, nil
}

func (i HandlerInput[T, R, A]) VerifyQueriesRetrieve(r *http.Request) ([]int32, error) {
	err := VerifyQueryParams[int32](r, i.Endpoint, i.QueryLookup, nil, nil)
	if err != nil {
		return nil, err
	}

	dbIDs, err := i.RetrieveQuery(r.Context())
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s", i.ResTypePlural), err)
	}

	return dbIDs, nil
}
