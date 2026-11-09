package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getPrimer(r *http.Request, i handlerinput.Main[seedtypes.Primer, apitypes.Primer, apiresource.Named], id int32) (apitypes.Primer, error) {
	primer, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Primer{}, err
	}

	rel, err := GetPrimerRelationships(r, primer)
	if err != nil {
		return apitypes.Primer{}, err
	}

	response := apitypes.Primer{
		ID:            primer.ID,
		Name:          primer.Name,
		KeyItem:       rel.KeyItem,
		Description:   rel.Description,
		AlBhedLetter:  primer.AlBhedLetter,
		EnglishLetter: primer.EnglishLetter,
		Treasures:     rel.Treasures,
		Areas:         rel.Areas,
	}

	return response, nil
}

func retrievePrimers(r *http.Request, i handlerinput.Main[seedtypes.Primer, apitypes.Primer, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{})
}
