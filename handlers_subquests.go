package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type Subquest struct {
	ID          		int32         		`json:"id"`
	Name        		string        		`json:"name"`
	ParentSidequest		NamedAPIResource	`json:"parent_sidequest"`
	Completions			[]QuestCompletion	`json:"completions"`
}


func (cfg *Config) HandleSubquests(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.subquests

	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(w, r, i)
		return

	case 1:
		handleEndpointIDOnly(cfg, w, r, i, segments)
		return

	case 2:
		handleEndpointSubsections(cfg, w, r, i, segments)
		return
		
	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: '/api/%s/{id}'.", i.endpoint), nil)
		return
	}
}


func (cfg *Config) getSubquest(r *http.Request, i handlerInput[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList], id int32) (Subquest, error) {
	subquest, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Subquest{}, err
	}

	

	response := Subquest{
		ID:         		subquest.ID,
		Name:       		subquest.Name,
		ParentSidequest: 	idToNamedAPIResource(cfg, cfg.e.sidequests, subquest.SidequestID),
		Completions: 		convertObjSlice(cfg, subquest.Completions, convertQuestCompletion),
	}

	return response, nil
}


func (cfg *Config) retrieveSubquests(r *http.Request, i handlerInput[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return i.resToListFunc(cfg, r, resources)
}