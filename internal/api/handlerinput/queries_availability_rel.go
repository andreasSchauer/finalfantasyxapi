package handlerinput

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

type RelAvlParams struct {
	ParentID     int32
	Availability []database.AvailabilityType
	Repeatable   sql.NullBool
}

func (i Main[T, R, A]) QueryAvlRelGetParams(r *http.Request, parentID int32) (RelAvlParams, error) {
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

func (i Main[T, R, A]) QueryAvlRelRun[K seedtypes.Lookupable](ctx context.Context, item K, params RelAvlParams, dbQuery DBQueryRelAvl) ([]A, error) {
	dbIDs, err := dbQuery(ctx, params)
	if err != nil {
		return nil, i.ErrorDB(item, err)
	}

	resources := i.ToAPIResources(dbIDs)
	return resources, nil
}
