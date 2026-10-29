package api

func verifyCritChanceParams(cfg *Config, params CritChanceParams, paramsRules ParamsRules, userValues map[FieldName]any) (CritChanceParams, error) {
	var err error

	params.Ability, err = verifyParamFieldObj(cfg, params.Ability, PfnAbility, userValues, paramsRules, vfAbilityCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	params.User, err = verifyParamFieldObj(cfg, params.User, PfnUser, userValues, paramsRules, vfUserCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	params.Target, err = verifyParamFieldObj(cfg, params.Target, PfnTarget, userValues, paramsRules, vfTargetCritChance)
	if err != nil {
		return CritChanceParams{}, err
	}

	return params, nil
}



func vfAbilityCritChance(cfg *Config, item AbilityCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (AbilityCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return AbilityCritChance{}, err
	}

	item.Lookup, err = verifyParamFieldObjPtr(cfg, item.Lookup, PfnLookup, userValues, paramsRules, vfAbilityLookupCritChance)
	if err != nil {
		return AbilityCritChance{}, err
	}

	item.Custom, err = verifyParamFieldObjPtr(cfg, item.Custom, PfnCustom, userValues, paramsRules, vfAbilityCustomCritChance)
	if err != nil {
		return AbilityCritChance{}, err
	}

	return item, nil
}


func vfAbilityLookupCritChance(cfg *Config, item AbilityLookupCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (AbilityLookupCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.AbilityID, err = verifyParamField(cfg, item.AbilityID, PfnAbilityID, userValues, paramsRules, vfId)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.Interaction, err = verifyParamFieldPtr(cfg, item.Interaction, PfnInteraction, userValues, paramsRules, vfId)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	item.Overrides, err = verifyParamFieldObjPtr(cfg, item.Overrides, PfnOverrides, userValues, paramsRules, vfAbilityOverridesCritChance)
	if err != nil {
		return AbilityLookupCritChance{}, err
	}

	return item, nil
}


func vfAbilityOverridesCritChance(cfg *Config, item AbilityOverridesCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (AbilityOverridesCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	item.Critical, err = verifyParamFieldPtr(cfg, item.Critical, PfnCritical, userValues, paramsRules, vfEnum)
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = verifyParamFieldPtr(cfg, item.CriticalPlusVal, PfnCriticalPlusVal, userValues, paramsRules, vfInt)
	if err != nil {
		return AbilityOverridesCritChance{}, err
	}

	return item, nil
}

func vfAbilityCustomCritChance(cfg *Config, item AbilityCustomCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (AbilityCustomCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	item.Critical, err = verifyParamField(cfg, item.Critical, PfnCritical, userValues, paramsRules, vfEnum)
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = verifyParamField(cfg, item.CriticalPlusVal, PfnCriticalPlusVal, userValues, paramsRules, vfInt)
	if err != nil {
		return AbilityCustomCritChance{}, err
	}

	return item, nil
}


func vfUserCritChance(cfg *Config, item UserCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (UserCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return UserCritChance{}, err
	}

	item.Lookup, err = verifyParamFieldObjPtr(cfg, item.Lookup, PfnLookup, userValues, paramsRules, vfUserLookupCritChance)
	if err != nil {
		return UserCritChance{}, err
	}

	item.Custom, err = verifyParamFieldObjPtr(cfg, item.Custom, PfnCustom, userValues, paramsRules, vfUserCustomCritChance)
	if err != nil {
		return UserCritChance{}, err
	}

	item.LuckStacks, err = verifyParamField(cfg, item.LuckStacks, PfnLuckStacks, userValues, paramsRules, vfInt)
	if err != nil {
		return UserCritChance{}, err
	}

	return item, nil
}


func vfUserLookupCritChance(cfg *Config, item UserLookupCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (UserLookupCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return UserLookupCritChance{}, err
	}

	item.MonsterID, err = verifyParamField(cfg, item.MonsterID, PfnMonsterID, userValues, paramsRules, vfId)
	if err != nil {
		return UserLookupCritChance{}, err
	}

	item.Overrides, err = verifyParamFieldObjPtr(cfg, item.Overrides, PfnOverrides, userValues, paramsRules, vfUserOverridesCritChance)
	if err != nil {
		return UserLookupCritChance{}, err
	}

	return item, nil
}


func vfUserOverridesCritChance(cfg *Config, item UserOverridesCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (UserOverridesCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	item.Luck, err = verifyParamFieldPtr(cfg, item.Luck, PfnLuck, userValues, paramsRules, vfInt)
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	item.CriticalPlusVal, err = verifyParamFieldPtr(cfg, item.CriticalPlusVal, PfnCriticalPlusVal, userValues, paramsRules, vfInt)
	if err != nil {
		return UserOverridesCritChance{}, err
	}

	return item, nil
}

func vfUserCustomCritChance(cfg *Config, item UserCustomCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (UserCustomCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return UserCustomCritChance{}, err
	}

	item.Luck, err = verifyParamField(cfg, item.Luck, PfnLuck, userValues, paramsRules, vfInt)
	if err != nil {
		return UserCustomCritChance{}, err
	}

	item.CriticalPlusVal, err = verifyParamField(cfg, item.CriticalPlusVal, PfnCriticalPlusVal, userValues, paramsRules, vfInt)
	if err != nil {
		return UserCustomCritChance{}, err
	}

	return item, nil
}


func vfTargetCritChance(cfg *Config, item TargetCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (TargetCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return TargetCritChance{}, err
	}

	item.Lookup, err = verifyParamFieldObjPtr(cfg, item.Lookup, PfnLookup, userValues, paramsRules, vfTargetLookupCritChance)
	if err != nil {
		return TargetCritChance{}, err
	}

	item.Custom, err = verifyParamFieldObjPtr(cfg, item.Custom, PfnCustom, userValues, paramsRules, vfTargetCustomCritChance)
	if err != nil {
		return TargetCritChance{}, err
	}

	item.JinxStacks, err = verifyParamField(cfg, item.JinxStacks, PfnJinxStacks, userValues, paramsRules, vfInt)
	if err != nil {
		return TargetCritChance{}, err
	}

	return item, nil
}


func vfTargetLookupCritChance(cfg *Config, item TargetLookupCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (TargetLookupCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	item.MonsterID, err = verifyParamField(cfg, item.MonsterID, PfnMonsterID, userValues, paramsRules, vfId)
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	item.Overrides, err = verifyParamFieldObjPtr(cfg, item.Overrides, PfnOverrides, userValues, paramsRules, vfTargetOverridesCritChance)
	if err != nil {
		return TargetLookupCritChance{}, err
	}

	return item, nil
}


func vfTargetOverridesCritChance(cfg *Config, item TargetOverridesCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (TargetOverridesCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return TargetOverridesCritChance{}, err
	}

	item.Luck, err = verifyParamFieldPtr(cfg, item.Luck, PfnLuck, userValues, paramsRules, vfInt)
	if err != nil {
		return TargetOverridesCritChance{}, err
	}

	return item, nil
}

func vfTargetCustomCritChance(cfg *Config, item TargetCustomCritChance, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (TargetCustomCritChance, error) {
	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return TargetCustomCritChance{}, err
	}

	item.Luck, err = verifyParamField(cfg, item.Luck, PfnLuck, userValues, paramsRules, vfInt)
	if err != nil {
		return TargetCustomCritChance{}, err
	}

	return item, nil
}