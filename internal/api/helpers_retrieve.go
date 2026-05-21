package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMultipleAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], name string) (L, error) {
	var zeroType L

	dbIDs, err := i.getMultipleQuery(r.Context(), name)
	if err != nil {
		return zeroType, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'.", i.resourceType, name), err)
	}

	return idsToAPIResourceList(cfg, r, i, dbIDs)
}

func retrieveAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) ([]A, error) {
	dbIDs, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return idsToAPIResources(cfg, i, dbIDs), nil
}

func filterAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], resources []A, filteredLists []filteredResList[A]) (L, error) {
	var zeroType L
	filteredRes := resources

	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return zeroType, filtered.err
		}
		filteredRes = getSharedResources(filteredRes, filtered.resources)
	}

	if !i.avlParams.IsZero() {
		var err error
		filteredRes, err = filterAvlMonsters(cfg, r, i, filteredRes)
		if err != nil {
			return zeroType, err
		}
	}

	flip, err := parseBooleanQuery(r, i.queryLookup["flip"])
	if errIsNotEmptyQuery(err) {
		return zeroType, err
	}

	if flip {
		filteredRes = removeResources(resources, filteredRes)
	}

	resourceList, err := i.resToListFunc(cfg, r, filteredRes)
	if err != nil {
		return zeroType, err
	}

	return resourceList, nil
}

func filterAvlMonsters[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], resources []A) ([]A, error) {
	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterMonsterIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		SourceType:   string(ViewSourceTypeMonster),
		AvlType:      string(AvlTypeSelf),
		Availability: availabilities,
	}

	locID, err := parseIdQuery(r, i.queryLookup["location"], len(cfg.l.Locations))
	if err == nil {
		params.AvlType = string(AvlTypeContext)
		params.ContextID = h.GetNullInt32(&locID)
		ct := string(ViewSourceTypeLocation)
		params.ContextType = h.GetNullString(&ct)
	}
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	subLocID, err := parseIdQuery(r, i.queryLookup["sublocation"], len(cfg.l.Sublocations))
	if err == nil {
		params.AvlType = string(AvlTypeContext)
		params.ContextID = h.GetNullInt32(&subLocID)
		ct := string(ViewSourceTypeSublocation)
		params.ContextType = h.GetNullString(&ct)
	}
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	
	areaID, err := parseIdQuery(r, i.queryLookup["area"], len(cfg.l.Areas))
	if err == nil {
		params.AvlType = string(AvlTypeArea)
		params.ContextID = h.GetNullInt32(&areaID)
		ct := string(ViewSourceTypeArea)
		params.ContextType = h.GetNullString(&ct)
	}
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterMonsterIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


type AvlType string

const (
	AvlTypeSelf    AvlType = "self"
	AvlTypeContext AvlType = "context"
	AvlTypeArea    AvlType = "area"
)

type ViewSourceType string

const (
	ViewSourceTypeMonster          ViewSourceType = "monster"
	ViewSourceTypeMonsterFormation ViewSourceType = "monster-formation"
	ViewSourceTypeLocation         ViewSourceType = "location"
	ViewSourceTypeSublocation      ViewSourceType = "sublocation"
	ViewSourceTypeArea             ViewSourceType = "area"
	ViewSourceTypeTreasure         ViewSourceType = "treasure"
	ViewSourceTypeShop             ViewSourceType = "shop"
	ViewSourceTypeQuest            ViewSourceType = "quest"
	ViewSourceTypeBlitzball        ViewSourceType = "blitzball"
	ViewSourceTypeItem        	   ViewSourceType = "item"
)

func resToIDs[A APIResource](resources []A) []int32 {
	ids := []int32{}

	for _, res := range resources {
		ids = append(ids, res.GetID())
	}

	return ids
}