package api

import "net/http"

func getMasterItemsByMethod(cfg *Config, r *http.Request, query string, queryParam QueryParam) ([]int32, error) {
	return filterByValues2(cfg, r, cfg.e.allItems, query, queryParam, cfg.db.GetMasterItemIDsByMethods)
}

func getItemsByMethod(cfg *Config, r *http.Request, query string, queryParam QueryParam) ([]int32, error) {
	return filterByValues2(cfg, r, cfg.e.items, query, queryParam, cfg.db.GetItemIDsByMethods)
}

func getKeyItemsByMethod(cfg *Config, r *http.Request, query string, queryParam QueryParam) ([]int32, error) {
	return filterByValues2(cfg, r, cfg.e.keyItems, query, queryParam, cfg.db.GetKeyItemIDsByMethods)
}
