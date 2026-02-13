package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Subquest struct {
	ID              int32             `json:"id"`
	Name            string            `json:"name"`
	ParentSidequest NamedAPIResource  `json:"parent_sidequest"`
	Completions     []QuestCompletion `json:"completions"`
}

func (cfg *Config) getSubquest(r *http.Request, i handlerInput[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList], id int32) (Subquest, error) {
	subquest, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Subquest{}, err
	}

	response := Subquest{
		ID:              subquest.ID,
		Name:            subquest.Name,
		ParentSidequest: idToNamedAPIResource(cfg, cfg.e.sidequests, subquest.SidequestID),
		Completions:     convertObjSlice(cfg, subquest.Completions, convertQuestCompletion),
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
