package api

import (
	"net/http"
	"slices"
)

func verifyDropChanceParams(cfg *Config, params DropChanceParams, valueMap map[FieldName]any) (DropChanceParams, error) {
	valTree := compileValidationTree(cfg.getDropChanceParamsDoc().Fields)

	err := vfExistingFields(valueMap, valTree)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.Monster, err = verifyParamField(cfg, params.Monster, pfnMonsterID, valueMap, valTree, vfId)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.CharacterID, err = verifyParamFieldPtr(cfg, params.CharacterID, pfnCharacterID, valueMap, valTree, vfId)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.PartyMembers, err = verifyParamFieldIdArr(cfg, params.PartyMembers, pfnPartyMembers, valueMap, valTree)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.AutoAbilities, err = verifyParamFieldIdArr(cfg, params.AutoAbilities, pfnAutoAbilities, valueMap, valTree)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.LenientAbilities, err = verifyParamField(cfg, params.LenientAbilities, pfnLenientAbilities, valueMap, valTree, nil)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.EquipType, err = verifyParamFieldPtr(cfg, params.EquipType, pfnEquipType, valueMap, valTree, vfEnum)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.MinEmptySlots, err = verifyParamFieldPtr(cfg, params.MinEmptySlots, pfnMinEmptySlots, valueMap, valTree, vfInt)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.TotalSlots, err = verifyParamFieldPtr(cfg, params.TotalSlots, pfnTotalSlots, valueMap, valTree, vfInt)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.Decimals, err = verifyParamField(cfg, params.Decimals, pfnDecimals, valueMap, valTree, nil)
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
	requiredSlots := getRequiredSlots(params)
	var totalSlots int32 = 4

	if params.TotalSlots != nil {
		totalSlots = *params.TotalSlots
	}

	if requiredSlots > totalSlots {
		if params.TotalSlots == nil {
			return newHTTPError(http.StatusBadRequest, "the amount of auto-abilities and empty slots combined can't exceed 4.", nil)
		}

		return newHTTPError(http.StatusBadRequest, "the amount of auto-abilities and empty slots combined can't exceed the total amount of slots.", nil)
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
