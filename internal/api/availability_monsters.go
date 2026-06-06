package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


func filterAvlMonsters(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.monsters

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterMonsterIDsByAvailability(r.Context(), database.FilterMonsterIDsByAvailabilityParams{
		Ids:            resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlMonsterFormations(cfg *Config, r *http.Request, resources []UnnamedAPIResource) ([]UnnamedAPIResource, error) {
	i := cfg.e.monsterFormations

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterMonsterFormationIDsByAvailability(r.Context(), database.FilterMonsterFormationIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}
