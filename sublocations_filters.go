package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getSublocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.sublocations
	resourceType := cfg.e.items.resourceType
	queryParam := i.queryLookup["method"]
	query := r.URL.Query().Get(queryParam.Name)

	var resources []NamedAPIResource
	var err error

	switch query {
	case "":
		resources, err = getSublocationsByItemAllMethods(cfg, r, i, id, resourceType)
		if err != nil {
			return nil, err
		}
	case "monster":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithItemFromMonster)
		if err != nil {
			return nil, err
		}
	case "treasure":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithItemFromTreasure)
		if err != nil {
			return nil, err
		}
	case "shop":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithItemFromShop)
		if err != nil {
			return nil, err
		}
	case "quest":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithItemFromQuest)
		if err != nil {
			return nil, err
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for 'method'. allowed values: %s.", query, h.FormatStringSlice(queryParam.AllowedValues)), err)
	}

	return resources, nil
}

func getSublocationsByItemAllMethods(cfg *Config, r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList], id int32, resourceType string) ([]NamedAPIResource, error) {
	filteredLists := []filteredResList[NamedAPIResource]{
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithItemFromMonster)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithItemFromTreasure)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithItemFromShop)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithItemFromQuest)),
	}

	resources, err := combineFilteredAPIResources(filteredLists)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func getSublocationsByKeyItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.sublocations
	resourceType := cfg.e.keyItems.resourceType

	filteredLists := []filteredResList[NamedAPIResource]{
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithKeyItemFromTreasure)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetSublocationIDsWithKeyItemFromQuest)),
	}

	resources, err := combineFilteredAPIResources(filteredLists)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
