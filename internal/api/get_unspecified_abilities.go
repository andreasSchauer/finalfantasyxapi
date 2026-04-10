package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getUnspecifiedAbility(r *http.Request, i handlerInput[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList], id int32) (UnspecifiedAbility, error) {
	ability, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return UnspecifiedAbility{}, err
	}

	response := UnspecifiedAbility{
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
		return UnspecifiedAbility{}, err
	}
	response.BattleInteractions = battleInteractions

	return response, nil
}

func (cfg *Config) retrieveUnspecifiedAbilities(r *http.Request, i handlerInput[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(intListQuery(cfg, r, i, resources, "rank", cfg.db.GetUnspecifiedAbilityIDsByRank)),
		frl(nameIdQuery(cfg, r, i, resources, "user", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetUnspecifiedAbilityIDsByCharClass)),
		frl(boolQuery(cfg, r, i, resources, "copycat", cfg.db.GetUnspecifiedAbilityIDsByCanCopycat)),
		frl(boolQuery(cfg, r, i, resources, "help_bar", cfg.db.GetUnspecifiedAbilityIDsByAppearsInHelpBar)),
		frl(boolQuery2(cfg, r, i, resources, "user_atk", cfg.db.GetUnspecifiedAbilityIDsBasedOnUserAttack)),
	})
}
