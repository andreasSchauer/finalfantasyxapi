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
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.e.abilities, ability.Ability.ID),
		Description:        ability.Description,
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		LearnedBy:          namesToNamedAPIResources(cfg, cfg.e.characterClasses, ability.LearnedBy),
		Topmenu:            namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, ability.Topmenu, nil),
		Submenu:            namePtrToNamedAPIResPtr(cfg, cfg.e.submenus, ability.Submenu, nil),
		OpenSubmenu:        namePtrToNamedAPIResPtr(cfg, cfg.e.submenus, ability.OpenSubmenu, nil),
		Cursor:             ability.Cursor,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	battleInteractions, err := applyUser(cfg, r, i, response, qpnAbilityUser)
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

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(intListQuery(cfg, r, i, ids, qpnRank, getTypedAbilityIDsByRank(cfg, abilityType))),
		fidl(nameIdQuery(r, i, ids, qpnUser, cfg.e.characterClasses.resTypeSing, cfg.l.CharClasses, cfg.db.GetMiscAbilityIDsByCharClass)),
		fidl(boolQuery(r, i, ids, qpnCopycat, getTypedAbilityIDsByCanCopycat(cfg, abilityType))),
		fidl(boolQuery(r, i, ids, qpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnUserAtk, getTypedAbilityIDsBasedOnUserAttack(cfg, abilityType))),
	})
}
