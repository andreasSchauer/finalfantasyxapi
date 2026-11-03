package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (cfg *Config) getMiscAbility(r *http.Request, i handlerInput[seedtypes.MiscAbility, MiscAbility, NamedAPIResource], id int32) (MiscAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return MiscAbility{}, err
	}

	response := MiscAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.E.Abilities, ability.Ability.ID),
		Description:        ability.Description,
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		LearnedBy:          namesToNamedAPIResources(cfg, cfg.E.CharacterClasses, ability.LearnedBy),
		Topmenu:            namePtrToNamedAPIResPtr(cfg, cfg.E.Topmenus, ability.Topmenu, nil),
		Submenu:            namePtrToNamedAPIResPtr(cfg, cfg.E.Submenus, ability.Submenu, nil),
		OpenSubmenu:        namePtrToNamedAPIResPtr(cfg, cfg.E.Submenus, ability.OpenSubmenu, nil),
		Cursor:             ability.Cursor,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(cfg, r, i, response, QpnAbilityUser)
	if err != nil {
		return MiscAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrieveMiscAbilities(r *http.Request, i handlerInput[seedtypes.MiscAbility, MiscAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeMiscAbility

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryIntList(cfg, r, ids, QpnRank, getTypedAbilityIDsByRank(cfg, abilityType)),
		i.queryNameID(r, ids, QpnUser, cfg.E.CharacterClasses.resTypeSingle, cfg.L.CharClasses, cfg.db.GetMiscAbilityIDsByCharClass),
		i.queryBool(r, ids, QpnCopycat, getTypedAbilityIDsByCanCopycat(cfg, abilityType)),
		i.queryBool(r, ids, QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType)),
		i.queryBool2(r, ids, QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(cfg, abilityType)),
	})
}
