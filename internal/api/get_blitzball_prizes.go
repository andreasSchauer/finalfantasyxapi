package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getBlitzballPrize(r *http.Request, i handlerInput[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedAPIResourceList], id int32) (BlitzballPrize, error) {
	bbPos, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveBlitzballPrizes(r *http.Request, i handlerInput[seeding.BlitzballPosition, BlitzballPrize, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		enumQuery(r, i, cfg.T.BlitzballTournamentCategory, ids, QpnCategory, cfg.db.GetBlitzballPrizeIDsByCategory),
	})
}
