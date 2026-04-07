package api

import "net/http"


func getMasterItemsByMethod(cfg *Config, r *http.Request, query string, queryParam QueryParam) ([]int32, error) {
	return filterByValues(r, cfg.e.allItems, query, queryParam, map[string]DbQueryNoInput{
		"monster": 	cfg.db.GetMasterItemIDsMonster,
		"treasure": cfg.db.GetMasterItemIDsTreasure,
		"shop": 	cfg.db.GetMasterItemIDsShop,
		"quest": 	cfg.db.GetMasterItemIDsQuest,
	})
}

func getItemsByMethod(cfg *Config, r *http.Request, query string, queryParam QueryParam) ([]int32, error) {
	return filterByValues(r, cfg.e.items, query, queryParam, map[string]DbQueryNoInput{
		"monster": 	cfg.db.GetItemIDsMonster,
		"treasure": cfg.db.GetItemIDsTreasure,
		"shop": 	cfg.db.GetItemIDsShop,
		"quest": 	cfg.db.GetItemIDsQuest,
	})
}

func getKeyItemsByMethod(cfg *Config, r *http.Request, query string, queryParam QueryParam) ([]int32, error) {
	return filterByValues(r, cfg.e.keyItems, query, queryParam, map[string]DbQueryNoInput{
		"treasure": cfg.db.GetKeyItemIDsTreasure,
		"quest": 	cfg.db.GetKeyItemIDsQuest,
	})
}