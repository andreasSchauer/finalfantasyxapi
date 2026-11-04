package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

func verifyTurnOrderParams(v paramsVerifyer, params TurnOrderParams) (TurnOrderParams, error) {
	var err error

	params.TurnsAmt, err = v.int(params.TurnsAmt, typestrings.PfnTurnsAmt)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = v.field(params.IgnFirstTurn, typestrings.PfnIgnFirstTurn)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnImmunities, err = v.field(params.IgnImmunities, typestrings.PfnIgnImmunities)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.RNG, err = v.enum(params.RNG, typestrings.PfnRNG)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.BattleStart, err = v.enum(params.BattleStart, typestrings.PfnBattleStart)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.FormationID, err = v.idPtr(params.FormationID, typestrings.PfnFormationID)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Party, err = v.array(params.Party, typestrings.PfnParty, vfTurnOrderParty)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Mons, err = v.array(params.Mons, typestrings.PfnMons, vfTurnOrderMon)
	if err != nil {
		return TurnOrderParams{}, err
	}

	return params, nil
}

func vfTurnOrderParty(v paramsVerifyer, item turnOrderParty, _ typestrings.FieldName) (turnOrderParty, error) {
	var err error

	item.UnitID, err = v.id(item.UnitID, typestrings.PfnUnitID)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Agility, err = v.int(item.Agility, typestrings.PfnAgility)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.FirstStrike, err = v.field(item.FirstStrike, typestrings.PfnFirstStrike)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Status, err = v.enumPtr(item.Status, typestrings.PfnStatus)
	if err != nil {
		return turnOrderParty{}, err
	}

	return item, nil
}

func vfTurnOrderMon(v paramsVerifyer, item turnOrderMon, _ typestrings.FieldName) (turnOrderMon, error) {
	var err error

	item.MonsterID, err = v.idPtr(item.MonsterID, typestrings.PfnUnitID)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.AltState, err = v.intPtr(item.AltState, typestrings.PfnAltState)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Name, err = v.fieldPtr(item.Name, typestrings.PfnName)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Agility, err = v.intPtr(item.Agility, typestrings.PfnAgility)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.FirstStrike, err = v.fieldPtr(item.FirstStrike, typestrings.PfnFirstStrike)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Status, err = v.enumPtr(item.Status, typestrings.PfnStatus)
	if err != nil {
		return turnOrderMon{}, err
	}

	return item, nil
}
