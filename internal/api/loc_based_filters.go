package api

import (
	"net/http"
)


func getAreasByItem(cfg *Config, r *http.Request, id int32) ([]AreaAPIResource, error) {
	return filterByIdAndValues(cfg, r, cfg.e.areas, id, "method", cfg.e.items.resourceType, map[string]DbQueryIntMany{
		"monster": 	cfg.db.GetAreaIDsWithItemFromMonster,
		"treasure": cfg.db.GetAreaIDsWithItemFromTreasure,
		"shop": 	cfg.db.GetAreaIDsWithItemFromShop,
		"quest": 	cfg.db.GetAreaIDsWithItemFromQuest,
	})
}

func getSublocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	return filterByIdAndValues(cfg, r, cfg.e.sublocations, id, "method", cfg.e.items.resourceType, map[string]DbQueryIntMany{
		"monster": 	cfg.db.GetSublocationIDsWithItemFromMonster,
		"treasure": cfg.db.GetSublocationIDsWithItemFromTreasure,
		"shop": 	cfg.db.GetSublocationIDsWithItemFromShop,
		"quest": 	cfg.db.GetSublocationIDsWithItemFromQuest,
	})
}

func getLocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	return filterByIdAndValues(cfg, r, cfg.e.locations, id, "method", cfg.e.items.resourceType, map[string]DbQueryIntMany{
		"monster": 	cfg.db.GetLocationIDsWithItemFromMonster,
		"treasure": cfg.db.GetLocationIDsWithItemFromTreasure,
		"shop": 	cfg.db.GetLocationIDsWithItemFromShop,
		"quest": 	cfg.db.GetLocationIDsWithItemFromQuest,
	})
}


func getAreasByKeyItem(cfg *Config, r *http.Request, id int32) ([]AreaAPIResource, error) {
	return dbQueriesToApiResources(cfg, r, cfg.e.areas, id, cfg.e.keyItems.resourceType, map[string]DbQueryIntMany{
		"treasure": cfg.db.GetAreaIDsWithKeyItemFromTreasure,
		"quest": 	cfg.db.GetAreaIDsWithKeyItemFromQuest,
	})
}

func getSublocationsByKeyItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	return dbQueriesToApiResources(cfg, r, cfg.e.sublocations, id, cfg.e.keyItems.resourceType, map[string]DbQueryIntMany{
		"treasure": cfg.db.GetSublocationIDsWithKeyItemFromTreasure,
		"quest": 	cfg.db.GetSublocationIDsWithKeyItemFromQuest,
	})
}

func getLocationsByKeyItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	return dbQueriesToApiResources(cfg, r, cfg.e.locations, id, cfg.e.keyItems.resourceType, map[string]DbQueryIntMany{
		"treasure": cfg.db.GetLocationIDsWithKeyItemFromTreasure,
		"quest": 	cfg.db.GetLocationIDsWithKeyItemFromQuest,
	})
}