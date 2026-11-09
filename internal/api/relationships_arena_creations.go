package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetArenaCreationRelationships(r *http.Request, creation seedtypes.ArenaCreation) (apitypes.ArenaCreation, error) {
	var rel apitypes.ArenaCreation
	var err error
	
	subquest, _ := lookup.GetResourceByID(creation.SubquestID, lookup.L.SubquestsID)
	ia := subquest.Completion.Reward
	rel.Reward = apilookups.E.AllItems.ToResAmtFromName(ia)

	rel.RequiredMonsters, err = getArenaCreationMonsters(r, creation)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	return rel, nil
}