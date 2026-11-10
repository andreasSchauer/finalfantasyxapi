package get

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetBlitzballPrize(r *http.Request, i handlerinput.Main[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.Named], id int32) (apitypes.BlitzballPrize, error) {
	bbPos, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.BlitzballPrize{}, err
	}

	response := apitypes.BlitzballPrize{
		ID:       bbPos.ID,
		Name:     fmt.Sprintf("%s - %s", bbPos.Category, bbPos.Slot),
		Category: bbPos.Category,
		Slot:     bbPos.Slot,
		Items:    ah.ConvertObjSlice(bbPos.Items, apiconv.ConvertBlitzballItem),
	}

	return response, nil
}
