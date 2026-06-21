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
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}
	abilityType := database.AbilityTypeOverdriveAbility

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.AttackType, ids, "attack_type", getTypedAbilityIDsByAttackType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.TargetType, ids, "target_type", getTypedAbilityIDsByTargetType(cfg, abilityType))),
		fidl(enumQuery(r, i, cfg.t.DamageFormula, ids, "damage_formula", getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		fidl(intListQuery(cfg, r, i, ids, "rank", getTypedAbilityIDsByRank(cfg, abilityType))),
		fidl(nameIdListQueryNul(cfg, r, i, ids, "element", cfg.e.elements.resourceType, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		fidl(nameIdQuery(r, i, ids, "user", cfg.e.characterClasses.resourceType, cfg.l.CharClasses, cfg.db.GetOverdriveAbilityIDsByCharClass)),
		fidl(nameIdQuery(r, i, ids, "related_stat", cfg.e.stats.resourceType, cfg.l.Stats, cfg.db.GetOverdriveAbilityIDsByRelatedStat)),
		fidl(idQueryNul(r, i, ids, "status_inflict", cfg.l.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, "status_remove", cfg.l.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "can_crit", getTypedAbilityIDsCanCrit(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "delay", getTypedAbilityIDsDealsDelay(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "stat_changes", getTypedAbilityIDsWithStatChanges(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "mod_changes", getTypedAbilityIDsWithModifierChanges(cfg, abilityType))),
	})
}
