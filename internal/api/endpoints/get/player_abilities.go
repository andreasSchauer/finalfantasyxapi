package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"

	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func PlayerAbility(r *http.Request, i handlerinput.Main[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.Named], id int32) (apitypes.PlayerAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}
	
	response := apitypes.PlayerAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification:         ability.Specification,
		UntypedAbility:        handlerinput.IDToTypedAPIResource(apilookups.E.Abilities, ability.Ability.ID),
		Description:           ability.Description,
		Effect:                ability.Effect,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		CanUseOutsideBattle:   ability.CanUseOutsideBattle,
		MpCost:                ability.MPCost,
		Category:              ability.Category,
		AeonLearnItem:         apilookups.E.AllItems.ToResAmtFromNamePtr(ability.AeonLearnItem),
		LearnedBy:             handlerinput.NamesToNamedAPIResources(apilookups.E.CharacterClasses, ability.LearnedBy),
		RelatedStats:          handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, ability.RelatedStats),
		StandardGridCharacter: handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Characters, ability.StandardGridPos, nil),
		ExpertGridCharacter:   handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Characters, ability.ExpertGridPos, nil),
		Topmenu:               handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, ability.Topmenu, nil),
		Submenu:               handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Submenus, ability.Submenu, nil),
		OpenSubmenu:           handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Submenus, ability.OpenSubmenu, nil),
		Cursor:                ability.Cursor,
		BattleInteractions:    ah.ConvertObjSlice(ability.BattleInteractions, apiconvert.BattleInteraction),
	}

	err = getPlayerAbilityRelationships(r, ability, &response)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}
	
	battleInteractions, err := applyUser(r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func getPlayerAbilityRelationships(r *http.Request, ability seedtypes.PlayerAbility, response *apitypes.PlayerAbility) error {
	var err error

	response.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return err
	}

	return nil
}
