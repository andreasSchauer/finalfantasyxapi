package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getAreasByItem(cfg *Config, r *http.Request, id int32) ([]AreaAPIResource, error) {
	i := cfg.e.areas
	resourceType := cfg.e.items.resourceType
	queryParam := i.queryLookup["method"]
	query := r.URL.Query().Get(queryParam.Name)

	var resources []AreaAPIResource
	var err error

	switch query {
	case "":
		resources, err = getAreasByItemAllMethods(cfg, r, i, id, resourceType)
		if err != nil {
			return nil, err
		}
	case "monster":
		resources, err = getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithItemFromMonster)
		if err != nil {
			return nil, err
		}
	case "treasure":
		resources, err = getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithItemFromTreasure)
		if err != nil {
			return nil, err
		}
	case "shop":
		resources, err = getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithItemFromShop)
		if err != nil {
			return nil, err
		}
	case "quest":
		resources, err = getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithItemFromQuest)
		if err != nil {
			return nil, err
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value '%s' used for 'method'. allowed values: %s.", query, h.FormatStringSlice(queryParam.AllowedValues)), err)
	}

	return resources, nil
}

func getAreasByItemAllMethods(cfg *Config, r *http.Request, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList], id int32, resourceType string) ([]AreaAPIResource, error) {
	filteredLists := []filteredResList[AreaAPIResource]{
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithItemFromMonster)),
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithItemFromTreasure)),
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithItemFromShop)),
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithItemFromQuest)),
	}

	resources, err := combineFilteredAPIResources(filteredLists)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func getAreasByKeyItem(cfg *Config, r *http.Request, id int32) ([]AreaAPIResource, error) {
	i := cfg.e.areas
	resourceType := cfg.e.keyItems.resourceType

	filteredLists := []filteredResList[AreaAPIResource]{
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithKeyItemFromTreasure)),
		frl(getResourcesDbID(cfg, r, i, id, resourceType, cfg.db.GetAreaIDsWithKeyItemFromQuest)),
	}

	resources, err := combineFilteredAPIResources(filteredLists)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
