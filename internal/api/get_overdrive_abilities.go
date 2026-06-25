package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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

func (cfg *Config) retrieveOverdriveAbilities(r *http.Request, i handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeOverdriveAbility

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.AttackType, ids, qpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.TargetType, ids, qpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType))),
		fidl(enumQuery(r, i, cfg.t.DamageFormula, ids, qpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		fidl(intListQuery(cfg, r, i, ids, qpnRank, getTypedAbilityIDsByRank(cfg, abilityType))),
		fidl(nameIdListQueryNul(cfg, r, i, ids, qpnElement, cfg.e.elements.resTypeSing, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		fidl(nameIdQuery(r, i, ids, qpnUser, cfg.e.characterClasses.resTypeSing, cfg.l.CharClasses, cfg.db.GetOverdriveAbilityIDsByCharClass)),
		fidl(nameIdQuery(r, i, ids, qpnRelatedStat, cfg.e.stats.resTypeSing, cfg.l.Stats, cfg.db.GetOverdriveAbilityIDsByRelatedStat)),
		fidl(idQueryNul(r, i, ids, qpnStatusInflict, cfg.l.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, qpnStatusRemove, cfg.l.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnCanCrit, getTypedAbilityIDsCanCrit(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType))),
	})
}
