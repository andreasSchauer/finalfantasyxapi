package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSubquest(r *http.Request, i handlerInput[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList], id int32) (Subquest, error) {
	subquest, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Subquest{}, err
	}

	response := Subquest{
		ID:              subquest.ID,
		Name:            subquest.Name,
		ParentSidequest: idToNamedAPIResource(cfg, cfg.e.sidequests, subquest.SidequestID),
		Completions:     convertObjSlice(cfg, subquest.Completions, convertQuestCompletion),
	}

	if response.ParentSidequest.Name == "monster arena" {
		response.ArenaCreation = namePtrToNamedAPIResPtr(cfg, cfg.e.arenaCreations, &subquest.Name, nil)
	}

	return response, nil
}

func (cfg *Config) retrieveSubquests(r *http.Request, i handlerInput[seeding.Subquest, Subquest, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(boolQuery2(cfg, r, i, resources, "post_airship", cfg.db.GetSubquestIDsByPostAirship)),
	})
}
