package api

func verifyTurnOrderParams(cfg *Config, params TurnOrderParams, docs ParamsDoc, userValues map[FieldName]any) (TurnOrderParams, error) {
	paramsRules := compileParamsRules(docs.Fields)

	err := vfExistingFields(userValues, paramsRules)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.TurnsAmt, err = verifyParamField(cfg, params.TurnsAmt, pfnTurnsAmt, userValues, paramsRules, vfInt)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = verifyParamField(cfg, params.IgnFirstTurn, pfnIgnFirstTurn, userValues, paramsRules, nil)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnImmunities, err = verifyParamField(cfg, params.IgnImmunities, pfnIgnImmunities, userValues, paramsRules, nil)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.RNG, err = verifyParamField(cfg, params.RNG, pfnRNG, userValues, paramsRules, vfEnum)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.BattleStart, err = verifyParamField(cfg, params.BattleStart, pfnBattleStart, userValues, paramsRules, vfEnum)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.FormationID, err = verifyParamFieldPtr(cfg, params.FormationID, pfnFormationID, userValues, paramsRules, vfId)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Party, err = verifyParamFieldArr(cfg, params.Party, pfnParty, userValues, paramsRules, vfTurnOrderParty)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Mons, err = verifyParamFieldArr(cfg, params.Mons, pfnMons, userValues, paramsRules, vfTurnOrderMon)
	if err != nil {
		return TurnOrderParams{}, err
	}

	return params, nil
}

func vfTurnOrderParty(cfg *Config, item turnOrderParty, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (turnOrderParty, error) {
	var err error

	err = vfExistingFields(userValues, paramsRules)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.UnitID, err = verifyParamField(cfg, item.UnitID, pfnUnitID, userValues, paramsRules, vfId)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Agility, err = verifyParamField(cfg, item.Agility, pfnAgility, userValues, paramsRules, vfInt)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.FirstStrike, err = verifyParamField(cfg, item.FirstStrike, pfnFirstStrike, userValues, paramsRules, nil)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Status, err = verifyParamFieldPtr(cfg, item.Status, pfnStatus, userValues, paramsRules, vfEnum)
	if err != nil {
		return turnOrderParty{}, err
	}

	return item, nil
}

func vfTurnOrderMon(cfg *Config, item turnOrderMon, _ FieldName, userValues map[FieldName]any, paramsRules ParamsRules) (turnOrderMon, error) {
	var err error

	err = vfExistingFields(userValues, paramsRules)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.MonsterID, err = verifyParamFieldPtr(cfg, item.MonsterID, pfnUnitID, userValues, paramsRules, vfId)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.AltState, err = verifyParamFieldPtr(cfg, item.AltState, pfnAltState, userValues, paramsRules, vfInt)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Name, err = verifyParamFieldPtr(cfg, item.Name, pfnName, userValues, paramsRules, nil)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Agility, err = verifyParamFieldPtr(cfg, item.Agility, pfnAgility, userValues, paramsRules, vfInt)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.FirstStrike, err = verifyParamFieldPtr(cfg, item.FirstStrike, pfnFirstStrike, userValues, paramsRules, nil)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Status, err = verifyParamFieldPtr(cfg, item.Status, pfnStatus, userValues, paramsRules, vfEnum)
	if err != nil {
		return turnOrderMon{}, err
	}

	return item, nil
}
