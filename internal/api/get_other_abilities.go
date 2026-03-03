package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getUnspecifiedAbility(r *http.Request, i handlerInput[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList], id int32) (UnspecifiedAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return UnspecifiedAbility{}, err
	}

	response := UnspecifiedAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
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

	response, err = applyUnspecifiedAbilityUser(cfg, r, response, "user")
	if err != nil {
		return UnspecifiedAbility{}, err
	}

	return response, nil
}

func (cfg *Config) retrieveUnspecifiedAbilities(r *http.Request, i handlerInput[seeding.UnspecifiedAbility, UnspecifiedAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.DamageType, resources, "damage_type", cfg.db.GetUnspecifiedAbilityIDsByDamageType)),
		frl(typeQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", cfg.db.GetUnspecifiedAbilityIDsByAttackType)),
		frl(typeQuery(cfg, r, i, cfg.t.TargetType, resources, "target_type", cfg.db.GetUnspecifiedAbilityIDsByTargetType)),
		frl(typeQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", cfg.db.GetUnspecifiedAbilityIDsByDamageFormula)),
		frl(intQueryNullable(cfg, r, i, resources, "rank", cfg.db.GetUnspecifiedAbilityIDsByRank)),
		frl(nameOrIdQuery(cfg, r, i, resources, "char_class", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetUnspecifiedAbilityIDsByCharClass)),
		frl(idQueryWrapper(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), getUnspecifiedAbilitiesInflictedStatus)),
		frl(boolQuery(cfg, r, i, resources, "copycat", cfg.db.GetUnspecifiedAbilityIDsByCanCopycat)),
		frl(boolQuery(cfg, r, i, resources, "help_bar", cfg.db.GetUnspecifiedAbilityIDsByAppearsInHelpBar)),
		frl(boolQuery2(cfg, r, i, resources, "phys_atk", cfg.db.GetUnspecifiedAbilityIDsBasedOnPhysAttack)),
		frl(boolQuery2(cfg, r, i, resources, "darkable", cfg.db.GetUnspecifiedAbilityIDsDarkable)),
	})
}
