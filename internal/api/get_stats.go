package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getStat(r *http.Request, i handlerinput.Main[seedtypes.Stat, apitypes.Stat, apiresource.Named], id int32) (apitypes.Stat, error) {
	stat, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Stat{}, err
	}

	rel, err := assemble.GetStatRelationships(r, stat)
	if err != nil {
		return apitypes.Stat{}, err
	}

	response := apitypes.Stat{
		ID:                 stat.ID,
		Name:               stat.Name,
		Effect:             stat.Effect,
		MinVal:             stat.MinVal,
		MaxVal:             stat.MaxVal,
		MaxVal2:            stat.MaxVal2,
		ActivationSphere:   handlerinput.NameToNamedAPIResource(apilookups.E.Spheres, stat.ActivationSphere, nil),
		Spheres:            rel.Spheres,
		AutoAbilities:      rel.AutoAbilities,
		PlayerAbilities:    rel.PlayerAbilities,
		OverdriveAbilities: rel.OverdriveAbilities,
		ItemAbilities:      rel.ItemAbilities,
		TriggerCommands:    rel.TriggerCommands,
		StatusConditions:   rel.StatusConditions,
		Properties:         rel.Properties,
	}

	return response, nil
}

func retrieveStats(r *http.Request, i handlerinput.Main[seedtypes.Stat, apitypes.Stat, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
