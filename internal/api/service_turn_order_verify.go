package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func verifyTurnOrderParams(cfg *Config, params TurnOrderParams, valueMap map[FieldName]any) (TurnOrderParams, error) {
	var err error
	valTree := compileValidationTree(cfg.getTurnOrderParamsDoc().Fields)

	params.TurnsAmt, err = verifyParamField(cfg, params.TurnsAmt, pfnTurnsAmt, valueMap, valTree, vfIntId)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.IgnFirstTurn, err = verifyParamField(cfg, params.IgnFirstTurn, pfnIgnFirstTurn, valueMap, valTree, nil)
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

	params.AglK, err = verifyParamField(cfg, params.AglK, pfnAglK, valueMap, valTree, vfIntId)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.AglY, err = verifyParamField(cfg, params.AglY, pfnAglY, valueMap, valTree, vfIntId)
	if err != nil {
		return TurnOrderParams{}, err
	}

	params.Battles, err = verifyParamField(cfg, params.Battles, pfnBattles, valueMap, valTree, vfIntId)
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

	params.MonsCustom, err = verifyParamFieldArr(cfg, params.MonsCustom, pfnMonsCustom, valueMap, valTree, vfTurnOrderMonCustom)
	if err != nil {
		return TurnOrderParams{}, err
	}

	return params, nil
}

func vfTurnOrderParty(cfg *Config, item turnOrderParty, _ FieldName, valueMap map[FieldName]any, valTree ValidationTree) (turnOrderParty, error) {
	var err error
	
	item.Name, err = verifyParamField(cfg, item.Name, pfnName, valueMap, valTree, nil)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Agl, err = verifyParamField(cfg, item.Agl, pfnAgl, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.FS, err = verifyParamField(cfg, item.FS, pfnFS, valueMap, valTree, nil)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Status, err = verifyParamFieldPtr(cfg, item.Status, pfnStatus, valueMap, valTree, vfEnum)
	if err != nil {
		return turnOrderParty{}, err
	}

	item.Offset, err = verifyParamField(cfg, item.Offset, pfnOffset, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderParty{}, err
	}

	return item, nil
}

func vfTurnOrderMon(cfg *Config, item turnOrderMon, _ FieldName, valueMap map[FieldName]any, valTree ValidationTree) (turnOrderMon, error) {
	var err error
	
	item.ID, err = verifyParamField(cfg, item.ID, pfnID, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.AltState, err = verifyParamFieldPtr(cfg, item.AltState, pfnAltState, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderMon{}, err
	}

	if item.AltState != nil {
		mon, _ := seeding.GetResourceByID(item.ID, cfg.l.MonstersID)
		altStateAmt := int32(len(mon.AlteredStates))

		if altStateAmt == 0 {
			return turnOrderMon{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("%s (id: %d) has no altered states.", mon.Error(), item.ID), nil)
		}

		if *item.AltState > altStateAmt {
			return turnOrderMon{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value for field '%s'. %s (id: %d) only has %d altered states.", pfnAltState, mon.Error(), item.ID, altStateAmt), nil)
		}
	}

	item.Status, err = verifyParamFieldPtr(cfg, item.Status, pfnStatus, valueMap, valTree, vfEnum)
	if err != nil {
		return turnOrderMon{}, err
	}

	item.Offset, err = verifyParamField(cfg, item.Offset, pfnOffset, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderMon{}, err
	}

	return item, nil
}

func vfTurnOrderMonCustom(cfg *Config, item turnOrderMonCustom, _ FieldName, valueMap map[FieldName]any, valTree ValidationTree) (turnOrderMonCustom, error) {
	var err error
	
	item.Name, err = verifyParamField(cfg, item.Name, pfnName, valueMap, valTree, nil)
	if err != nil {
		return turnOrderMonCustom{}, err
	}

	item.Agl, err = verifyParamField(cfg, item.Agl, pfnAgl, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderMonCustom{}, err
	}

	item.FS, err = verifyParamField(cfg, item.FS, pfnFS, valueMap, valTree, nil)
	if err != nil {
		return turnOrderMonCustom{}, err
	}

	item.Status, err = verifyParamFieldPtr(cfg, item.Status, pfnStatus, valueMap, valTree, vfEnum)
	if err != nil {
		return turnOrderMonCustom{}, err
	}

	item.Offset, err = verifyParamField(cfg, item.Offset, pfnOffset, valueMap, valTree, vfIntId)
	if err != nil {
		return turnOrderMonCustom{}, err
	}

	return item, nil
}