package api

func verifyCritChanceParams(v paramsVerifyer, params CritChanceParams) (CritChanceParams, error) {
	var err error

	params.Ability, err = v.verifyParamFieldObj(params.Ability, PfnAbility, vfAbilityCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	params.User, err = v.verifyParamFieldObj(params.User, PfnUser, vfUserCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	params.Target, err = v.verifyParamFieldObj(params.Target, PfnTarget, vfTargetCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	return params, nil
}

func vfAbilityCritChance(v paramsVerifyer, item AbilityCritChance, _ FieldName) (AbilityCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return AbilityCritChance{}, err
	}

	item.Lookup, err = v.verifyParamFieldObjPtr(item.Lookup, PfnLookup, vfAbilityLookupCritChance)
	if err != nil {
		return AbilityCritChance{}, err
	}

	item.Custom, err = v.verifyParamFieldObjPtr(item.Custom, PfnCustom, vfAbilityCustomCritChance)
	if err != nil {
		return AbilityCritChance{}, err
	}

	return item, nil
}

func vfAbilityLookupCritChance(v paramsVerifyer, item AbilityLookupCritChance, _ FieldName) (AbilityLookupCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.AbilityID, err = v.vfIDField(item.AbilityID, PfnAbilityID)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.Interaction, err = v.verifyParamFieldPtr(item.Interaction, PfnInteraction, vfID)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.Overrides, err = v.verifyParamFieldObjPtr(item.Overrides, PfnOverrides, vfAbilityOverridesCritChance)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	return item, nil
}

func vfAbilityOverridesCritChance(v paramsVerifyer, item AbilityOverridesCritChance, _ FieldName) (AbilityOverridesCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	item.Critical, err = v.verifyParamFieldPtr(item.Critical, PfnCritical, vfEnum)
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = v.verifyParamFieldPtr(item.CriticalPlusVal, PfnCriticalPlusVal, vfInt)
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	return item, nil
}

func vfAbilityCustomCritChance(v paramsVerifyer, item AbilityCustomCritChance, _ FieldName) (AbilityCustomCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	item.Critical, err = v.vfEnumField(item.Critical, PfnCritical)
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = v.vfIntField(item.CriticalPlusVal, PfnCriticalPlusVal)
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	return item, nil
}

func vfUserCritChance(v paramsVerifyer, item UserCritChance, _ FieldName) (UserCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return UserCritChance{}, err
	}

	item.Lookup, err = v.verifyParamFieldObjPtr(item.Lookup, PfnLookup, vfUserLookupCritChance)
	if err != nil {
		return UserCritChance{}, err
	}

	item.Custom, err = v.verifyParamFieldObjPtr(item.Custom, PfnCustom, vfUserCustomCritChance)
	if err != nil {
		return UserCritChance{}, err
	}

	item.LuckStacks, err = v.vfIntField(item.LuckStacks, PfnLuckStacks)
	if err != nil {
		return UserCritChance{}, err
	}

	return item, nil
}

func vfUserLookupCritChance(v paramsVerifyer, item UserLookupCritChance, _ FieldName) (UserLookupCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return UserLookupCritChance{}, err
	}

	item.MonsterID, err = v.vfIDField(item.MonsterID, PfnMonsterID)
	if err != nil {
		return UserLookupCritChance{}, err
	}

	item.Overrides, err = v.verifyParamFieldObjPtr(item.Overrides, PfnOverrides, vfUserOverridesCritChance)
	if err != nil {
		return UserLookupCritChance{}, err
	}

	return item, nil
}

func vfUserOverridesCritChance(v paramsVerifyer, item UserOverridesCritChance, _ FieldName) (UserOverridesCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	item.Luck, err = v.verifyParamFieldPtr(item.Luck, PfnLuck, vfInt)
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = v.verifyParamFieldPtr(item.CriticalPlusVal, PfnCriticalPlusVal, vfInt)
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	return item, nil
}

func vfUserCustomCritChance(v paramsVerifyer, item UserCustomCritChance, _ FieldName) (UserCustomCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return UserCustomCritChance{}, err
	}

	item.Luck, err = v.vfIntField(item.Luck, PfnLuck)
	if err != nil {
		return UserCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = v.vfIntField(item.CriticalPlusVal, PfnCriticalPlusVal)
	if err != nil {
		return UserCustomCritChance{}, err
	}

	return item, nil
}

func vfTargetCritChance(v paramsVerifyer, item TargetCritChance, _ FieldName) (TargetCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return TargetCritChance{}, err
	}

	item.Lookup, err = v.verifyParamFieldObjPtr(item.Lookup, PfnLookup, vfTargetLookupCritChance)
	if err != nil {
		return TargetCritChance{}, err
	}

	item.Custom, err = v.verifyParamFieldObjPtr(item.Custom, PfnCustom, vfTargetCustomCritChance)
	if err != nil {
		return TargetCritChance{}, err
	}

	item.JinxStacks, err = v.vfIntField(item.JinxStacks, PfnJinxStacks)
	if err != nil {
		return TargetCritChance{}, err
	}

	return item, nil
}

func vfTargetLookupCritChance(v paramsVerifyer, item TargetLookupCritChance, _ FieldName) (TargetLookupCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	item.MonsterID, err = v.vfIDField(item.MonsterID, PfnMonsterID)
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	item.Overrides, err = v.verifyParamFieldObjPtr(item.Overrides, PfnOverrides, vfTargetOverridesCritChance)
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	return item, nil
}

func vfTargetOverridesCritChance(v paramsVerifyer, item TargetOverridesCritChance, _ FieldName) (TargetOverridesCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return TargetOverridesCritChance{}, err
	}

	item.Luck, err = v.verifyParamFieldPtr(item.Luck, PfnLuck, vfInt)
	if err != nil {
		return TargetOverridesCritChance{}, err
	}

	return item, nil
}

func vfTargetCustomCritChance(v paramsVerifyer, item TargetCustomCritChance, _ FieldName) (TargetCustomCritChance, error) {
	err := v.vfExistingFields()
	if err != nil {
		return TargetCustomCritChance{}, err
	}

	item.Luck, err = v.vfIntField(item.Luck, PfnLuck)
	if err != nil {
		return TargetCustomCritChance{}, err
	}

	return item, nil
}
