package main

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// all boolean queries can be generalized
// if a filter isn't tied to a second parameter, these functions can be generalized
func (cfg *Config) getAreasLocation(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["location"]

	id, err := parseIDOnlyQuery(r, queryParam, len(cfg.l.Locations))
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetLocationAreas(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by location.", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetLocationAreasRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasSublocation(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["sublocation"]

	id, err := parseIDOnlyQuery(r, queryParam, len(cfg.l.SubLocations))
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetSublocationAreas(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas by sublocation.", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetSublocationAreasRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasItem(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["item"]

	id, err := parseIDOnlyQuery(r, queryParam, len(cfg.l.Items))
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	resources, err := cfg.queryAreasByItemMethod(r, id)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (cfg *Config) getAreasKeyItem(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["key-item"]

	id, err := parseIDOnlyQuery(r, queryParam, len(cfg.l.KeyItems))
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	resources, err := cfg.getAreasByKeyItem(r, id)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (cfg *Config) getAreasStoryBased(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["story-based"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasStoryOnly(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve story-based areas.", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasStoryOnlyRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasSaveSphere(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["save-sphere"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithSaveSphere(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with save sphere.", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithSaveSphereRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasCompSphere(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["comp-sphere"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithCompSphere(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with compilation sphere.", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithCompSphereRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasDropOff(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["airship"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithDropOff(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with airship drop off.", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithDropOffRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasChocobo(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["chocobo"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithChocobo(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas where you can ride a chocobo.", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithChocoboRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	return resources, nil
}

func (cfg *Config) getAreasCharacters(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
	queryParam := cfg.q.areas["characters"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithCharacters(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with characters.", err)
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
	queryParam := cfg.q.areas["aeons"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithAeons(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with aeons.", err)
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
	queryParam := cfg.q.areas["monsters"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithMonsters(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with monsters.", err)
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
	queryParam := cfg.q.areas["boss-fights"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithBosses(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with bosses.", err)
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
	queryParam := cfg.q.areas["shops"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithShops(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with shops.", err)
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
	queryParam := cfg.q.areas["treasures"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithTreasures(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with treasures.", err)
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
	queryParam := cfg.q.areas["sidequests"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithSidequests(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with sidequests.", err)
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
	queryParam := cfg.q.areas["fmvs"]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputAreas, nil
	}
	if err != nil {
		return nil, err
	}

	dbAreas, err := cfg.db.GetAreasWithFMVs(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas with fmvs.", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithFMVsRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Area, h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	return resources, nil
}
