package api

func verifyTurnOrderParams(cfg *Config, params TurnOrderParams, valueMap map[FieldName]any) (TurnOrderParams, error) {
	valTree := compileValidationTree(cfg.getTurnOrderParamsDoc().Fields)
	
	err := vfExistingFields(valueMap, valTree)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.TurnsAmt, err = verifyParamField(cfg, params.TurnsAmt, pfnTurnsAmt, valueMap, valTree, vfIntId)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = verifyParamField(cfg, params.IgnFirstTurn, pfnIgnFirstTurn, valueMap, valTree, nil)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnImmunities, err = verifyParamField(cfg, params.IgnImmunities, pfnIgnImmunities, valueMap, valTree, nil)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.RNG, err = verifyParamField(cfg, params.RNG, pfnRNG, valueMap, valTree, vfEnum)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.BattleStart, err = verifyParamField(cfg, params.BattleStart, pfnBattleStart, valueMap, valTree, vfEnum)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Formation, err = verifyParamFieldPtr(cfg, params.Formation, pfnFormation, valueMap, valTree, vfIntId)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Party, err = verifyParamFieldArr(cfg, params.Party, pfnParty, valueMap, valTree, vfTurnOrderParty)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Mons, err = verifyParamFieldArr(cfg, params.Mons, pfnMons, valueMap, valTree, vfTurnOrderMon)
	if err != nil {
		return TurnOrderParams{}, err
	}

	return params, nil
}

func vfTurnOrderParty(cfg *Config, item turnOrderParty, _ FieldName, valueMap map[FieldName]any, valTree ValidationTree) (turnOrderParty, error) {
	var err error

	err = vfExistingFields(valueMap, valTree)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.ID, err = verifyParamField(cfg, item.ID, pfnID, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Agility, err = verifyParamField(cfg, item.Agility, pfnAgility, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.FS, err = verifyParamField(cfg, item.FS, pfnFirstStrike, valueMap, valTree, nil)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Status, err = verifyParamFieldPtr(cfg, item.Status, pfnStatus, valueMap, valTree, vfEnum)
	if err != nil {
		return turnOrderParty{}, err
	}

	return item, nil
}

func vfTurnOrderMon(cfg *Config, item turnOrderMon, _ FieldName, valueMap map[FieldName]any, valTree ValidationTree) (turnOrderMon, error) {
	var err error

	err = vfExistingFields(valueMap, valTree)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.ID, err = verifyParamFieldPtr(cfg, item.ID, pfnID, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.AltState, err = verifyParamFieldPtr(cfg, item.AltState, pfnAltState, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Name, err = verifyParamFieldPtr(cfg, item.Name, pfnName, valueMap, valTree, nil)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Agility, err = verifyParamFieldPtr(cfg, item.Agility, pfnAgility, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.FirstStrike, err = verifyParamFieldPtr(cfg, item.FirstStrike, pfnFirstStrike, valueMap, valTree, nil)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Status, err = verifyParamFieldPtr(cfg, item.Status, pfnStatus, valueMap, valTree, vfEnum)
	if err != nil {
		return turnOrderMon{}, err
	}

	return item, nil
}