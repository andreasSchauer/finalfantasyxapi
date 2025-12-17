package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func (cfg *Config) getAreasByItem(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	filterFuncs := []func (*http.Request, int32) ([]LocationAPIResource, error){
		cfg.getAreasByItemMonster,
		cfg.getAreasByItemTreasure,
		cfg.getAreasByItemShop,
		cfg.getAreasByItemQuest,
	}

	resources := []LocationAPIResource{}

	for _, function := range filterFuncs {
		addedResources, err := function(r, itemID)
		if err != nil {
			return nil, err
		}

		resources = combineResources(resources, addedResources)
	}

	return resources, nil
}


func (cfg *Config) getAreasByItemMonster(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreasWithItemMonster(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by monster item", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithItemMonsterRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}


func (cfg *Config) getAreasByItemTreasure(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreasWithItemTreasure(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by treasure item", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithItemTreasureRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}


func (cfg *Config) getAreasByItemShop(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreasWithItemShop(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by shop item", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithItemShopRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}


func (cfg *Config) getAreasByItemQuest(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreasWithItemQuest(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by quest item", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithItemQuestRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}


func (cfg *Config) getAreasByKeyItem(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	filterFuncs := []func (*http.Request, int32) ([]LocationAPIResource, error){
		cfg.getAreasByKeyItemTreasure,
		cfg.getAreasByKeyItemQuest,
	}

	resources := []LocationAPIResource{}

	for _, function := range filterFuncs {
		addedResources, err := function(r, itemID)
		if err != nil {
			return nil, err
		}

		resources = combineResources(resources, addedResources)
	}

	return resources, nil
}


func (cfg *Config) getAreasByKeyItemTreasure(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreasWithKeyItemTreasure(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by treasure key item", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithKeyItemTreasureRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}


func (cfg *Config) getAreasByKeyItemQuest(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreasWithKeyItemQuest(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by quest key item", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithKeyItemQuestRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}