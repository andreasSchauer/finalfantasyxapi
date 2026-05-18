package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AvailabilityParams struct {
	ParentID     int32
	Availability []database.AvailabilityType
	Method       sql.NullString
}

func getAvailabilityParams[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], parentID int32) (AvailabilityParams, error) {
	queryParam := i.queryLookup["availability"]

	availabilities, err := parseEnumListQuery(cfg, r, i.endpoint, queryParam, cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return AvailabilityParams{}, err
	}

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if err != nil {
		return AvailabilityParams{}, err
	}

	availabilityParams := AvailabilityParams{
		ParentID:     parentID,
		Availability: availabilities,
		Method:       h.GetNullString(method),
	}

	return availabilityParams, nil
}

func runAvailabilityQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, pResType string, dbQuery AvailabilityDbQuery) ([]A, error) {
	params, err := getAvailabilityParams(cfg, r, i, id)
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss.", i.resourceType), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}
