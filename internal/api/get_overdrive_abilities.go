package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getOverdriveAbility(r *http.Request, i handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource], id int32) (OverdriveAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return OverdriveAbility{}, err
	}

	overdriveIDs, err := cfg.db.GetOverdriveAbilityOverdriveIDs(r.Context(), id)
	if err != nil {
		return OverdriveAbility{}, newHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive.", err)
	}
	overdrive, _ := seeding.GetResourceByID(overdriveIDs[0], cfg.L.OverdrivesID)

	response := OverdriveAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.E.Abilities, ability.Ability.ID),
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Overdrives:         cfg.E.Overdrives.toResources(cfg, overdriveIDs),
		OverdriveCommand:   namePtrToNamedAPIResPtr(cfg, cfg.E.OverdriveCommands, overdrive.OverdriveCommand, nil),
		User:               nameToNamedAPIResource(cfg, cfg.E.CharacterClasses, overdrive.User, nil),
		RelatedStats:       namesToNamedAPIResources(cfg, cfg.E.Stats, ability.RelatedStats),
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (cfg *Config) retrieveOverdriveAbilities(r *http.Request, i handlerInput[seeding.OverdriveAbility, OverdriveAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}
	abilityType := database.AbilityTypeOverdriveAbility

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, cfg.T.AttackType, ids, QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		i.queryEnumList(cfg, r, cfg.T.TargetType, ids, QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		i.queryEnum(r, cfg.T.DamageFormula, ids, QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		i.queryIntList(cfg, r, ids, QpnRank, getTypedAbilityIDsByRank(cfg, abilityType)),
		i.queryNameIDListNul(cfg, r, ids, QpnElement, cfg.E.Elements.resTypeSingle, cfg.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		i.queryNameID(r, ids, QpnUser, cfg.E.CharacterClasses.resTypeSingle, cfg.L.CharClasses, cfg.db.GetOverdriveAbilityIDsByCharClass),
		i.queryNameID(r, ids, QpnRelatedStat, cfg.E.Stats.resTypeSingle, cfg.L.Stats, cfg.db.GetOverdriveAbilityIDsByRelatedStat),
		i.queryIDNul(r, ids, QpnStatusInflict, cfg.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		i.queryIDNul(r, ids, QpnStatusRemove, cfg.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		i.queryBool2(r, ids, QpnCanCrit, getTypedAbilityIDsCanCrit(cfg, abilityType)),
		i.queryBool2(r, ids, QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
		i.queryBool2(r, ids, QpnStatChanges, getTypedAbilityIDsWithStatChanges(cfg, abilityType)),
		i.queryBool2(r, ids, QpnModChanges, getTypedAbilityIDsWithModifierChanges(cfg, abilityType)),
	})
}
