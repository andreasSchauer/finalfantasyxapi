package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getBlitzballPrize(r *http.Request, i handlerInput[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedApiResourceList], id int32) (BlitzballPrize, error) {
	bbPos, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return BlitzballPrize{}, err
	}

	response := BlitzballPrize{
		ID:       bbPos.ID,
		Name:     fmt.Sprintf("%s - %s", bbPos.Category, bbPos.Slot),
		Category: bbPos.Category,
		Slot:     bbPos.Slot,
		Items:    convertObjSlice(cfg, bbPos.Items, convertBlitzballItem),
	}

	return response, nil
}

func (cfg *Config) retrieveBlitzballPrizes(r *http.Request, i handlerInput[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumQuery(cfg, r, i, cfg.t.BlitzballTournamentCategory, resources, "category", cfg.db.GetBlitzballPrizeIDsByCategory)),
	})
}
