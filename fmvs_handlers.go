package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type FMV struct {
	ID          		int32         		`json:"id"`
	Name        		string        		`json:"name"`
	Translation         *string      		`json:"translation,omitempty"`
	CutsceneDescription string       		`json:"cutscene_description"`
	Area				LocationAPIResource	`json:"area"`
	Song				*NamedAPIResource	`json:"song"`
}


func (cfg *Config) HandleFMVs(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.fmvs

	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(w, r, i)
		return

	case 1:
		handleEndpointNameOrID(cfg, w, r, i, segments)
		return

	case 2:
		handleEndpointSubsections(cfg, w, r, i, segments)
		return
		
	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: '/api/%s/{id}'.", i.endpoint), nil)
		return
	}
}


func (cfg *Config) getFMV(r *http.Request, i handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList], id int32) (FMV, error) {
	fmv, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return FMV{}, err
	}

	response := FMV{
		ID:          			fmv.ID,
		Name:        			fmv.Name,
		Translation: fmv.Translation,
		CutsceneDescription: fmv.CutsceneDescription,
		Area: locAreaToLocationAPIResource(cfg, cfg.e.areas, fmv.LocationArea),
		Song: namePtrToNamedAPIResPtr(cfg, cfg.e.songs, fmv.SongName, nil),
	}

	return response, nil
}


func (cfg *Config) retrieveFMVs(r *http.Request, i handlerInput[seeding.FMV, FMV, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	filteredLists := []filteredResList[NamedAPIResource]{
		frl(idOnlyQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationFmvIDs)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}