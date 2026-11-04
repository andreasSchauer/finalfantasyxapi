package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getEnemyAbility(r *http.Request, i handlerInput[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource], id int32) (EnemyAbility, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return EnemyAbility{}, err
	}

	monsters, err := cfg.E.Monsters.dbGetRelations(cfg, r.Context(), ability, cfg.db.GetEnemyAbilityMonsterIDs)
	if err != nil {
		return EnemyAbility{}, err
	}

	response := EnemyAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.E.Abilities, ability.Ability.ID),
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Monsters:           monsters,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (cfg *Config) retrieveEnemyAbilities(r *http.Request, i handlerInput[seedtypes.EnemyAbility, EnemyAbility, NamedAPIResource]) ([]int32, error) {
	abilityType := database.AbilityTypeEnemyAbility
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, typestrings.QpnMonster, lookup.L.Monsters, cfg.db.GetEnemyAbilityIDsByMonster),
		i.queryEnumList(cfg, r, enums.T.DamageType, ids, typestrings.QpnDamageType, getTypedAbilityIDsByDamageType(cfg, abilityType)),
		i.queryEnumList(cfg, r, enums.T.AttackType, ids, typestrings.QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		i.queryEnumList(cfg, r, enums.T.TargetType, ids, typestrings.QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		i.queryEnum(r, enums.T.DamageFormula, ids, typestrings.QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		i.queryIntList(cfg, r, ids, typestrings.QpnRank, getTypedAbilityIDsByRank(cfg, abilityType)),
		i.queryNameIDListNul(cfg, r, ids, typestrings.QpnElement, cfg.E.Elements.resTypeSingle, lookup.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		i.queryIDNul(r, ids, typestrings.QpnStatusInflict, lookup.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		i.queryIDNul(r, ids, typestrings.QpnStatusRemove, lookup.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		i.queryBool(r, ids, typestrings.QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnCanCrit, getTypedAbilityIDsCanCrit(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnBDL, getTypedAbilityIDsBreakDmgLimit(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnDarkable, getTypedAbilityIDsDarkable(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnSilenceable, getTypedAbilityIDsSilenceable(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnReflectable, getTypedAbilityIDsReflectable(cfg, abilityType)),
		i.queryBool2(r, ids, typestrings.QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
	})
}
