package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type BlitzballPrize struct {
	ID       int32          `json:"id"`
	Category string         `json:"category"`
	Slot     string         `json:"slot"`
	Items    []PossibleItem `json:"items"`
}

func convertBlitzballItem(cfg *Config, bi seeding.BlitzballItem) PossibleItem {
	return convertPossibleItem(cfg, bi.PossibleItem)
}

func (cfg *Config) getBlitzballPrize(r *http.Request, i handlerInput[seeding.BlitzballPosition, BlitzballPrize, UnnamedAPIResource, UnnamedApiResourceList], id int32) (BlitzballPrize, error) {
	bbPos, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return BlitzballPrize{}, err
	}

	response := BlitzballPrize{
		ID:       bbPos.ID,
		Category: bbPos.Category,
		Slot:     bbPos.Slot,
		Items:    convertObjSlice(cfg, bbPos.Items, convertBlitzballItem),
	}

	return response, nil
}

func (cfg *Config) retrieveBlitzballPrizes(r *http.Request, i handlerInput[seeding.BlitzballPosition, BlitzballPrize, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[UnnamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.BlitzballTournamentCategory, resources, "category", cfg.db.GetBlitzballPrizeIDsByCategory)),
	})
}
