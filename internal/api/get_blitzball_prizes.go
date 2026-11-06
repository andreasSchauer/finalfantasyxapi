package api

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getBlitzballPrize(r *http.Request, i handlerInput[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.NamedAPIResource], id int32) (apitypes.BlitzballPrize, error) {
	bbPos, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.BlitzballPrize{}, err
	}

	response := apitypes.BlitzballPrize{
		ID:       bbPos.ID,
		Name:     fmt.Sprintf("%s - %s", bbPos.Category, bbPos.Slot),
		Category: bbPos.Category,
		Slot:     bbPos.Slot,
		Items:    ah.ConvertObjSlice(bbPos.Items, convertBlitzballItem),
	}

	return response, nil
}

func retrieveBlitzballPrizes(r *http.Request, i handlerInput[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnum(r, enums.T.BlitzballTournamentCategory, ids, ts.QpnCategory, config.Cfg.DB.GetBlitzballPrizeIDsByCategory),
	})
}
