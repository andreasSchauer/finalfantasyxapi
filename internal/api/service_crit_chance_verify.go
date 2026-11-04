package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

func verifyCritChanceParams(v paramsVerifyer, params CritChanceParams) (CritChanceParams, error) {
	var err error

	params.Ability, err = v.obj(params.Ability, typestrings.PfnAbility, vfAbilityCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	params.User, err = v.obj(params.User, typestrings.PfnUser, vfUserCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	params.Target, err = v.obj(params.Target, typestrings.PfnTarget, vfTargetCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	return params, nil
}

func vfAbilityCritChance(v paramsVerifyer, item AbilityCritChance, _ typestrings.FieldName) (AbilityCritChance, error) {
	var err error

	item.Lookup, err = v.objPtr(item.Lookup, typestrings.PfnLookup, vfAbilityLookupCritChance)
	if err != nil {
		return AbilityCritChance{}, err
	}

	item.Custom, err = v.objPtr(item.Custom, typestrings.PfnCustom, vfAbilityCustomCritChance)
	if err != nil {
		return AbilityCritChance{}, err
	}

	return item, nil
}

func vfAbilityLookupCritChance(v paramsVerifyer, item AbilityLookupCritChance, _ typestrings.FieldName) (AbilityLookupCritChance, error) {
	var err error

	item.AbilityID, err = v.id(item.AbilityID, typestrings.PfnAbilityID)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.Interaction, err = v.intPtr(item.Interaction, typestrings.PfnInteraction)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.Overrides, err = v.objPtr(item.Overrides, typestrings.PfnOverrides, vfAbilityOverridesCritChance)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	return item, nil
}

func vfAbilityOverridesCritChance(v paramsVerifyer, item AbilityOverridesCritChance, _ typestrings.FieldName) (AbilityOverridesCritChance, error) {
	var err error

	item.Critical, err = v.enumPtr(item.Critical, typestrings.PfnCritical)
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = v.intPtr(item.CriticalPlusVal, typestrings.PfnCriticalPlusVal)
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	return item, nil
}

func vfAbilityCustomCritChance(v paramsVerifyer, item AbilityCustomCritChance, _ typestrings.FieldName) (AbilityCustomCritChance, error) {
	var err error

	item.Critical, err = v.enum(item.Critical, typestrings.PfnCritical)
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = v.int(item.CriticalPlusVal, typestrings.PfnCriticalPlusVal)
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	return item, nil
}

func vfUserCritChance(v paramsVerifyer, item UserCritChance, _ typestrings.FieldName) (UserCritChance, error) {
	var err error

	item.Lookup, err = v.objPtr(item.Lookup, typestrings.PfnLookup, vfUserLookupCritChance)
	if err != nil {
		return UserCritChance{}, err
	}

	item.Custom, err = v.objPtr(item.Custom, typestrings.PfnCustom, vfUserCustomCritChance)
	if err != nil {
		return UserCritChance{}, err
	}

	item.LuckStacks, err = v.int(item.LuckStacks, typestrings.PfnLuckStacks)
	if err != nil {
		return UserCritChance{}, err
	}

	return item, nil
}

func vfUserLookupCritChance(v paramsVerifyer, item UserLookupCritChance, _ typestrings.FieldName) (UserLookupCritChance, error) {
	var err error

	item.MonsterID, err = v.id(item.MonsterID, typestrings.PfnMonsterID)
	if err != nil {
		return UserLookupCritChance{}, err
	}

	item.Overrides, err = v.objPtr(item.Overrides, typestrings.PfnOverrides, vfUserOverridesCritChance)
	if err != nil {
		return UserLookupCritChance{}, err
	}

	return item, nil
}

func vfUserOverridesCritChance(v paramsVerifyer, item UserOverridesCritChance, _ typestrings.FieldName) (UserOverridesCritChance, error) {
	var err error

	item.Luck, err = v.intPtr(item.Luck, typestrings.PfnLuck)
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = v.intPtr(item.CriticalPlusVal, typestrings.PfnCriticalPlusVal)
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	return item, nil
}

func vfUserCustomCritChance(v paramsVerifyer, item UserCustomCritChance, _ typestrings.FieldName) (UserCustomCritChance, error) {
	var err error

	item.Luck, err = v.int(item.Luck, typestrings.PfnLuck)
	if err != nil {
		return UserCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = v.int(item.CriticalPlusVal, typestrings.PfnCriticalPlusVal)
	if err != nil {
		return UserCustomCritChance{}, err
	}

	return item, nil
}

func vfTargetCritChance(v paramsVerifyer, item TargetCritChance, _ typestrings.FieldName) (TargetCritChance, error) {
	var err error

	item.Lookup, err = v.objPtr(item.Lookup, typestrings.PfnLookup, vfTargetLookupCritChance)
	if err != nil {
		return TargetCritChance{}, err
	}

	item.Custom, err = v.objPtr(item.Custom, typestrings.PfnCustom, vfTargetCustomCritChance)
	if err != nil {
		return TargetCritChance{}, err
	}

	item.JinxStacks, err = v.int(item.JinxStacks, typestrings.PfnJinxStacks)
	if err != nil {
		return TargetCritChance{}, err
	}

	return item, nil
}

func vfTargetLookupCritChance(v paramsVerifyer, item TargetLookupCritChance, _ typestrings.FieldName) (TargetLookupCritChance, error) {
	var err error

	item.MonsterID, err = v.id(item.MonsterID, typestrings.PfnMonsterID)
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	item.Overrides, err = v.objPtr(item.Overrides, typestrings.PfnOverrides, vfTargetOverridesCritChance)
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	return item, nil
}

func vfTargetOverridesCritChance(v paramsVerifyer, item TargetOverridesCritChance, _ typestrings.FieldName) (TargetOverridesCritChance, error) {
	var err error

	item.Luck, err = v.intPtr(item.Luck, typestrings.PfnLuck)
	if err != nil {
		return TargetOverridesCritChance{}, err
	}

	return item, nil
}

func vfTargetCustomCritChance(v paramsVerifyer, item TargetCustomCritChance, _ typestrings.FieldName) (TargetCustomCritChance, error) {
	var err error

	item.Luck, err = v.int(item.Luck, typestrings.PfnLuck)
	if err != nil {
		return TargetCustomCritChance{}, err
	}

	return item, nil
}
