package api

import (
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
	queryParam := i.queryLookup[qpnRelAvailability]

	availabilities, err := parseEnumListQuery(cfg, r, i.endpoint, queryParam, cfg.t.AvailabilityType)
	if errExceptEmptyQuery(err) {
		return RelAvlParams{}, err
	}

	var repeatable *bool

	_, ok := i.queryLookup[qpnRelRepeatable]
	if ok {
		repeatable, err = getQueryBoolPtr(r, qpnRelRepeatable, i.queryLookup)
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

func runRelAvailabilityQuery[T, K seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], item K, params RelAvlParams, dbQuery RelAvailabilityDbQuery) ([]A, error) {
	dbIDs, err := dbQuery(r.Context(), params)
	if err != nil {
		return nil, newHTTPErrorDB(i.resTypePlural, item, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}
