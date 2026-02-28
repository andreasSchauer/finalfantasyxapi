package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getOtherAbility(r *http.Request, i handlerInput[seeding.OtherAbility, OtherAbility, NamedAPIResource, NamedApiResourceList], id int32) (OtherAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return OtherAbility{}, err
	}

	response := OtherAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification: 		   ability.Specification,
		Description:           ability.Description,
		Effect:                ability.Effect,
		Rank:                  ability.Rank,
		AppearsInHelpBar:      ability.AppearsInHelpBar,
		CanCopycat:            ability.CanCopycat,
		LearnedBy:             namesToNamedAPIResources(cfg, cfg.e.characterClasses, ability.LearnedBy),
		Topmenu:               namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, ability.Topmenu, nil),
		Submenu:               namePtrToNamedAPIResPtr(cfg, cfg.e.submenus, ability.Submenu, nil),
		OpenSubmenu:           namePtrToNamedAPIResPtr(cfg, cfg.e.submenus, ability.OpenSubmenu, nil),
		Cursor:                ability.Cursor,
		BattleInteractions:    convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	response, err = applyOtherAbilityUser(cfg, r, response, "user")
	if err != nil {
		return OtherAbility{}, err
	}

	return response, nil
}



func (cfg *Config) retrieveOtherAbilities(r *http.Request, i handlerInput[seeding.OtherAbility, OtherAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.DamageType, resources, "damage_type", cfg.db.GetOtherAbilityIDsByDamageType)),
		frl(typeQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", cfg.db.GetOtherAbilityIDsByAttackType)),
		frl(typeQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", cfg.db.GetOtherAbilityIDsByDamageFormula)),
		frl(intQueryNullable(cfg, r, i, resources, "rank", cfg.db.GetOtherAbilityIDsByRank)),
		frl(nameOrIdQuery(cfg, r, i, resources, "char_class", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetOtherAbilityIDsByCharClass)),
		frl(idQueryWrapper(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), getOtherAbilitiesInflictedStatus)),
		frl(boolQuery(cfg, r, i, resources, "copycat", cfg.db.GetOtherAbilityIDsByCanCopycat)),
		frl(boolQuery(cfg, r, i, resources, "help_bar", cfg.db.GetOtherAbilityIDsByAppearsInHelpBar)),
		frl(boolQuery2(cfg, r, i, resources, "phys_atk", cfg.db.GetOtherAbilityIDsBasedOnPhysAttack)),
		frl(boolQuery2(cfg, r, i, resources, "darkable", cfg.db.GetOtherAbilityIDsDarkable)),
	})
}
