package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)



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
