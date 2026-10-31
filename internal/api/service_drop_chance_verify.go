package api

import (
	"fmt"
	"net/http"
	"slices"
)

func verifyDropChanceParams(v paramsVerifyer, params DropChanceParams) (DropChanceParams, error) {
	var err error

	params.Monster, err = v.vfIDField(params.Monster, PfnMonsterID)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.CharacterID, err = v.verifyParamFieldPtr(params.CharacterID, PfnCharacterID, vfID)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.PartyMembers, err = v.verifyParamFieldIDArr(params.PartyMembers, PfnPartyMembers)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.AutoAbilities, err = v.verifyParamFieldIDArr(params.AutoAbilities, PfnAutoAbilities)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.OptionalAbilities, err = v.vfField(params.OptionalAbilities, PfnOptionalAbilities)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.EquipType, err = v.verifyParamFieldPtr(params.EquipType, PfnEquipType, vfEnum)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.MinEmptySlots, err = v.verifyParamFieldPtr(params.MinEmptySlots, PfnMinEmptySlots, vfInt)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.TotalSlots, err = v.verifyParamFieldPtr(params.TotalSlots, PfnTotalSlots, vfInt)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.Decimals, err = v.vfField(params.Decimals, PfnDecimals)
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
