package api

func verifyTurnOrderParams(cfg *Config, params TurnOrderParams, paramsRules ParamsRules, userValues map[FieldName]any) (TurnOrderParams, error) {
	var err error

	params.TurnsAmt, err = verifyParamField(cfg, params.TurnsAmt, PfnTurnsAmt, userValues, paramsRules, vfInt)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = verifyParamField(cfg, params.IgnFirstTurn, PfnIgnFirstTurn, userValues, paramsRules, nil)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnImmunities, err = verifyParamField(cfg, params.IgnImmunities, PfnIgnImmunities, userValues, paramsRules, nil)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.RNG, err = verifyParamField(cfg, params.RNG, PfnRNG, userValues, paramsRules, vfEnum)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.BattleStart, err = verifyParamField(cfg, params.BattleStart, PfnBattleStart, userValues, paramsRules, vfEnum)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.FormationID, err = verifyParamFieldPtr(cfg, params.FormationID, PfnFormationID, userValues, paramsRules, vfId)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Party, err = verifyParamFieldArr(cfg, params.Party, PfnParty, userValues, paramsRules, vfTurnOrderParty)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Mons, err = verifyParamFieldArr(cfg, params.Mons, PfnMons, userValues, paramsRules, vfTurnOrderMon)
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

	item.UnitID, err = verifyParamField(cfg, item.UnitID, PfnUnitID, userValues, paramsRules, vfId)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Agility, err = verifyParamField(cfg, item.Agility, PfnAgility, userValues, paramsRules, vfInt)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.FirstStrike, err = verifyParamField(cfg, item.FirstStrike, PfnFirstStrike, userValues, paramsRules, nil)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Status, err = verifyParamFieldPtr(cfg, item.Status, PfnStatus, userValues, paramsRules, vfEnum)
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

	item.MonsterID, err = verifyParamFieldPtr(cfg, item.MonsterID, PfnUnitID, userValues, paramsRules, vfId)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.AltState, err = verifyParamFieldPtr(cfg, item.AltState, PfnAltState, userValues, paramsRules, vfInt)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Name, err = verifyParamFieldPtr(cfg, item.Name, PfnName, userValues, paramsRules, nil)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Agility, err = verifyParamFieldPtr(cfg, item.Agility, PfnAgility, userValues, paramsRules, vfInt)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.FirstStrike, err = verifyParamFieldPtr(cfg, item.FirstStrike, PfnFirstStrike, userValues, paramsRules, nil)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Status, err = verifyParamFieldPtr(cfg, item.Status, PfnStatus, userValues, paramsRules, vfEnum)
	if err != nil {
		return turnOrderMon{}, err
	}

	return item, nil
}
