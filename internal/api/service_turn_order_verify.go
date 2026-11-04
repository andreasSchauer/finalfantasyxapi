package api

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

func verifyTurnOrderParams(v paramsVerifyer, params TurnOrderParams) (TurnOrderParams, error) {
	var err error

	params.TurnsAmt, err = v.int(params.TurnsAmt, ts.PfnTurnsAmt)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = v.field(params.IgnFirstTurn, ts.PfnIgnFirstTurn)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnImmunities, err = v.field(params.IgnImmunities, ts.PfnIgnImmunities)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.RNG, err = v.enum(params.RNG, ts.PfnRNG)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.BattleStart, err = v.enum(params.BattleStart, ts.PfnBattleStart)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.FormationID, err = v.idPtr(params.FormationID, ts.PfnFormationID)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Party, err = v.array(params.Party, ts.PfnParty, vfTurnOrderParty)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Mons, err = v.array(params.Mons, ts.PfnMons, vfTurnOrderMon)
	if err != nil {
		return TurnOrderParams{}, err
	}

	return params, nil
}

func vfTurnOrderParty(v paramsVerifyer, item turnOrderParty, _ ts.FieldName) (turnOrderParty, error) {
	var err error

	item.UnitID, err = v.id(item.UnitID, ts.PfnUnitID)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Agility, err = v.int(item.Agility, ts.PfnAgility)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.FirstStrike, err = v.field(item.FirstStrike, ts.PfnFirstStrike)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Status, err = v.enumPtr(item.Status, ts.PfnStatus)
	if err != nil {
		return turnOrderParty{}, err
	}

	return item, nil
}

func vfTurnOrderMon(v paramsVerifyer, item turnOrderMon, _ ts.FieldName) (turnOrderMon, error) {
	var err error

	item.MonsterID, err = v.idPtr(item.MonsterID, ts.PfnUnitID)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.AltState, err = v.intPtr(item.AltState, ts.PfnAltState)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Name, err = v.fieldPtr(item.Name, ts.PfnName)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Agility, err = v.intPtr(item.Agility, ts.PfnAgility)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.FirstStrike, err = v.fieldPtr(item.FirstStrike, ts.PfnFirstStrike)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Status, err = v.enumPtr(item.Status, ts.PfnStatus)
	if err != nil {
		return turnOrderMon{}, err
	}

	return item, nil
}
