package api

func verifyTurnOrderParams(v paramsVerifyer, params TurnOrderParams) (TurnOrderParams, error) {
	var err error

	params.TurnsAmt, err = v.vfIntField(params.TurnsAmt, PfnTurnsAmt)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = v.vfField(params.IgnFirstTurn, PfnIgnFirstTurn)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnImmunities, err = v.vfField(params.IgnImmunities, PfnIgnImmunities)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.RNG, err = v.vfEnumField(params.RNG, PfnRNG)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.BattleStart, err = v.vfEnumField(params.BattleStart, PfnBattleStart)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.FormationID, err = v.verifyParamFieldPtr(params.FormationID, PfnFormationID, vfID)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Party, err = v.verifyParamFieldArr(params.Party, PfnParty, vfTurnOrderParty)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Mons, err = v.verifyParamFieldArr(params.Mons, PfnMons, vfTurnOrderMon)
	if err != nil {
		return TurnOrderParams{}, err
	}

	return params, nil
}

func vfTurnOrderParty(v paramsVerifyer, item turnOrderParty, _ FieldName) (turnOrderParty, error) {
	var err error

	err = v.vfExistingFields()
	if err != nil {
		return turnOrderParty{}, err
	}

	item.UnitID, err = v.vfIDField(item.UnitID, PfnUnitID)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Agility, err = v.vfIntField(item.Agility, PfnAgility)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.FirstStrike, err = v.vfField(item.FirstStrike, PfnFirstStrike)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Status, err = v.verifyParamFieldPtr(item.Status, PfnStatus, vfEnum)
	if err != nil {
		return turnOrderParty{}, err
	}

	return item, nil
}

func vfTurnOrderMon(v paramsVerifyer, item turnOrderMon, _ FieldName) (turnOrderMon, error) {
	var err error

	err = v.vfExistingFields()
	if err != nil {
		return turnOrderMon{}, err
	}

	item.MonsterID, err = v.verifyParamFieldPtr(item.MonsterID, PfnUnitID, vfID)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.AltState, err = v.verifyParamFieldPtr(item.AltState, PfnAltState, vfInt)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Name, err = v.verifyParamFieldPtr(item.Name, PfnName, nil)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Agility, err = v.verifyParamFieldPtr(item.Agility, PfnAgility, vfInt)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.FirstStrike, err = v.verifyParamFieldPtr(item.FirstStrike, PfnFirstStrike, nil)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Status, err = v.verifyParamFieldPtr(item.Status, PfnStatus, vfEnum)
	if err != nil {
		return turnOrderMon{}, err
	}

	return item, nil
}
