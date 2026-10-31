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

func getRelAvailabilityParams[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], parentID int32) (RelAvlParams, error) {
	queryParam := i.queryLookup[QpnRelAvailability]

	availabilities, err := parseEnumListQuery(cfg, r, i.Endpoint, queryParam, cfg.T.AvailabilityType)
	if errExceptEmptyQuery(err) {
		return RelAvlParams{}, err
	}

	var repeatable *bool

	_, ok := i.queryLookup[QpnRelRepeatable]
	if ok {
		repeatable, err = getQueryBoolPtr(r, QpnRelRepeatable, i.queryLookup)
		if errExceptEmptyQuery(err) {
			return RelAvlParams{}, err
		}
	}

	availabilityParams := RelAvlParams{
		ParentID:     parentID,
		Availability: h.SliceOrNil(availabilities),
		Repeatable:   h.GetNullBool(repeatable),
	}

	return availabilityParams, nil
}

func runRelAvailabilityQuery[T, K seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, ctx context.Context, i handlerInput[T, R, A, L], item K, params RelAvlParams, dbQuery RelAvailabilityDBQuery) ([]A, error) {
	dbIDs, err := dbQuery(ctx, params)
	if err != nil {
		return nil, newHTTPErrorDB(i.resTypePlural, item, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}
