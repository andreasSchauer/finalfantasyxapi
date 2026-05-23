package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


func getAreasByItem(cfg *Config, r *http.Request, id int32) ([]AreaAPIResource, error) {
	i := cfg.e.areas
	queryParamMethod := i.queryLookup["method"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethod)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetAreaIDsWithItemFromMethod(r.Context(), database.GetAreaIDsWithItemFromMethodParams{
		ID: 	id,
		Method: methods,
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethod, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getSublocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	return filterByIdAndValues(cfg, r, cfg.e.sublocations, id, "method", cfg.e.items.resourceType, map[string]DbQueryIntMany{
		"monster":  cfg.db.GetSublocationIDsWithItemFromMonster,
		"treasure": cfg.db.GetSublocationIDsWithItemFromTreasure,
		"shop":     cfg.db.GetSublocationIDsWithItemFromShop,
		"quest":    cfg.db.GetSublocationIDsWithItemFromQuest,
	})
}

func getLocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	return filterByIdAndValues(cfg, r, cfg.e.locations, id, "method", cfg.e.items.resourceType, map[string]DbQueryIntMany{
		"monster":  cfg.db.GetLocationIDsWithItemFromMonster,
		"treasure": cfg.db.GetLocationIDsWithItemFromTreasure,
		"shop":     cfg.db.GetLocationIDsWithItemFromShop,
		"quest":    cfg.db.GetLocationIDsWithItemFromQuest,
	})
}


func getSublocationsByKeyItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	return dbQueriesToApiResources(cfg, r, cfg.e.sublocations, id, cfg.e.keyItems.resourceType, map[string]DbQueryIntMany{
		"treasure": cfg.db.GetSublocationIDsWithKeyItemFromTreasure,
		"quest":    cfg.db.GetSublocationIDsWithKeyItemFromQuest,
	})
}

func getLocationsByKeyItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	return dbQueriesToApiResources(cfg, r, cfg.e.locations, id, cfg.e.keyItems.resourceType, map[string]DbQueryIntMany{
		"treasure": cfg.db.GetLocationIDsWithKeyItemFromTreasure,
		"quest":    cfg.db.GetLocationIDsWithKeyItemFromQuest,
	})
}
