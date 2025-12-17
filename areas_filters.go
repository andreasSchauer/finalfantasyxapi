package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (cfg *apiConfig) getAreasLocation(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}

func (cfg *apiConfig) getAreasSublocation(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}

func (cfg *apiConfig) getAreasStoryBased(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}

func (cfg *apiConfig) getAreasSaveSphere(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}

func (cfg *apiConfig) getAreasCompSphere(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}

func (cfg *apiConfig) getAreasDropOff(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}

func (cfg *apiConfig) getAreasChocobo(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}



func (cfg *apiConfig) getAreasCharacters(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getAreasAeons(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getAreasMonsters(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getAreasBosses(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getAreasShops(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getAreasTreasures(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getAreasSidequests(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getAreasFMVs(r *http.Request, inputAreas []LocationAPIResource) ([]LocationAPIResource, error) {
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

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}