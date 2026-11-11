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

	rel := getItemAbilityRelationships(ability)

	response := apitypes.ItemAbility{
		ID:                  ability.ID,
		Name:                ability.Name,
		Item:                rel.Item,
		UntypedAbility:      handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, ability.Ability.ID),
		Description:         rel.Description,
		Effect:              rel.Effect,
		Category:            rel.Category,
		CanUseOutsideBattle: rel.CanUseOutsideBattle,
		Rank:                ability.Rank,
		AppearsInHelpBar:    ability.AppearsInHelpBar,
		CanCopycat:          ability.CanCopycat,
		RelatedStats:        rel.RelatedStats,
		Cursor:              rel.Cursor,
		BattleInteractions:  ah.ConvertObjSlice(ability.BattleInteractions, apiconvert.BattleInteraction),
	}

	return response, nil
}

func getItemAbilityRelationships(ability seedtypes.ItemAbility) apitypes.ItemAbility {
	item, _ := lookup.GetResourceByID(ability.ID, lookup.L.ItemsID)

	rel := apitypes.ItemAbility{
		Item:         handlerinput.IDToNamedAPIResource(apilookups.E.Items, item.ID),
		Description:  item.Description,
		Effect:       item.Effect,
		Category:     item.Category,
		RelatedStats: handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, item.RelatedStats),
		Cursor:       item.Cursor,
	}

	if item.Usability == string(database.ItemUsabilityOutsideBattle) || item.Usability == string(database.ItemUsabilityAlways) {
		rel.CanUseOutsideBattle = true
	}

	return rel
}
