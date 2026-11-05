package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getBlitzballPrize(r *http.Request, i handlerInput[seedtypes.BlitzballPosition, BlitzballPrize, NamedAPIResource], id int32) (BlitzballPrize, error) {
	bbPos, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return BlitzballPrize{}, err
	}

	response := BlitzballPrize{
		ID:       bbPos.ID,
		Name:     fmt.Sprintf("%s - %s", bbPos.Category, bbPos.Slot),
		Category: bbPos.Category,
		Slot:     bbPos.Slot,
		Items:    convertObjSlice(bbPos.Items, convertBlitzballItem),
	}

	return response, nil
}

func retrieveBlitzballPrizes(r *http.Request, i handlerInput[seedtypes.BlitzballPosition, BlitzballPrize, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnum(r, enums.T.BlitzballTournamentCategory, ids, ts.QpnCategory, config.Cfg.DB.GetBlitzballPrizeIDsByCategory),
	})
}
