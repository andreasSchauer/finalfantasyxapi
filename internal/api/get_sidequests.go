package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSidequest(r *http.Request, i handlerInput[seeding.Sidequest, Sidequest, NamedAPIResource, NamedApiResourceList], id int32) (Sidequest, error) {
	sidequest, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Sidequest{}, err
	}

	subquests, err := getResourcesDbItem(cfg, r, cfg.e.subquests, sidequest, cfg.db.GetSidequestSubquestIDs)
	if err != nil {
		return Sidequest{}, err
	}

	response := Sidequest{
		ID:         	sidequest.ID,
		Name:       	sidequest.Name,
		IsPostAirship: 	sidequest.IsPostAirship,
		Completion: 	convertObjPtr(cfg, sidequest.Completion, convertQuestCompletion),
		Subquests:  	subquests,
	}

	return response, nil
}

func (cfg *Config) retrieveSidequests(r *http.Request, i handlerInput[seeding.Sidequest, Sidequest, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(boolQuery(cfg, r, i, resources, "post_airship", cfg.db.GetSidequestIDsByPostAirship)),
	})
}
