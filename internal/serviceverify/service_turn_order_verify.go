package serviceverify

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func VerifyTurnOrderParams(v ParamsVerifyer, params servicetypes.TurnOrderParams) (servicetypes.TurnOrderParams, error) {
	var err error

	params.TurnsAmt, err = v.Int(params.TurnsAmt, ts.PfnTurnsAmt)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = v.Bool(params.IgnFirstTurn, ts.PfnIgnFirstTurn)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.IgnImmunities, err = v.Bool(params.IgnImmunities, ts.PfnIgnImmunities)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.RNG, err = v.Enum(params.RNG, ts.PfnRNG)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.BattleStart, err = v.Enum(params.BattleStart, ts.PfnBattleStart)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.FormationID, err = v.IDPtr(params.FormationID, ts.PfnFormationID)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.Party, err = v.Array(params.Party, ts.PfnParty, vfTurnOrderParty)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.Mons, err = v.Array(params.Mons, ts.PfnMons, vfTurnOrderMon)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	return params, nil
}

func vfTurnOrderParty(v ParamsVerifyer, item servicetypes.TurnOrderParty, _ ts.FieldName) (servicetypes.TurnOrderParty, error) {
	var err error

	item.UnitID, err = v.ID(item.UnitID, ts.PfnUnitID)
	if err != nil {
		return servicetypes.TurnOrderParty{}, err
	}

	item.Agility, err = v.Int(item.Agility, ts.PfnAgility)
	if err != nil {
		return servicetypes.TurnOrderParty{}, err
	}

	item.FirstStrike, err = v.Bool(item.FirstStrike, ts.PfnFirstStrike)
	if err != nil {
		return servicetypes.TurnOrderParty{}, err
	}

	item.Status, err = v.EnumPtr(item.Status, ts.PfnStatus)
	if err != nil {
		return servicetypes.TurnOrderParty{}, err
	}

	return item, nil
}

func vfTurnOrderMon(v ParamsVerifyer, item servicetypes.TurnOrderMon, _ ts.FieldName) (servicetypes.TurnOrderMon, error) {
	var err error

	item.MonsterID, err = v.IDPtr(item.MonsterID, ts.PfnUnitID)
	if err != nil {
		return servicetypes.TurnOrderMon{}, err
	}

	item.AltState, err = v.IntPtr(item.AltState, ts.PfnAltState)
	if err != nil {
		return servicetypes.TurnOrderMon{}, err
	}

	item.Name, err = v.StringPtr(item.Name, ts.PfnName)
	if err != nil {
		return servicetypes.TurnOrderMon{}, err
	}

	item.Agility, err = v.IntPtr(item.Agility, ts.PfnAgility)
	if err != nil {
		return servicetypes.TurnOrderMon{}, err
	}

	item.FirstStrike, err = v.BoolPtr(item.FirstStrike, ts.PfnFirstStrike)
	if err != nil {
		return servicetypes.TurnOrderMon{}, err
	}

	item.Status, err = v.EnumPtr(item.Status, ts.PfnStatus)
	if err != nil {
		return servicetypes.TurnOrderMon{}, err
	}

	return item, nil
}
