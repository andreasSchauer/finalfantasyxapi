package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Sidequest struct {
	ID         int32              `json:"id"`
	Name       string             `json:"name"`
	Completion *QuestCompletion   `json:"completion"`
	Subquests  []NamedAPIResource `json:"subquests"`
}

type QuestCompletion struct {
	Areas  []CompletionArea `json:"areas"`
	Reward ItemAmount       `json:"reward"`
}

func convertQuestCompletion(cfg *Config, qc seeding.QuestCompletion) QuestCompletion {
	return QuestCompletion{
		Areas:  convertObjSlice(cfg, qc.Areas, convertCompletionArea),
		Reward: convertItemAmount(cfg, qc.Reward),
	}
}

type CompletionArea struct {
	Area  AreaAPIResource `json:"area"`
	Notes *string         `json:"notes,omitempty"`
}

func (ca CompletionArea) GetAPIResource() APIResource {
	return ca.Area
}

func convertCompletionArea(cfg *Config, cl seeding.CompletionArea) CompletionArea {
	return CompletionArea{
		Area:  locAreaToAreaAPIResource(cfg, cfg.e.areas, cl.LocationArea),
		Notes: cl.Notes,
	}
}

func (cfg *Config) HandleSidequests(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.sidequests

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
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", getUsageString(i)), nil)
		return
	}
}

func (cfg *Config) getSidequest(r *http.Request, i handlerInput[seeding.Sidequest, Sidequest, NamedAPIResource, NamedApiResourceList], id int32) (Sidequest, error) {
	sidequest, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Sidequest{}, err
	}

	subquests, err := getResourcesDB(cfg, r, cfg.e.subquests, sidequest, cfg.db.GetSidequestSubquestIDs)
	if err != nil {
		return Sidequest{}, err
	}

	response := Sidequest{
		ID:         sidequest.ID,
		Name:       sidequest.Name,
		Completion: convertObjPtr(cfg, sidequest.Completion, convertQuestCompletion),
		Subquests:  subquests,
	}

	return response, nil
}

func (cfg *Config) retrieveSidequests(r *http.Request, i handlerInput[seeding.Sidequest, Sidequest, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return i.resToListFunc(cfg, r, resources)
}
