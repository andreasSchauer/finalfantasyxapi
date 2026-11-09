package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/filters"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetPlayerAbility(r *http.Request, i handlerinput.Main[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.Named], id int32) (apitypes.PlayerAbility, error) {
	ability, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}

	rel, err := assemble.GetPlayerAbilityRelationships(r, ability)
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

	battleInteractions, err := assemble.ApplyUser(r, i, response, ts.QpnAbilityUser)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func RetrievePlayerAbilities(r *http.Request, i handlerinput.Main[seedtypes.PlayerAbility, apitypes.PlayerAbility, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypePlayerAbility

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.PlayerAbilityCategory, ids, ts.QpnCategory, config.Cfg.DB.GetPlayerAbilityIDsByCategory),
		i.QueryEnumList(r, enums.T.DamageType, ids, ts.QpnDamageType, filters.GetTypedAbilityIDsByDamageType(abilityType)),
		i.QueryEnumList(r, enums.T.AttackType, ids, ts.QpnAttackType, filters.GetTypedAbilityIDsByAttackType(abilityType)),
		i.QueryEnumList(r, enums.T.TargetType, ids, ts.QpnTargetType, filters.GetTypedAbilityIDsByTargetType(abilityType)),
		i.QueryEnum(r, enums.T.DamageFormula, ids, ts.QpnDamageFormula, filters.GetTypedAbilityIDsByDamageFormula(abilityType)),
		i.QueryIntList(r, ids, ts.QpnMp, config.Cfg.DB.GetPlayerAbilityIDsByMpCost),
		i.QueryInt(r, ids, ts.QpnMpMin, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMin),
		i.QueryInt(r, ids, ts.QpnMpMax, config.Cfg.DB.GetPlayerAbilityIDsByMpCostMax),
		i.QueryIntList(r, ids, ts.QpnRank, filters.GetTypedAbilityIDsByRank(abilityType)),
		i.QueryNameIDListNul(r, ids, ts.QpnElement, apilookups.E.Elements.ResTypeSingle, lookup.L.Elements, filters.GetTypedAbilityIDsByElement(abilityType)),
		i.QueryNameID(r, ids, ts.QpnUser, apilookups.E.CharacterClasses.ResTypeSingle, lookup.L.CharClasses, config.Cfg.DB.GetPlayerAbilityIDsByCharClass),
		i.QueryID(r, ids, ts.QpnLearnItem, lookup.L.Items, config.Cfg.DB.GetPlayerAbilityIDsByLearnItem),
		i.QueryNameID(r, ids, ts.QpnRelatedStat, apilookups.E.Stats.ResTypeSingle, lookup.L.Stats, config.Cfg.DB.GetPlayerAbilityIDsByRelatedStat),
		i.QueryIDNul(r, ids, ts.QpnStatusInflict, lookup.L.StatusConditions, filters.GetTypedAbilityIDsByInflictedStatus(abilityType)),
		i.QueryIDNul(r, ids, ts.QpnStatusRemove, lookup.L.StatusConditions, filters.GetTypedAbilityIDsByRemovedStatus(abilityType)),
		i.QueryNameID(r, ids, ts.QpnStdSg, apilookups.E.Characters.ResTypeSingle, lookup.L.Characters, handlerinput.ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsStdSgChar)),
		i.QueryNameID(r, ids, ts.QpnExpSg, apilookups.E.Characters.ResTypeSingle, lookup.L.Characters, handlerinput.ToIntManyNull(config.Cfg.DB.GetPlayerAbilityIDsExpSgChar)),
		i.QueryBool(r, ids, ts.QpnOutsideBattle, config.Cfg.DB.GetPlayerAbilityIDsCanUseOutsideBattle),
		i.QueryBool(r, ids, ts.QpnCopycat, filters.GetTypedAbilityIDsByCanCopycat(abilityType)),
		i.QueryBool(r, ids, ts.QpnHelpBar, filters.GetTypedAbilityIDsByAppearsInHelpBar(abilityType)),
		i.QueryBool2(r, ids, ts.QpnUserAtk, filters.GetTypedAbilityIDsBasedOnUserAttack(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDarkable, filters.GetTypedAbilityIDsDarkable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnSilenceable, filters.GetTypedAbilityIDsSilenceable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnReflectable, filters.GetTypedAbilityIDsReflectable(abilityType)),
		i.QueryBool2(r, ids, ts.QpnDelay, filters.GetTypedAbilityIDsDealsDelay(abilityType)),
		i.QueryBool2(r, ids, ts.QpnStatChanges, filters.GetTypedAbilityIDsWithStatChanges(abilityType)),
		i.QueryBool2(r, ids, ts.QpnModChanges, filters.GetTypedAbilityIDsWithModifierChanges(abilityType)),
	})
}
