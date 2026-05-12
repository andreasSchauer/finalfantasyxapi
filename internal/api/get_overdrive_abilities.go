package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getOverdriveAbility(r *http.Request, i handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList], id int32) (OverdriveAbility, error) {
	ability, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return OverdriveAbility{}, err
	}

	overdriveIDs, err := cfg.db.GetOverdriveAbilityOverdriveIDs(r.Context(), id)
	if err != nil {
		return OverdriveAbility{}, newHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive.", err)
	}
	overdrive, _ := seeding.GetResourceByID(overdriveIDs[0], cfg.l.OverdrivesID)

	response := OverdriveAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.e.abilities, ability.Ability.ID),
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Overdrives:         idsToAPIResources(cfg, cfg.e.overdrives, overdriveIDs),
		OverdriveCommand:   namePtrToNamedAPIResPtr(cfg, cfg.e.overdriveCommands, overdrive.OverdriveCommand, nil),
		User:               nameToNamedAPIResource(cfg, cfg.e.characterClasses, overdrive.User, nil),
		RelatedStats:       namesToNamedAPIResources(cfg, cfg.e.stats, ability.RelatedStats),
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (cfg *Config) retrieveOverdriveAbilities(r *http.Request, i handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}
	abilityType := database.AbilityTypeOverdriveAbility

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", getTypedAbilityIDsByAttackType(cfg, abilityType))),
		frl(enumListQuery(cfg, r, i, cfg.t.TargetType, resources, "target_type", getTypedAbilityIDsByTargetType(cfg, abilityType))),
		frl(enumQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		frl(intListQuery(cfg, r, i, resources, "rank", getTypedAbilityIDsByRank(cfg, abilityType))),
		frl(nameIdListQueryNul(cfg, r, i, resources, "element", cfg.e.elements.resourceType, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		frl(nameIdQuery(cfg, r, i, resources, "user", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetOverdriveAbilityIDsByCharClass)),
		frl(nameIdQuery(cfg, r, i, resources, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetOverdriveAbilityIDsByRelatedStat)),
		frl(idQueryNul(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		frl(idQueryNul(cfg, r, i, resources, "status_remove", len(cfg.l.StatusConditions), getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "can_crit", getTypedAbilityIDsCanCrit(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "delay", getTypedAbilityIDsDealsDelay(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "stat_changes", getTypedAbilityIDsWithStatChanges(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "mod_changes", getTypedAbilityIDsWithModifierChanges(cfg, abilityType))),
	})
}
