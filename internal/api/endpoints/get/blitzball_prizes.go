package get

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func BlitzballPrize(r *http.Request, i handlerinput.Main[seedtypes.BlitzballPosition, apitypes.BlitzballPrize, apiresource.Named], id int32) (apitypes.BlitzballPrize, error) {
	bbPos, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.BlitzballPrize{}, err
	}

	response := apitypes.BlitzballPrize{
		ID:       bbPos.ID,
		Name:     fmt.Sprintf("%s - %s", bbPos.Category, bbPos.Slot),
		Category: bbPos.Category,
		Slot:     bbPos.Slot,
		Items:    ah.ConvertObjSlice(bbPos.Items, apiconvert.BlitzballItem),
	}

	return response, nil
}
