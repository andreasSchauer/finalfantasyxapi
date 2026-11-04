package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (s *Server) getPrimer(r *http.Request, i handlerInput[seedtypes.Primer, Primer, NamedAPIResource], id int32) (Primer, error) {
	primer, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Primer{}, err
	}

	rel, err := getPrimerRelationships(s, r, primer)
	if err != nil {
		return Primer{}, err
	}

	keyItem, _ := lookup.GetResource(primer, lookup.L.KeyItems)

	response := Primer{
		ID:            primer.ID,
		Name:          primer.Name,
		KeyItem:       nameToNamedAPIResource(s, E.KeyItems, keyItem.Name, nil),
		Description:   keyItem.Description,
		AlBhedLetter:  primer.AlBhedLetter,
		EnglishLetter: primer.EnglishLetter,
		Treasures:     rel.Treasures,
		Areas:         rel.Areas,
	}

	return response, nil
}

func (s *Server) retrievePrimers(r *http.Request, i handlerInput[seedtypes.Primer, Primer, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{})
}
