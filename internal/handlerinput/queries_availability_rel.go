package handlerinput

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

func (i HandlerInput[T, R, A]) QueryAvlRelGetParams(r *http.Request, parentID int32) (RelAvlParams, error) {
	queryParam := i.QueryLookup[ts.QpnRelAvailability]

	availabilities, err := queryParam.ParseEnumList(r, enums.T.AvailabilityType)
	if errs.ErrExceptEmptyQuery(err) {
		return RelAvlParams{}, err
	}

	var repeatable sql.NullBool

	_, ok := i.QueryLookup[ts.QpnRelRepeatable]
	if ok {
		repeatable, err = NullBool(r, ts.QpnRelRepeatable, i.QueryLookup)
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

func (i HandlerInput[T, R, A]) QueryAvlRelRun[K seedtypes.Lookupable](ctx context.Context, item K, params RelAvlParams, dbQuery RelAvailabilityDBQuery) ([]A, error) {
	dbIDs, err := dbQuery(ctx, params)
	if err != nil {
		return nil, i.ErrorDB(item, err)
	}

	resources := i.ToAPIResources(dbIDs)
	return resources, nil
}
