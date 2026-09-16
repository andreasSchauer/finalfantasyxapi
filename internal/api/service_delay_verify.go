package api

func verifyDelayParams(cfg *Config, params DelayParams, valueMap map[FieldName]any) (DelayParams, error) {
	valTree := compileValidationTree(cfg.getDelayParamsDoc().Fields)

	err := vfExistingFields(valueMap, valTree)
	if err != nil {
		return DelayParams{}, err
	}

	params.IgnImmunities, err = verifyParamField(cfg, params.IgnImmunities, pfnIgnImmunities, valueMap, valTree, nil)
	if err != nil {
		return DelayParams{}, err
	}

	params.Delay, err = verifyParamFieldObj(cfg, params.Delay, pfnDelay, valueMap, valTree, vfCustomDelay)
	if err != nil {
		return DelayParams{}, err
	}

	params.Target, err = verifyParamFieldObj(cfg, params.Target, pfnTarget, valueMap, valTree, vfDelayTarget)
	if err != nil {
		return DelayParams{}, err
	}

	return params, nil
}

func vfCustomDelay(cfg *Config, delay CustomDelay, _ FieldName, valueMap map[FieldName]any, valTree ValidationTree) (CustomDelay, error) {
	var err error

	err = vfExistingFields(valueMap, valTree)
	if err != nil {
		return CustomDelay{}, err
	}

	delay.DelayType, err = verifyParamField(cfg, delay.DelayType, pfnDelayType, valueMap, valTree, vfEnum)
	if err != nil {
		return CustomDelay{}, err
	}

	delay.AttackType, err = verifyParamField(cfg, delay.AttackType, pfnAttackType, valueMap, valTree, vfEnum)
	if err != nil {
		return CustomDelay{}, err
	}

	delay.Strength, err = verifyParamFieldPtr(cfg, delay.Strength, pfnStrength, valueMap, valTree, vfEnum)
	if err != nil {
		return CustomDelay{}, err
	}

	delay.DelayConstant, err = verifyParamFieldPtr(cfg, delay.DelayConstant, pfnDelayConstant, valueMap, valTree, vfIntId)
	if err != nil {
		return CustomDelay{}, err
	}

	return delay, nil
}

func vfDelayTarget(cfg *Config, target DelayTarget, _ FieldName, valueMap map[FieldName]any, valTree ValidationTree) (DelayTarget, error) {
	var err error

	err = vfExistingFields(valueMap, valTree)
	if err != nil {
		return DelayTarget{}, err
	}

	target.MonsterID, err = verifyParamFieldPtr(cfg, target.MonsterID, pfnMonsterID, valueMap, valTree, vfIntId)
	if err != nil {
		return DelayTarget{}, err
	}

	target.AltState, err = verifyParamFieldPtr(cfg, target.AltState, pfnAltState, valueMap, valTree, vfIntId)
	if err != nil {
		return DelayTarget{}, err
	}

	target.Agility, err = verifyParamFieldPtr(cfg, target.Agility, pfnAgility, valueMap, valTree, vfIntId)
	if err != nil {
		return DelayTarget{}, err
	}

	target.Status, err = verifyParamFieldPtr(cfg, target.Status, pfnStatus, valueMap, valTree, vfEnum)
	if err != nil {
		return DelayTarget{}, err
	}

	target.RemainingTicks, err = verifyParamFieldPtr(cfg, target.RemainingTicks, pfnRemainingTicks, valueMap, valTree, vfIntId)
	if err != nil {
		return DelayTarget{}, err
	}

	return target, nil
}