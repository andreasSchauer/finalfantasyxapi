package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type RelAvlParams struct {
	ParentID     int32
	Availability []database.AvailabilityType
	Repeatable   sql.NullBool
}

func (i handlerInput[T, R, A]) queryAvlRelGetParams(r *http.Request, parentID int32) (RelAvlParams, error) {
	queryParam := i.queryLookup[ts.QpnRelAvailability]

	availabilities, err := queryParam.ParseEnumList(r, enums.T.AvailabilityType)
	if errs.ErrExceptEmptyQuery(err) {
		return RelAvlParams{}, err
	}

	var repeatable sql.NullBool

	_, ok := i.queryLookup[ts.QpnRelRepeatable]
	if ok {
		repeatable, err = nullBool(r, ts.QpnRelRepeatable, i.queryLookup)
		if errs.ErrExceptEmptyQuery(err) {
			return RelAvlParams{}, err
		}
	}

	availabilityParams := RelAvlParams{
		ParentID:     parentID,
		Availability: h.SliceOrNil(availabilities),
		Repeatable:   repeatable,
	}

	return availabilityParams, nil
}

func (i handlerInput[T, R, A]) queryAvlRelRun[K seedtypes.Lookupable](ctx context.Context, item K, params RelAvlParams, dbQuery RelAvailabilityDBQuery) ([]A, error) {
	dbIDs, err := dbQuery(ctx, params)
	if err != nil {
		return nil, i.ErrorDB(item, err)
	}

	resources := i.toAPIResources(dbIDs)
	return resources, nil
}
