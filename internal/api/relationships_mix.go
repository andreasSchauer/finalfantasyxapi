package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetMixRelationships(r *http.Request, mix seedtypes.Mix) (apitypes.Mix, error) {
	overdrive, _ := lookup.GetResourceByID(mix.OverdriveID, lookup.L.OverdrivesID)
	var err error
	
	rel := apitypes.Mix{
		Overdrive:    handlerinput.NameToNamedAPIResource(apilookups.E.Overdrives, overdrive.Name, nil),
		Description:  overdrive.Description,
		Effect:       overdrive.Effect,
	}

	rel.Combinations, err = getMixCombinations(r, mix)
	if err != nil {
		return apitypes.Mix{}, err
	}

	return rel, nil
}