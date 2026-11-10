package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetPlayerAbility(r *http.Request, i handlerinput.Main[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.Named], id int32) (apitypes.PlayerAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}

	rel, err := getPlayerAbilityRelationships(r, ability)
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
		Monsters:              rel.Monsters,
		BattleInteractions:    ah.ConvertObjSlice(ability.BattleInteractions, apiconv.ConvertBattleInteraction),
	}

	battleInteractions, err := ApplyUser(r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func getPlayerAbilityRelationships(r *http.Request, ability seedtypes.PlayerAbility) (apitypes.PlayerAbility, error) {
	var rel apitypes.PlayerAbility
	var err error

	rel.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}

	return rel, nil
}
