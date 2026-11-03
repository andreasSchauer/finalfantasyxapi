package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type RelAvlParams struct {
	ParentID     int32
	Availability []database.AvailabilityType
	Repeatable   sql.NullBool
}

func (i handlerInput[T, R, A]) queryAvlRelGetParams(cfg *Config, r *http.Request, parentID int32) (RelAvlParams, error) {
	queryParam := i.queryLookup[QpnRelAvailability]

	availabilities, err := queryParam.parseEnumList(cfg, r, cfg.T.AvailabilityType)
	if errExceptEmptyQuery(err) {
		return RelAvlParams{}, err
	}

	var repeatable sql.NullBool

	_, ok := i.queryLookup[QpnRelRepeatable]
	if ok {
		repeatable, err = nullBool(r, QpnRelRepeatable, i.queryLookup)
		if errExceptEmptyQuery(err) {
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

func (i handlerInput[T, R, A]) queryAvlRelRun[K seeding.Lookupable](cfg *Config, ctx context.Context, item K, params RelAvlParams, dbQuery RelAvailabilityDBQuery) ([]A, error) {
	dbIDs, err := dbQuery(ctx, params)
	if err != nil {
		return nil, newHTTPErrorDB(i.resTypePlural, item, err)
	}

	resources := i.toResources(cfg, dbIDs)
	return resources, nil
}
