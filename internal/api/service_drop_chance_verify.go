package api

import (
	"net/http"
)

func verifyDropChanceParams(cfg *Config, params DropChanceParams, valueMap map[FieldName]any) (DropChanceParams, error) {
	valTree := compileValidationTree(cfg.getDropChanceParamsDoc().Fields)

	err := vfExistingFields(valueMap, valTree)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.Monster, err = verifyParamField(cfg, params.Monster, pfnMonster, valueMap, valTree, vfIntId)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.Character, err = verifyParamFieldPtr(cfg, params.Character, pfnCharacter, valueMap, valTree, vfIntId)
	if err != nil {
		return DropChanceParams{}, err
	}
	
	params.PartyMembers, err = verifyParamFieldIntIdArr(cfg, params.PartyMembers, pfnPartyMembers, valueMap, valTree)
	if err != nil {
		return DropChanceParams{}, err
	}
	
	params.AutoAbilities, err = verifyParamFieldIntIdArr(cfg, params.AutoAbilities, pfnAutoAbilities, valueMap, valTree)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.LenientAbilities, err = verifyParamField(cfg, params.LenientAbilities, pfnLenientAbilities, valueMap, valTree, nil)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.MinEmptySlots, err = verifyParamFieldPtr(cfg, params.MinEmptySlots, pfnMinEmptySlots, valueMap, valTree, vfIntId)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.TotalSlots, err = verifyParamFieldPtr(cfg, params.TotalSlots, pfnTotalSlots, valueMap, valTree, vfIntId)
	if err != nil {
		return DropChanceParams{}, err
	}

	err = vfRequiredSlotAmount(params)
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

func getRequiredSlots(params DropChanceParams) int32 {
	var minEmptySlots int32 = 0

	if params.MinEmptySlots != nil {
		minEmptySlots = *params.MinEmptySlots
	}

	return int32(len(params.AutoAbilities)) + minEmptySlots
}