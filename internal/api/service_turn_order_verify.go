package api

func verifyTurnOrderParams(v paramsVerifyer, params TurnOrderParams) (TurnOrderParams, error) {
	var err error

	params.TurnsAmt, err = v.int(params.TurnsAmt, PfnTurnsAmt)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = v.field(params.IgnFirstTurn, PfnIgnFirstTurn)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnImmunities, err = v.field(params.IgnImmunities, PfnIgnImmunities)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.RNG, err = v.enum(params.RNG, PfnRNG)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.BattleStart, err = v.enum(params.BattleStart, PfnBattleStart)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.FormationID, err = v.idPtr(params.FormationID, PfnFormationID)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Party, err = v.array(params.Party, PfnParty, vfTurnOrderParty)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Mons, err = v.array(params.Mons, PfnMons, vfTurnOrderMon)
	if err != nil {
		return TurnOrderParams{}, err
	}

	return params, nil
}

func vfTurnOrderParty(v paramsVerifyer, item turnOrderParty, _ FieldName) (turnOrderParty, error) {
	var err error

	item.UnitID, err = v.id(item.UnitID, PfnUnitID)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Agility, err = v.int(item.Agility, PfnAgility)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.FirstStrike, err = v.field(item.FirstStrike, PfnFirstStrike)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Status, err = v.enumPtr(item.Status, PfnStatus)
	if err != nil {
		return turnOrderParty{}, err
	}

	return item, nil
}

func vfTurnOrderMon(v paramsVerifyer, item turnOrderMon, _ FieldName) (turnOrderMon, error) {
	var err error

	item.MonsterID, err = v.idPtr(item.MonsterID, PfnUnitID)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.AltState, err = v.intPtr(item.AltState, PfnAltState)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Name, err = v.fieldPtr(item.Name, PfnName)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Agility, err = v.intPtr(item.Agility, PfnAgility)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.FirstStrike, err = v.fieldPtr(item.FirstStrike, PfnFirstStrike)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Status, err = v.enumPtr(item.Status, PfnStatus)
	if err != nil {
		return turnOrderMon{}, err
	}

	return item, nil
}
