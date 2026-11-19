package serviceverify

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func TurnOrderParams(v ParamsVerifyer, params servicetypes.TurnOrderParams) (servicetypes.TurnOrderParams, error) {
	var err error

	params.TurnsAmt, err = v.int(params.TurnsAmt, ts.PfnTurnsAmt)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = v.bool(params.IgnFirstTurn, ts.PfnIgnFirstTurn)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.IgnImmunities, err = v.bool(params.IgnImmunities, ts.PfnIgnImmunities)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.RNG, err = v.enum(params.RNG, ts.PfnRNG)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.BattleStart, err = v.enum(params.BattleStart, ts.PfnBattleStart)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.FormationID, err = v.idPtr(params.FormationID, ts.PfnFormationID)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.Party, err = v.array(params.Party, ts.PfnParty, vfTurnOrderParty)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	params.Mons, err = v.array(params.Mons, ts.PfnMons, vfTurnOrderMon)
	if err != nil {
		return servicetypes.TurnOrderParams{}, err
	}

	return params, nil
}

func vfTurnOrderParty(v ParamsVerifyer, item servicetypes.TurnOrderParty, _ ts.FieldName) (servicetypes.TurnOrderParty, error) {
	var err error

	item.UnitID, err = v.id(item.UnitID, ts.PfnUnitID)
	if err != nil {
		return servicetypes.TurnOrderParty{}, err
	}

	item.Agility, err = v.int(item.Agility, ts.PfnAgility)
	if err != nil {
		return servicetypes.TurnOrderParty{}, err
	}

	item.FirstStrike, err = v.bool(item.FirstStrike, ts.PfnFirstStrike)
	if err != nil {
		return servicetypes.TurnOrderParty{}, err
	}

	item.Status, err = v.enumPtr(item.Status, ts.PfnStatus)
	if err != nil {
		return servicetypes.TurnOrderParty{}, err
	}

	return item, nil
}

func vfTurnOrderMon(v ParamsVerifyer, item servicetypes.TurnOrderMon, _ ts.FieldName) (servicetypes.TurnOrderMon, error) {
	var err error

	item.Lookup, err = v.objPtr(item.Lookup, ts.PfnLookup, vfTurnOrderMonLookup)
	if err != nil {
		return servicetypes.TurnOrderMon{}, err
	}

	item.Custom, err = v.objPtr(item.Custom, ts.PfnCustom, verifyFn[servicetypes.TurnOrderMonCustom](vfTurnOrderMonCustom))
	if err != nil {
		return servicetypes.TurnOrderMon{}, err
	}

	return item, nil
}

func vfTurnOrderMonLookup(v ParamsVerifyer, item servicetypes.TurnOrderMonLookup, _ ts.FieldName) (servicetypes.TurnOrderMonLookup, error) {
	var err error

	item.MonsterID, err = v.id(item.MonsterID, ts.PfnMonsterID)
	if err != nil {
		return servicetypes.TurnOrderMonLookup{}, err
	}

	item.AltState, err = v.intPtr(item.AltState, ts.PfnAltState)
	if err != nil {
		return servicetypes.TurnOrderMonLookup{}, err
	}

	item.Overrides, err = v.objPtr(item.Overrides, ts.PfnOverrides, vfTurnOrderMonOverrides)
	if err != nil {
		return servicetypes.TurnOrderMonLookup{}, err
	}

	return item, nil
}

func vfTurnOrderMonOverrides(v ParamsVerifyer, item servicetypes.TurnOrderMonOverrides, _ ts.FieldName) (servicetypes.TurnOrderMonOverrides, error) {
	var err error

	item.Agility, err = v.intPtr(item.Agility, ts.PfnAgility)
	if err != nil {
		return servicetypes.TurnOrderMonOverrides{}, err
	}

	item.FirstStrike, err = v.boolPtr(item.FirstStrike, ts.PfnFirstStrike)
	if err != nil {
		return servicetypes.TurnOrderMonOverrides{}, err
	}

	item.Status, err = v.enumPtr(item.Status, ts.PfnStatus)
	if err != nil {
		return servicetypes.TurnOrderMonOverrides{}, err
	}

	return item, nil
}

func vfTurnOrderMonCustom(v ParamsVerifyer, item servicetypes.TurnOrderMonCustom, _ ts.FieldName) (servicetypes.TurnOrderMonCustom, error) {
	var err error

	item.Name, err = v.string(item.Name, ts.PfnName)
	if err != nil {
		return servicetypes.TurnOrderMonCustom{}, err
	}

	item.Agility, err = v.int(item.Agility, ts.PfnAgility)
	if err != nil {
		return servicetypes.TurnOrderMonCustom{}, err
	}

	item.FirstStrike, err = v.bool(item.FirstStrike, ts.PfnFirstStrike)
	if err != nil {
		return servicetypes.TurnOrderMonCustom{}, err
	}

	item.Status, err = v.enumPtr(item.Status, ts.PfnStatus)
	if err != nil {
		return servicetypes.TurnOrderMonCustom{}, err
	}

	return item, nil
}
