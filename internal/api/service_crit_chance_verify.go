package api

func verifyCritChanceParams(v paramsVerifyer, params CritChanceParams) (CritChanceParams, error) {
	var err error

	params.Ability, err = v.obj(params.Ability, PfnAbility, vfAbilityCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	params.User, err = v.obj(params.User, PfnUser, vfUserCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	params.Target, err = v.obj(params.Target, PfnTarget, vfTargetCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	return params, nil
}

func vfAbilityCritChance(v paramsVerifyer, item AbilityCritChance, _ FieldName) (AbilityCritChance, error) {
	var err error

	item.Lookup, err = v.objPtr(item.Lookup, PfnLookup, vfAbilityLookupCritChance)
	if err != nil {
		return AbilityCritChance{}, err
	}

	item.Custom, err = v.objPtr(item.Custom, PfnCustom, vfAbilityCustomCritChance)
	if err != nil {
		return AbilityCritChance{}, err
	}

	return item, nil
}

func vfAbilityLookupCritChance(v paramsVerifyer, item AbilityLookupCritChance, _ FieldName) (AbilityLookupCritChance, error) {
	var err error

	item.AbilityID, err = v.id(item.AbilityID, PfnAbilityID)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.Interaction, err = v.intPtr(item.Interaction, PfnInteraction)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.Overrides, err = v.objPtr(item.Overrides, PfnOverrides, vfAbilityOverridesCritChance)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	return item, nil
}

func vfAbilityOverridesCritChance(v paramsVerifyer, item AbilityOverridesCritChance, _ FieldName) (AbilityOverridesCritChance, error) {
	var err error

	item.Critical, err = v.enumPtr(item.Critical, PfnCritical)
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = v.intPtr(item.CriticalPlusVal, PfnCriticalPlusVal)
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	return item, nil
}

func vfAbilityCustomCritChance(v paramsVerifyer, item AbilityCustomCritChance, _ FieldName) (AbilityCustomCritChance, error) {
	var err error

	item.Critical, err = v.enum(item.Critical, PfnCritical)
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = v.int(item.CriticalPlusVal, PfnCriticalPlusVal)
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	return item, nil
}

func vfUserCritChance(v paramsVerifyer, item UserCritChance, _ FieldName) (UserCritChance, error) {
	var err error

	item.Lookup, err = v.objPtr(item.Lookup, PfnLookup, vfUserLookupCritChance)
	if err != nil {
		return UserCritChance{}, err
	}

	item.Custom, err = v.objPtr(item.Custom, PfnCustom, vfUserCustomCritChance)
	if err != nil {
		return UserCritChance{}, err
	}

	item.LuckStacks, err = v.int(item.LuckStacks, PfnLuckStacks)
	if err != nil {
		return UserCritChance{}, err
	}

	return item, nil
}

func vfUserLookupCritChance(v paramsVerifyer, item UserLookupCritChance, _ FieldName) (UserLookupCritChance, error) {
	var err error

	item.MonsterID, err = v.id(item.MonsterID, PfnMonsterID)
	if err != nil {
		return UserLookupCritChance{}, err
	}

	item.Overrides, err = v.objPtr(item.Overrides, PfnOverrides, vfUserOverridesCritChance)
	if err != nil {
		return UserLookupCritChance{}, err
	}

	return item, nil
}

func vfUserOverridesCritChance(v paramsVerifyer, item UserOverridesCritChance, _ FieldName) (UserOverridesCritChance, error) {
	var err error

	item.Luck, err = v.intPtr(item.Luck, PfnLuck)
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = v.intPtr(item.CriticalPlusVal, PfnCriticalPlusVal)
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	return item, nil
}

func vfUserCustomCritChance(v paramsVerifyer, item UserCustomCritChance, _ FieldName) (UserCustomCritChance, error) {
	var err error

	item.Luck, err = v.int(item.Luck, PfnLuck)
	if err != nil {
		return UserCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = v.int(item.CriticalPlusVal, PfnCriticalPlusVal)
	if err != nil {
		return UserCustomCritChance{}, err
	}

	return item, nil
}

func vfTargetCritChance(v paramsVerifyer, item TargetCritChance, _ FieldName) (TargetCritChance, error) {
	var err error

	item.Lookup, err = v.objPtr(item.Lookup, PfnLookup, vfTargetLookupCritChance)
	if err != nil {
		return TargetCritChance{}, err
	}

	item.Custom, err = v.objPtr(item.Custom, PfnCustom, vfTargetCustomCritChance)
	if err != nil {
		return TargetCritChance{}, err
	}

	item.JinxStacks, err = v.int(item.JinxStacks, PfnJinxStacks)
	if err != nil {
		return TargetCritChance{}, err
	}

	return item, nil
}

func vfTargetLookupCritChance(v paramsVerifyer, item TargetLookupCritChance, _ FieldName) (TargetLookupCritChance, error) {
	var err error

	item.MonsterID, err = v.id(item.MonsterID, PfnMonsterID)
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	item.Overrides, err = v.objPtr(item.Overrides, PfnOverrides, vfTargetOverridesCritChance)
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	return item, nil
}

func vfTargetOverridesCritChance(v paramsVerifyer, item TargetOverridesCritChance, _ FieldName) (TargetOverridesCritChance, error) {
	var err error

	item.Luck, err = v.intPtr(item.Luck, PfnLuck)
	if err != nil {
		return TargetOverridesCritChance{}, err
	}

	return item, nil
}

func vfTargetCustomCritChance(v paramsVerifyer, item TargetCustomCritChance, _ FieldName) (TargetCustomCritChance, error) {
	var err error

	item.Luck, err = v.int(item.Luck, PfnLuck)
	if err != nil {
		return TargetCustomCritChance{}, err
	}

	return item, nil
}
