package handlerinput

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func (i Main[T, R, A]) VerifyQueriesGet(r *http.Request, id int32) (T, error) {
	var zeroType T

	err := verifyQueryParams(r, i.Endpoint, i.QueryLookup, &id, nil)
	if err != nil {
		return zeroType, err
	}

	resource, err := lookup.GetResourceByID(id, i.ObjLookupID)
	if err != nil {
		return zeroType, errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s with id '%d' doesn't exist", i.ResTypeSingle, id), err)
	}

	return resource, nil
}

func (i Main[T, R, A]) VerifyQueriesRetrieve(r *http.Request) ([]int32, error) {
	err := verifyQueryParams[int32](r, i.Endpoint, i.QueryLookup, nil, nil)
	if err != nil {
		return nil, err
	}

	dbIDs, err := i.RetrieveQuery(r.Context())
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s", i.ResTypePlural), err)
	}

	return dbIDs, nil
}

func (i Main[T, R, A]) VerifyIDsQuery(r *http.Request) error {
	sn := string(ts.SnSimple)
	return verifyQueryParams[int32](r, i.Endpoint, i.QueryLookup, nil, &sn)
}

func (i Main[T, R, A]) VerifyQueriesAltList(r *http.Request, sectionName ts.SectionName) error {
	sn := string(sectionName)
	return verifyQueryParamsAltList(r, i.Endpoint, i.QueryLookup, &sn)
}
