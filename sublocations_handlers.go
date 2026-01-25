package main

import (
	"fmt"
	"net/http"

	//"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type Sublocation struct {
	ID                	int32                	`json:"id"`
	Name              	string               	`json:"name"`
	ParentLocation   	NamedAPIResource     	`json:"parent_location"`
	Areas				[]LocationAPIResource	`json:"areas"`
	LocRel
}


func (cfg *Config) HandleSublocations(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.sublocations
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
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: '/api/%s/{id}', or '/api/%s/{id}/{subsection}'. supported subsections: %s.", i.endpoint, i.endpoint, h.GetMapKeyStr(i.subsections)), nil)
		return
	}
}


func (cfg *Config) getSublocation(r *http.Request, i handlerInput[seeding.SubLocation, Sublocation, NamedAPIResource, NamedApiResourceList], id int32) (Sublocation, error) {
	sublocation, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Sublocation{}, err
	}

	areas, err := getResourcesDB(cfg, r, cfg.e.areas, sublocation, cfg.db.GetSublocationAreaIDs)
	if err != nil {
		return Sublocation{}, err
	}

	rel, err := getSublocationRelationships(cfg, r, sublocation)
	if err != nil {
		return Sublocation{}, err
	}

	response := Sublocation{
		ID:                	sublocation.ID,
		Name:              	sublocation.Name,
		ParentLocation:    	nameToNamedAPIResource(cfg, cfg.e.locations, sublocation.Location.Name, nil),
		Areas: 				areas,
		LocRel: 			rel,
	}

	return response, nil
}


func (cfg *Config) retrieveSublocations(r *http.Request, i handlerInput[seeding.SubLocation, Sublocation, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	filteredLists := []filteredResList[NamedAPIResource]{
		frl(idOnlyQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationSublocationIDs)),
		frl(idOnlyQueryWrapper(cfg, r, i, resources, "item", len(cfg.l.Items), getSublocationsByItem)),
		frl(idOnlyQueryWrapper(cfg, r, i, resources, "key-item", len(cfg.l.KeyItems), getSublocationsByKeyItem)),
		frl(boolQuery2(cfg, r, i, resources, "characters", cfg.db.GetSublocationIDsWithCharacters)),
		frl(boolQuery2(cfg, r, i, resources, "aeons", cfg.db.GetSublocationIDsWithAeons)),
		frl(boolQuery2(cfg, r, i, resources, "monsters", cfg.db.GetSublocationIDsWithMonsters)),
		frl(boolQuery2(cfg, r, i, resources, "boss_fights", cfg.db.GetSublocationIDsWithBosses)),
		frl(boolQuery2(cfg, r, i, resources, "shops", cfg.db.GetSublocationIDsWithShops)),
		frl(boolQuery2(cfg, r, i, resources, "treasures", cfg.db.GetSublocationIDsWithTreasures)),
		frl(boolQuery2(cfg, r, i, resources, "sidequests", cfg.db.GetSublocationIDsWithSidequests)),
		frl(boolQuery2(cfg, r, i, resources, "fmvs", cfg.db.GetSublocationIDsWithFMVs)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}