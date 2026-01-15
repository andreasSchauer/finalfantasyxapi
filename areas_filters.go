package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (cfg *Config) queryAreasByItemMethod(r *http.Request, id int32) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["method"]
	query := r.URL.Query().Get(queryParam.Name)

	var resources []LocationAPIResource
	var err error

	switch query {
	case "":
		resources, err = cfg.getAreasByItem(r, id)
		if err != nil {
			return nil, err
		}
	case "monster":
		resources, err = cfg.getAreasByItemMonster(r, id)
		if err != nil {
			return nil, err
		}
	case "treasure":
		resources, err = cfg.getAreasByItemTreasure(r, id)
		if err != nil {
			return nil, err
		}
	case "shop":
		resources, err = cfg.getAreasByItemShop(r, id)
		if err != nil {
			return nil, err
		}
	case "quest":
		resources, err = cfg.getAreasByItemQuest(r, id)
		if err != nil {
			return nil, err
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid method value '%s'. allowed values: %s.", query, h.FormatStringSlice(queryParam.AllowedValues)), err)
	}

	return resources, nil
}

func (cfg *Config) getAreasByItem(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	filterFuncs := []func(*http.Request, int32) ([]LocationAPIResource, error){
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
	dbAreas, err := cfg.db.GetAreaIDsItemMonster(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by monster-item.", err)
	}

	resources := idsToLocationAPIResources(cfg, cfg.e.areas, dbAreas)

	return resources, nil
}

func (cfg *Config) getAreasByItemTreasure(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreaIDsItemTreasure(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by treasure-item.", err)
	}

	resources := idsToLocationAPIResources(cfg, cfg.e.areas, dbAreas)

	return resources, nil
}

func (cfg *Config) getAreasByItemShop(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreaIDsItemShop(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by shop-item.", err)
	}

	resources := idsToLocationAPIResources(cfg, cfg.e.areas, dbAreas)

	return resources, nil
}

func (cfg *Config) getAreasByItemQuest(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreaIDsItemQuest(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by quest-item.", err)
	}

	resources := idsToLocationAPIResources(cfg, cfg.e.areas, dbAreas)

	return resources, nil
}

func (cfg *Config) getAreasByKeyItem(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	filterFuncs := []func(*http.Request, int32) ([]LocationAPIResource, error){
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
	dbAreas, err := cfg.db.GetAreaIDsKeyItemTreasure(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by treasure-key-item.", err)
	}

	resources := idsToLocationAPIResources(cfg, cfg.e.areas, dbAreas)

	return resources, nil
}

func (cfg *Config) getAreasByKeyItemQuest(r *http.Request, itemID int32) ([]LocationAPIResource, error) {
	dbAreas, err := cfg.db.GetAreaIDsKeyItemQuest(r.Context(), itemID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by quest-key-item", err)
	}

	resources := idsToLocationAPIResources(cfg, cfg.e.areas, dbAreas)

	return resources, nil
}
