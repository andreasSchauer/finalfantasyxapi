package api

import (
	"fmt"
	"net/http"
	"slices"
)

func verifyDropChanceParams(v paramsVerifyer, params DropChanceParams) (DropChanceParams, error) {
	var err error

	params.Monster, err = v.id(params.Monster, PfnMonsterID)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.CharacterID, err = v.idPtr(params.CharacterID, PfnCharacterID)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.PartyMembers, err = v.idArray(params.PartyMembers, PfnPartyMembers)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.AutoAbilities, err = v.idArray(params.AutoAbilities, PfnAutoAbilities)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.OptionalAbilities, err = v.field(params.OptionalAbilities, PfnOptionalAbilities)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.EquipType, err = v.enumPtr(params.EquipType, PfnEquipType)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.MinEmptySlots, err = v.intPtr(params.MinEmptySlots, PfnMinEmptySlots)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.TotalSlots, err = v.intPtr(params.TotalSlots, PfnTotalSlots)
	if err != nil {
		return DropChanceParams{}, err
	}

	params.Decimals, err = v.field(params.Decimals, PfnDecimals)
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
