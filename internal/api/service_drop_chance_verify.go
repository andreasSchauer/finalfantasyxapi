package api

func verifyDropChanceParams(cfg *Config, params DropChanceParams, valueMap map[FieldName]any) (DropChanceParams, error) {
	valTree := compileValidationTree(cfg.getDropChanceParamsDoc().Fields)

	err := vfExistingFields(valueMap, valTree)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.Monster, err = verifyParamField(cfg, params.Monster, pfnMonster, valueMap, valTree, vfIntId)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.Character, err = verifyParamFieldPtr(cfg, params.Character, pfnCharacter, valueMap, valTree, vfIntId)
	if err != nil {
		return DropChanceParams{}, err
	}
	
	params.PartyMembers, err = verifyParamFieldIntIdArr(cfg, params.PartyMembers, pfnPartyMembers, valueMap, valTree)
	if err != nil {
		return DropChanceParams{}, err
	}
	
	params.AutoAbilities, err = verifyParamFieldIntIdArr(cfg, params.AutoAbilities, pfnAutoAbilities, valueMap, valTree)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.LenientAbilities, err = verifyParamField(cfg, params.LenientAbilities, pfnLenientAbilities, valueMap, valTree, nil)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.MinEmptySlots, err = verifyParamFieldPtr(cfg, params.MinEmptySlots, pfnMinEmptySlots, valueMap, valTree, vfIntId)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.TotalSlots, err = verifyParamFieldPtr(cfg, params.TotalSlots, pfnTotalSlots, valueMap, valTree, vfIntId)
	if err != nil {
		return DropChanceParams{}, err
	}

	return params, nil
}