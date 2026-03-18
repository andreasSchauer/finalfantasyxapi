package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getLocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.locations
	resourceType := cfg.e.items.resourceType
	queryParamMethod := i.queryLookup["method"]
	queryMethod := r.URL.Query().Get(queryParamMethod.Name)

	switch queryMethod {
	case "":
		return getLocationsByItemAllMethods(cfg, r, i, id, resourceType)
		
	case "monster":
		return getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromMonster)
		
	case "treasure":
		return getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromTreasure)
		
	case "shop":
		return getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromShop)
		
	case "quest":
		return getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromQuest)
		
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for 'method'. allowed values: %s.", queryMethod, h.FormatStringSlice(queryParamMethod.AllowedValues)), nil)
	}
}

func getLocationsByItemAllMethods(cfg *Config, r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList], id int32, resourceType string) ([]NamedAPIResource, error) {
	filteredLists := []filteredResList[NamedAPIResource]{
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromMonster)),
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromTreasure)),
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromShop)),
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithItemFromQuest)),
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
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithKeyItemFromTreasure)),
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetLocationIDsWithKeyItemFromQuest)),
	}

	resources, err := combineFilteredAPIResources(filteredLists)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
