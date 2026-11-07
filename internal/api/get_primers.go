package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getPrimer(r *http.Request, i handlerinput.HandlerInput[seedtypes.Primer, apitypes.Primer, apiresource.NamedAPIResource], id int32) (apitypes.Primer, error) {
	primer, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Primer{}, err
	}

	rel, err := getPrimerRelationships(r, primer)
	if err != nil {
		return apitypes.Primer{}, err
	}

	keyItem, _ := lookup.GetResource(primer, lookup.L.KeyItems)

	response := apitypes.Primer{
		ID:            primer.ID,
		Name:          primer.Name,
		KeyItem:       handlerinput.NameToNamedAPIResource(E.KeyItems, keyItem.Name, nil),
		Description:   keyItem.Description,
		AlBhedLetter:  primer.AlBhedLetter,
		EnglishLetter: primer.EnglishLetter,
		Treasures:     rel.Treasures,
		Areas:         rel.Areas,
	}

	return response, nil
}

func retrievePrimers(r *http.Request, i handlerinput.HandlerInput[seedtypes.Primer, apitypes.Primer, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{})
}
