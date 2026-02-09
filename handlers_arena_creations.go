package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ArenaCreation struct {
	ID                        	int32              	`json:"id"`
	Name                      	string             	`json:"name"`
	Category                  	string             	`json:"category"`
	Monster                   	NamedAPIResource   	`json:"monster"`
	ParentSubquest            	NamedAPIResource  	`json:"parent_subquest"`
	Reward					  	ItemAmount		 	`json:"reward"`
	RequiredCatchAmount       	int32              	`json:"required_catch_amount"`
	UnlockedCreationsCategory	*string				`json:"unlocked_creations_category,omitempty"`
	RequiredMonsters          	[]NamedAPIResource 	`json:"required_monsters,omitempty"`
}

type MonsterFilter struct {
	RequiredArea              *string
	RequiredSpecies           *string
	CreationsUnlockedCategory *string
	UnderwaterOnly            bool
}

func (mf MonsterFilter) IsZero() bool {
	return 	mf.RequiredArea == nil &&
			mf.RequiredSpecies == nil &&
			mf.CreationsUnlockedCategory == nil &&
			mf.UnderwaterOnly == false
}

func createMonsterFilter(creation seeding.ArenaCreation) MonsterFilter {
	return MonsterFilter{
		RequiredArea: 				creation.RequiredArea,
		RequiredSpecies: 			creation.RequiredSpecies,
		CreationsUnlockedCategory: 	creation.CreationsUnlockedCategory,
		UnderwaterOnly: 			creation.UnderwaterOnly,
	}
}


func (cfg *Config) HandleArenaCreations(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.arenaCreations

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

func (cfg *Config) getArenaCreation(r *http.Request, i handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList], id int32) (ArenaCreation, error) {
	creation, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return ArenaCreation{}, err
	}

	subquest, _ := seeding.GetResourceByID(creation.SubquestID, cfg.l.SubquestsID)
	ia := subquest.Completions[0].Reward

	monsters, err := getArenaCreationMonsters(cfg, r, creation)
	if err != nil {
		return ArenaCreation{}, err
	}

	response := ArenaCreation{
		ID:                     	creation.ID,
		Name:                   	creation.Name,
		Category:               	creation.Category,
		Monster:                	idToNamedAPIResource(cfg, cfg.e.monsters, *creation.MonsterID),
		ParentSubquest:         	idToNamedAPIResource(cfg, cfg.e.subquests, creation.SubquestID),
		Reward: 					convertItemAmount(cfg, ia),
		RequiredCatchAmount:    	creation.Amount,
		UnlockedCreationsCategory: 	creation.CreationsUnlockedCategory,
		RequiredMonsters: 			monsters,
	}

	

	return response, nil
}


func (cfg *Config) retrieveArenaCreations(r *http.Request, i handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	filteredLists := []filteredResList[NamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.ArenaCreationCategory, resources, "category", cfg.db.GetArenaCreationIDsByCategory)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}


func getArenaCreationMonsters(cfg *Config, r *http.Request, creation seeding.ArenaCreation) ([]NamedAPIResource, error) {
	mf := createMonsterFilter(creation)

	if mf.IsZero() || mf.CreationsUnlockedCategory != nil {
		return nil, nil
	}
	
	monsterIdSlices := []filteredIdList{}

	if mf.RequiredArea != nil {
		area := h.NullMaCreationArea(mf.RequiredArea)
		idSlice := fidl(cfg.db.GetCaptureMonsterIDsByMaCreationArea(r.Context(), area))
		monsterIdSlices = append(monsterIdSlices, idSlice)
	}

	if mf.RequiredSpecies != nil {
		species := database.MonsterSpecies(*mf.RequiredSpecies)
		idSlice := fidl(cfg.db.GetCaptureMonsterIDsBySpecies(r.Context(), species))
		monsterIdSlices = append(monsterIdSlices, idSlice)
	}

	if mf.UnderwaterOnly {
		idSlice := fidl(cfg.db.GetCaptureMonsterIDsByIsUnderwater(r.Context()))
		monsterIdSlices = append(monsterIdSlices, idSlice)
	}

	monsterIDs, err := filterIdSlices(monsterIdSlices)
	if err != nil {
		return nil, err
	}

	monsters := idsToAPIResources(cfg, cfg.e.monsters, monsterIDs)

	return monsters, nil
}