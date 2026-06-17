package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


func filterAvlShops(cfg *Config, r *http.Request, resources []UnnamedAPIResource) ([]UnnamedAPIResource, error) {
	i := cfg.e.shops

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	sources, err := getShopSources(cfg, r, i)
	if err != nil {
		return nil, err
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterShopIDsByAvailability(r.Context(), database.FilterShopIDsByAvailabilityParams{
		Ids:          		resToIDs(resources),
		Availability: 		avlParams.availabilities,
		AvlType:      		sources.AvlType,
		LocContextID: 		locContext.ID,
		LocContextType: 	locContext.Type,
		RequiredSources:    sources.RequiredSources,
		ExcludedSources: 	sources.ExcludedSources,
		AutoAbilityID: 		sources.AutoAbilityID,
		CharacterID: 		sources.CharacterID,
		EmptySlots: 		sources.EmptySlots,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlTreasures(cfg *Config, r *http.Request, resources []UnnamedAPIResource) ([]UnnamedAPIResource, error) {
	i := cfg.e.treasures

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterTreasureIDsByAvailability(r.Context(), database.FilterTreasureIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability: 	avlParams.availabilities,
		PreAirship:		avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlQuests(cfg *Config, r *http.Request, resources []QuestAPIResource) ([]QuestAPIResource, error) {
	i := cfg.e.quests

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterQuestIDsByAvailability(r.Context(), database.FilterQuestIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlSidequests(cfg *Config, r *http.Request, resources []QuestAPIResource) ([]QuestAPIResource, error) {
	i := cfg.e.sidequests

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterSidequestIDsByAvailability(r.Context(), database.FilterSidequestIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlSubquests(cfg *Config, r *http.Request, resources []QuestAPIResource) ([]QuestAPIResource, error) {
	i := cfg.e.subquests

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterSubquestIDsByAvailability(r.Context(), database.FilterSubquestIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}