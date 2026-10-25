package api

import (
	"fmt"
	"net/http"
	"slices"
)

func verifyDropChanceParams(cfg *Config, params DropChanceParams, paramsRules ParamsRules, userValues map[FieldName]any) (DropChanceParams, error) {
	var err error

	params.Monster, err = verifyParamField(cfg, params.Monster, pfnMonsterID, userValues, paramsRules, vfId)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.CharacterID, err = verifyParamFieldPtr(cfg, params.CharacterID, pfnCharacterID, userValues, paramsRules, vfId)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.PartyMembers, err = verifyParamFieldIdArr(cfg, params.PartyMembers, pfnPartyMembers, userValues, paramsRules)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.AutoAbilities, err = verifyParamFieldIdArr(cfg, params.AutoAbilities, pfnAutoAbilities, userValues, paramsRules)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.OptionalAbilities, err = verifyParamField(cfg, params.OptionalAbilities, pfnOptionalAbilities, userValues, paramsRules, nil)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.EquipType, err = verifyParamFieldPtr(cfg, params.EquipType, pfnEquipType, userValues, paramsRules, vfEnum)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.MinEmptySlots, err = verifyParamFieldPtr(cfg, params.MinEmptySlots, pfnMinEmptySlots, userValues, paramsRules, vfInt)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.TotalSlots, err = verifyParamFieldPtr(cfg, params.TotalSlots, pfnTotalSlots, userValues, paramsRules, vfInt)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.Decimals, err = verifyParamField(cfg, params.Decimals, pfnDecimals, userValues, paramsRules, nil)
	if err != nil {
		return DropChanceParams{}, err
	}

	err = vfRequiredSlotAmount(params)
	if err != nil {
		return DropChanceParams{}, err
	}

	err = vfCharInParty(params)
	if err != nil {
		return DropChanceParams{}, err
	}

	return params, nil
}

func vfRequiredSlotAmount(params DropChanceParams) error {
	requiredSlots := params.GetRequiredSlots()
	var totalSlots int32 = 4

	if params.TotalSlots != nil {
		totalSlots = *params.TotalSlots
	}

	if requiredSlots > totalSlots {
		if params.TotalSlots == nil {
			return newHTTPError(http.StatusBadRequest, "the amount of auto-abilities and empty slots combined can't exceed 4.", nil)
		}

		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("the amount of auto-abilities and empty slots combined (%d) can't exceed the given total amount of slots (%d).", requiredSlots, totalSlots), nil)
	}

	return nil
}

func vfCharInParty(params DropChanceParams) error {
	if params.CharacterID == nil {
		return nil
	}

	if !slices.Contains(params.PartyMembers, *params.CharacterID) {
		return newHTTPError(http.StatusBadRequest, "character must be in the party to be eligible for receiving equipment.", nil)
	}

	return nil
}
