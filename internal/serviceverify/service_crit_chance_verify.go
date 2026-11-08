package serviceverify

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func VerifyCritChanceParams(v ParamsVerifyer, params servicetypes.CritChanceParams) (servicetypes.CritChanceParams, error) {
	var err error

	params.Ability, err = v.Obj(params.Ability, ts.PfnAbility, vfAbilityCritChance)
	if err != nil {
		return servicetypes.CritChanceParams{}, err
	}

	params.User, err = v.Obj(params.User, ts.PfnUser, vfUserCritChance)
	if err != nil {
		return servicetypes.CritChanceParams{}, err
	}

	params.Target, err = v.Obj(params.Target, ts.PfnTarget, vfTargetCritChance)
	if err != nil {
		return servicetypes.CritChanceParams{}, err
	}

	return params, nil
}

func vfAbilityCritChance(v ParamsVerifyer, item servicetypes.AbilityCritChance, _ ts.FieldName) (servicetypes.AbilityCritChance, error) {
	var err error

	item.Lookup, err = v.ObjPtr(item.Lookup, ts.PfnLookup, vfAbilityLookupCritChance)
	if err != nil {
		return servicetypes.AbilityCritChance{}, err
	}

	item.Custom, err = v.ObjPtr(item.Custom, ts.PfnCustom, vfAbilityCustomCritChance)
	if err != nil {
		return servicetypes.AbilityCritChance{}, err
	}

	return item, nil
}

func vfAbilityLookupCritChance(v ParamsVerifyer, item servicetypes.AbilityLookupCritChance, _ ts.FieldName) (servicetypes.AbilityLookupCritChance, error) {
	var err error

	item.AbilityID, err = v.ID(item.AbilityID, ts.PfnAbilityID)
	if err != nil {
		return servicetypes.AbilityLookupCritChance{}, err
	}

	item.Interaction, err = v.IntPtr(item.Interaction, ts.PfnInteraction)
	if err != nil {
		return servicetypes.AbilityLookupCritChance{}, err
	}

	item.Overrides, err = v.ObjPtr(item.Overrides, ts.PfnOverrides, vfAbilityOverridesCritChance)
	if err != nil {
		return servicetypes.AbilityLookupCritChance{}, err
	}

	return item, nil
}

func vfAbilityOverridesCritChance(v ParamsVerifyer, item servicetypes.AbilityOverridesCritChance, _ ts.FieldName) (servicetypes.AbilityOverridesCritChance, error) {
	var err error

	item.Critical, err = v.EnumPtr(item.Critical, ts.PfnCritical)
	if err != nil {
		return servicetypes.AbilityOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = v.IntPtr(item.CriticalPlusVal, ts.PfnCriticalPlusVal)
	if err != nil {
		return servicetypes.AbilityOverridesCritChance{}, err
	}

	return item, nil
}

func vfAbilityCustomCritChance(v ParamsVerifyer, item servicetypes.AbilityCustomCritChance, _ ts.FieldName) (servicetypes.AbilityCustomCritChance, error) {
	var err error

	item.Critical, err = v.Enum(item.Critical, ts.PfnCritical)
	if err != nil {
		return servicetypes.AbilityCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = v.Int(item.CriticalPlusVal, ts.PfnCriticalPlusVal)
	if err != nil {
		return servicetypes.AbilityCustomCritChance{}, err
	}

	return item, nil
}

func vfUserCritChance(v ParamsVerifyer, item servicetypes.UserCritChance, _ ts.FieldName) (servicetypes.UserCritChance, error) {
	var err error

	item.Lookup, err = v.ObjPtr(item.Lookup, ts.PfnLookup, vfUserLookupCritChance)
	if err != nil {
		return servicetypes.UserCritChance{}, err
	}

	item.Custom, err = v.ObjPtr(item.Custom, ts.PfnCustom, vfUserCustomCritChance)
	if err != nil {
		return servicetypes.UserCritChance{}, err
	}

	item.LuckStacks, err = v.Int(item.LuckStacks, ts.PfnLuckStacks)
	if err != nil {
		return servicetypes.UserCritChance{}, err
	}

	return item, nil
}

func vfUserLookupCritChance(v ParamsVerifyer, item servicetypes.UserLookupCritChance, _ ts.FieldName) (servicetypes.UserLookupCritChance, error) {
	var err error

	item.MonsterID, err = v.ID(item.MonsterID, ts.PfnMonsterID)
	if err != nil {
		return servicetypes.UserLookupCritChance{}, err
	}

	item.Overrides, err = v.ObjPtr(item.Overrides, ts.PfnOverrides, vfUserOverridesCritChance)
	if err != nil {
		return servicetypes.UserLookupCritChance{}, err
	}

	return item, nil
}

func vfUserOverridesCritChance(v ParamsVerifyer, item servicetypes.UserOverridesCritChance, _ ts.FieldName) (servicetypes.UserOverridesCritChance, error) {
	var err error

	item.Luck, err = v.IntPtr(item.Luck, ts.PfnLuck)
	if err != nil {
		return servicetypes.UserOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = v.IntPtr(item.CriticalPlusVal, ts.PfnCriticalPlusVal)
	if err != nil {
		return servicetypes.UserOverridesCritChance{}, err
	}

	return item, nil
}

func vfUserCustomCritChance(v ParamsVerifyer, item servicetypes.UserCustomCritChance, _ ts.FieldName) (servicetypes.UserCustomCritChance, error) {
	var err error

	item.Luck, err = v.Int(item.Luck, ts.PfnLuck)
	if err != nil {
		return servicetypes.UserCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = v.Int(item.CriticalPlusVal, ts.PfnCriticalPlusVal)
	if err != nil {
		return servicetypes.UserCustomCritChance{}, err
	}

	return item, nil
}

func vfTargetCritChance(v ParamsVerifyer, item servicetypes.TargetCritChance, _ ts.FieldName) (servicetypes.TargetCritChance, error) {
	var err error

	item.Lookup, err = v.ObjPtr(item.Lookup, ts.PfnLookup, vfTargetLookupCritChance)
	if err != nil {
		return servicetypes.TargetCritChance{}, err
	}

	item.Custom, err = v.ObjPtr(item.Custom, ts.PfnCustom, vfTargetCustomCritChance)
	if err != nil {
		return servicetypes.TargetCritChance{}, err
	}

	item.JinxStacks, err = v.Int(item.JinxStacks, ts.PfnJinxStacks)
	if err != nil {
		return servicetypes.TargetCritChance{}, err
	}

	return item, nil
}

func vfTargetLookupCritChance(v ParamsVerifyer, item servicetypes.TargetLookupCritChance, _ ts.FieldName) (servicetypes.TargetLookupCritChance, error) {
	var err error

	item.MonsterID, err = v.ID(item.MonsterID, ts.PfnMonsterID)
	if err != nil {
		return servicetypes.TargetLookupCritChance{}, err
	}

	item.Overrides, err = v.ObjPtr(item.Overrides, ts.PfnOverrides, vfTargetOverridesCritChance)
	if err != nil {
		return servicetypes.TargetLookupCritChance{}, err
	}

	return item, nil
}

func vfTargetOverridesCritChance(v ParamsVerifyer, item servicetypes.TargetOverridesCritChance, _ ts.FieldName) (servicetypes.TargetOverridesCritChance, error) {
	var err error

	item.Luck, err = v.IntPtr(item.Luck, ts.PfnLuck)
	if err != nil {
		return servicetypes.TargetOverridesCritChance{}, err
	}

	return item, nil
}

func vfTargetCustomCritChance(v ParamsVerifyer, item servicetypes.TargetCustomCritChance, _ ts.FieldName) (servicetypes.TargetCustomCritChance, error) {
	var err error

	item.Luck, err = v.Int(item.Luck, ts.PfnLuck)
	if err != nil {
		return servicetypes.TargetCustomCritChance{}, err
	}

	return item, nil
}
