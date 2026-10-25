package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMiscAbility(r *http.Request, i handlerInput[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedApiResourceList], id int32) (MiscAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveMiscAbilities(r *http.Request, i handlerInput[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeMiscAbility

	return filterIDs(cfg, r, i, ids, []IdFilter{
		intListQuery(cfg, r, i, ids, QpnRank, getTypedAbilityIDsByRank(cfg, abilityType)),
		nameIdQuery(r, i, ids, QpnUser, cfg.E.CharacterClasses.resTypeSingle, cfg.L.CharClasses, cfg.db.GetMiscAbilityIDsByCharClass),
		boolQuery(r, i, ids, QpnCopycat, getTypedAbilityIDsByCanCopycat(cfg, abilityType)),
		boolQuery(r, i, ids, QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(cfg, abilityType)),
	})
}
