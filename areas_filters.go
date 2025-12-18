package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (cfg *Config) getAreasLocation(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "location"
	location, isEmpty, err := parseUniqueNameQuery(r, queryParam, cfg.l.Locations)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetLocationAreas(r.Context(), location.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by location", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetLocationAreasRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasSublocation(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "sublocation"
	sublocation, isEmpty, err := parseUniqueNameQuery(r, queryParam, cfg.l.SubLocations)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetSublocationAreas(r.Context(), sublocation.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by sublocation", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetSublocationAreasRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasItem(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryItem := "item"
	queryMethod := r.URL.Query().Get("method")

	item, itemIsEmpty, err := parseUniqueNameQuery(r, queryItem, cfg.l.Items)
	if err != nil {
		return nil, err
	}
	if itemIsEmpty {
		if queryMethod != "" {
			return nil, newHTTPError(http.StatusBadRequest, "invalid input. method parameter must be paired with item or key-item parameter. usage: item={item}&method={monster/treasure/shop/quest}; or item={item}&method={treasure/quest}", nil)
		}
		return inputAreas, nil
	}

	var resources []LocationAPIResource

	switch queryMethod {
	case "":
		resources, err = cfg.getAreasByItem(r, item.ID)
		if err != nil {
			return nil, err
		}
	case "monster":
		resources, err = cfg.getAreasByItemMonster(r, item.ID)
		if err != nil {
			return nil, err
		}
	case "treasure":
		resources, err = cfg.getAreasByItemTreasure(r, item.ID)
		if err != nil {
			return nil, err
		}
	case "shop":
		resources, err = cfg.getAreasByItemShop(r, item.ID)
		if err != nil {
			return nil, err
		}
	case "quest":
		resources, err = cfg.getAreasByItemQuest(r, item.ID)
		if err != nil {
			return nil, err
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid method value: %s. allowed methods: monster, treasure, shop, quest.", queryMethod), err)
	}

	return resources, nil
}

func (cfg *Config) getAreasKeyItem(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "key-item"

	item, isEmpty, err := parseUniqueNameQuery(r, queryParam, cfg.l.KeyItems)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	resources, err := cfg.getAreasByKeyItem(r, item.ID)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (cfg *Config) getAreasStoryBased(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "story-based"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasStoryOnly(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve story-based areas", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasStoryOnlyRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasSaveSphere(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "save-sphere"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithSaveSphere(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with save sphere", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithSaveSphereRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasCompSphere(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "comp-sphere"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithCompSphere(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with compilation sphere", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithCompSphereRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasDropOff(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "airship"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithDropOff(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with airship drop off", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithDropOffRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasChocobo(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "chocobo"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithChocobo(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas where you can ride a chocobo", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithChocoboRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasCharacters(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "characters"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithCharacters(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with characters", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithCharactersRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}

func (cfg *Config) getAreasAeons(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "aeons"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithAeons(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with aeons", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithAeonsRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}

func (cfg *Config) getAreasMonsters(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "monsters"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithMonsters(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with monsters", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithMonstersRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}

func (cfg *Config) getAreasBosses(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "boss-fights"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithBosses(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with bosses", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithBossesRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}

func (cfg *Config) getAreasShops(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "shops"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithShops(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with shops", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithShopsRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}

func (cfg *Config) getAreasTreasures(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "treasures"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithTreasures(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with treasures", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithTreasuresRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}

func (cfg *Config) getAreasSidequests(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "sidequests"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithSidequests(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with sidequests", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithSidequestsRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}

func (cfg *Config) getAreasFMVs(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := "fmvs"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputAreas, nil
	}

	dbAreas, err := cfg.db.GetAreasWithFMVs(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with FMVs", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithFMVsRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}
