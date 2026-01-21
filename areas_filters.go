package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getAreasByItem(cfg *Config, r *http.Request, id int32) ([]LocationAPIResource, error) {
	i := cfg.e.areas
	resourceType := cfg.e.items.resourceType
	queryParam := i.queryLookup["method"]
	query := r.URL.Query().Get(queryParam.Name)

	var resources []LocationAPIResource
	var err error

	switch query {
	case "":
		resources, err = getAreasByItemAllMethods(cfg, r, i, id, resourceType)
		if err != nil {
			return nil, err
		}
	case "monster":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsItemMonster)
		if err != nil {
			return nil, err
		}
	case "treasure":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsItemTreasure)
		if err != nil {
			return nil, err
		}
	case "shop":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsItemShop)
		if err != nil {
			return nil, err
		}
	case "quest":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsItemQuest)
		if err != nil {
			return nil, err
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid method value '%s'. allowed values: %s.", query, h.FormatStringSlice(queryParam.AllowedValues)), err)
	}

	return resources, nil
}

// same pattern as retrieve endpoint
func getAreasByItemAllMethods(cfg *Config, r *http.Request, i handlerInput[seeding.Area, Area, LocationAPIResource, LocationApiResourceList], id int32, resourceType string) ([]LocationAPIResource, error) {
	filteredLists := []filteredResList[LocationAPIResource]{
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsItemMonster)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsItemTreasure)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsItemShop)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsItemQuest)),
	}

	resources, err := combineFilteredAPIResources(filteredLists)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func getAreasByKeyItem(cfg *Config, r *http.Request, id int32) ([]LocationAPIResource, error) {
	i := cfg.e.areas
	resourceType := cfg.e.keyItems.resourceType

	filteredLists := []filteredResList[LocationAPIResource]{
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsKeyItemTreasure)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsKeyItemQuest)),
	}

	resources, err := combineFilteredAPIResources(filteredLists)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
