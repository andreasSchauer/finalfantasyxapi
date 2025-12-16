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
		return area.Location, h.NullStringToVal(area.Sublocation), h.NullStringToVal(area.Area), h.NullInt32ToPtr(area.Version)
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
		return area.Location, h.NullStringToVal(area.Sublocation), h.NullStringToVal(area.Area), h.NullInt32ToPtr(area.Version)
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
		return area.Location, h.NullStringToVal(area.Sublocation), h.NullStringToVal(area.Area), h.NullInt32ToPtr(area.Version)
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
		return area.Location, h.NullStringToVal(area.Sublocation), h.NullStringToVal(area.Area), h.NullInt32ToPtr(area.Version)
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
		return area.Location, h.NullStringToVal(area.Sublocation), h.NullStringToVal(area.Area), h.NullInt32ToPtr(area.Version)
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
		return area.Location, h.NullStringToVal(area.Sublocation), h.NullStringToVal(area.Area), h.NullInt32ToPtr(area.Version)
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
		return area.Location, h.NullStringToVal(area.Sublocation), h.NullStringToVal(area.Area), h.NullInt32ToPtr(area.Version)
	})

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
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas where you can ride a chocobo", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasWithBossesRow) (string, string, string, *int32) {
		return h.NullStringToVal(area.Location), h.NullStringToVal(area.Sublocation), h.NullStringToVal(area.Area), h.NullInt32ToPtr(area.Version)
	})

	if !b {
		resources = removeResources(inputAreas, resources)
	}

	sharedResources := getSharedResources(inputAreas, resources)

	return sharedResources, nil
}
