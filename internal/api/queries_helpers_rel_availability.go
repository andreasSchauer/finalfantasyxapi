package api

import (
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type RelAvailabilityParams struct {
	ParentID     int32
	Availability []database.AvailabilityType
	Repeatable   sql.NullBool
}


func getRelAvailabilityParams[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], parentID int32) (RelAvailabilityParams, error) {
	queryParam := i.queryLookup["rel_availability"]

	availabilities, err := parseEnumListQuery(cfg, r, i.endpoint, queryParam, cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return RelAvailabilityParams{}, err
	}

	repeatable, err := getQueryBoolPtr(r, "repeatable", cfg.e.monsters.queryLookup)
	if err != nil {
		return RelAvailabilityParams{}, err
	}

	availabilityParams := RelAvailabilityParams{
		ParentID:     parentID,
		Availability: availabilities,
		Repeatable:   h.GetNullBool(repeatable),
	}

	return availabilityParams, nil
}


func runRelAvailabilityQuery[T, K seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], item K, params RelAvailabilityParams, dbQuery RelAvailabilityDbQuery) ([]A, error) {
	dbIDs, err := dbQuery(r.Context(), params)
	if err != nil {
		return nil, newHTTPErrorDB(i.resourceType, item, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}