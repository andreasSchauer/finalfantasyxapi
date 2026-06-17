package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func filterAvlAutoAbilities(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.autoAbilities

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	reqItem, err := parseBooleanQuery(r, i.queryLookup["req_item"])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	charID, err := getQueryIdPtr(r, cfg.e.characters, "character", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup["methods"])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterAutoAbilityIDsByAvailability(r.Context(), database.FilterAutoAbilityIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
		CharacterID: 	h.GetNullInt32(charID),
		Methods: 		methods,
		ReqItem:      	reqItem,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}
