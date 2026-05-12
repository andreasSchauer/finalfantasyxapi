package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMiscAbility(r *http.Request, i handlerInput[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedApiResourceList], id int32) (MiscAbility, error) {
	ability, err := verifyParamsAndGet(cfg, r, i, id)
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

	battleInteractions, err := applyUser(cfg, r, i, response, "ability_user")
	if err != nil {
		return MiscAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrieveMiscAbilities(r *http.Request, i handlerInput[seeding.MiscAbility, MiscAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}
	abilityType := database.AbilityTypeMiscAbility

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(intListQuery(cfg, r, i, resources, "rank", getTypedAbilityIDsByRank(cfg, abilityType))),
		frl(nameIdQuery(cfg, r, i, resources, "user", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetMiscAbilityIDsByCharClass)),
		frl(boolQuery(cfg, r, i, resources, "copycat", getTypedAbilityIDsByCanCopycat(cfg, abilityType))),
		frl(boolQuery(cfg, r, i, resources, "help_bar", getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "user_atk", getTypedAbilityIDsBasedOnUserAttack(cfg, abilityType))),
	})
}
