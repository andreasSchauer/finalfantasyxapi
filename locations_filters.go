package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getLocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.locations
	resourceType := cfg.e.items.resourceType
	queryParam := i.queryLookup["method"]
	query := r.URL.Query().Get(queryParam.Name)

	var resources []NamedAPIResource
	var err error

	switch query {
	case "":
		resources, err = getLocationsByItemAllMethods(cfg, r, i, id, resourceType)
		if err != nil {
			return nil, err
		}
	case "monster":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromMonster)
		if err != nil {
			return nil, err
		}
	case "treasure":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromTreasure)
		if err != nil {
			return nil, err
		}
	case "shop":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromShop)
		if err != nil {
			return nil, err
		}
	case "quest":
		resources, err = filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromQuest)
		if err != nil {
			return nil, err
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid method value '%s'. allowed values: %s.", query, h.FormatStringSlice(queryParam.AllowedValues)), err)
	}

	return resources, nil
}


func getLocationsByItemAllMethods(cfg *Config, r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList], id int32, resourceType string) ([]NamedAPIResource, error) {
	filteredLists := []filteredResList[NamedAPIResource]{
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromMonster)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromTreasure)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromShop)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromQuest)),
	}

	resources, err := combineFilteredAPIResources(filteredLists)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func getLocationsByKeyItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.sublocations
	resourceType := cfg.e.keyItems.resourceType

	filteredLists := []filteredResList[NamedAPIResource]{
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithKeyItemFromTreasure)),
		frl(filterResourcesDB(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithKeyItemFromQuest)),
	}

	resources, err := combineFilteredAPIResources(filteredLists)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
