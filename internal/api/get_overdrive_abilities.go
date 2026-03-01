package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getOverdriveAbility(r *http.Request, i handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList], id int32) (OverdriveAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return OverdriveAbility{}, err
	}

	overdriveIDs, err := cfg.db.GetOverdriveAbilityOverdriveIDs(r.Context(), id)
	if err != nil {
		return OverdriveAbility{}, newHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive.", err)
	}
	overdrive, _ := seeding.GetResourceByID(overdriveIDs[0], cfg.l.OverdrivesID)

	response := OverdriveAbility{
		ID:                    ability.ID,
		Name:                  ability.Name,
		Version:               ability.Version,
		Specification: 		   ability.Specification,
		Rank:                  overdrive.Rank,
		AppearsInHelpBar:      overdrive.AppearsInHelpBar,
		CanCopycat:            overdrive.CanCopycat,
		Overdrives: 		   idsToAPIResources(cfg, cfg.e.overdrives, overdriveIDs),
		OverdriveCommand: 	   namePtrToNamedAPIResPtr(cfg, cfg.e.overdriveCommands, overdrive.OverdriveCommand, nil),
		User:             	   nameToNamedAPIResource(cfg, cfg.e.characterClasses, overdrive.User, nil),
		RelatedStats:          namesToNamedAPIResources(cfg, cfg.e.stats, ability.RelatedStats),
		Topmenu:               namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, overdrive.Topmenu, nil),
		Cursor:                overdrive.Cursor,
		BattleInteractions:    convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}



func (cfg *Config) retrieveOverdriveAbilities(r *http.Request, i handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", cfg.db.GetOverdriveAbilityIDsByAttackType)),
		frl(typeQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", cfg.db.GetOverdriveAbilityIDsByDamageFormula)),
		frl(intQueryNullable(cfg, r, i, resources, "rank", cfg.db.GetOverdriveAbilityIDsByRank)),
		frl(nameOrIdQuery(cfg, r, i, resources, "element", cfg.e.elements.resourceType, cfg.l.Elements, cfg.db.GetOverdriveAbilityIDsByElement)),
		frl(nameOrIdQuery(cfg, r, i, resources, "char_class", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetOverdriveAbilityIDsByCharClass)),
		frl(nameOrIdQuery(cfg, r, i, resources, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetOverdriveAbilityIDsByRelatedStat)),
		frl(idQueryWrapper(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), getOverdriveAbilitiesInflictedStatus)),
		frl(idQuery(cfg, r, i, resources, "status_remove", len(cfg.l.StatusConditions), cfg.db.GetOverdriveAbilityIDsByRemovedStatus)),
		frl(boolQuery2(cfg, r, i, resources, "delay", cfg.db.GetOverdriveAbilityIDsDealsDelay)),
		frl(boolQuery2(cfg, r, i, resources, "stat_changes", cfg.db.GetOverdriveAbilityIDsWithStatChanges)),
		frl(boolQuery2(cfg, r, i, resources, "mod_changes", cfg.db.GetOverdriveAbilityIDsWithModifierChanges)),
	})
}
