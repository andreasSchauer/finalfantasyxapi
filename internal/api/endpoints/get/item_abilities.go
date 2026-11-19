package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func ItemAbility(r *http.Request, i handlerinput.Main[seedtypes.ItemAbility, apitypes.ItemAbility, apiresource.Named], id int32) (apitypes.ItemAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.ItemAbility{}, err
	}

	response := apitypes.ItemAbility{
		ID:                  ability.ID,
		Name:                ability.Name,
		UntypedAbility:      handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, ability.Ability.ID),
		Rank:                ability.Rank,
		AppearsInHelpBar:    ability.AppearsInHelpBar,
		CanCopycat:          ability.CanCopycat,
		BattleInteractions:  ah.ConvertObjSlice(ability.BattleInteractions, apiconvert.BattleInteraction),
	}

	getItemAbilityRelationships(ability, &response)

	return response, nil
}

func getItemAbilityRelationships(ability seedtypes.ItemAbility, response *apitypes.ItemAbility) {
	item, _ := lookup.GetResourceByID(ability.ID, lookup.L.ItemsID)

	response.Item = handlerinput.IDToNamedAPIResource(apilookups.E.Items, item.ID)
	response.Description = item.Description
	response.Effect = item.Effect
	response.Category = item.Category
	response.RelatedStats = handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, item.RelatedStats)
	response.Cursor = item.Cursor

	if item.Usability == string(database.ItemUsabilityOutsideBattle) || item.Usability == string(database.ItemUsabilityAlways) {
		response.CanUseOutsideBattle = true
	}
}
