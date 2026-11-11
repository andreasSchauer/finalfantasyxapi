package serviceverify

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func DropChanceParams(v ParamsVerifyer, params servicetypes.DropChanceParams) (servicetypes.DropChanceParams, error) {
	var err error

	params.Monster, err = v.id(params.Monster, ts.PfnMonsterID)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	params.CharacterID, err = v.idPtr(params.CharacterID, ts.PfnCharacterID)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	params.PartyMembers, err = v.idArray(params.PartyMembers, ts.PfnPartyMembers)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	params.AutoAbilities, err = v.idArray(params.AutoAbilities, ts.PfnAutoAbilities)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	params.OptionalAbilities, err = v.bool(params.OptionalAbilities, ts.PfnOptionalAbilities)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	params.EquipType, err = v.enumPtr(params.EquipType, ts.PfnEquipType)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	params.MinEmptySlots, err = v.intPtr(params.MinEmptySlots, ts.PfnMinEmptySlots)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	params.TotalSlots, err = v.intPtr(params.TotalSlots, ts.PfnTotalSlots)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	params.Decimals, err = v.bool(params.Decimals, ts.PfnDecimals)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	err = vfRequiredSlotAmount(params)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	err = vfCharInParty(params)
	if err != nil {
		return servicetypes.DropChanceParams{}, err
	}

	return params, nil
}

func vfRequiredSlotAmount(params servicetypes.DropChanceParams) error {
	requiredSlots := params.GetRequiredSlots()
	var totalSlots int32 = 4

	if params.TotalSlots != nil {
		totalSlots = *params.TotalSlots
	}

	if requiredSlots > totalSlots {
		if params.TotalSlots == nil {
			return errs.NewHTTPError(http.StatusBadRequest, "the amount of auto-abilities and empty slots combined can't exceed 4", nil)
		}

		return errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("the amount of auto-abilities and empty slots combined (%d) can't exceed the given total amount of slots (%d)", requiredSlots, totalSlots), nil)
	}

	return nil
}

func vfCharInParty(params servicetypes.DropChanceParams) error {
	if params.CharacterID == nil {
		return nil
	}

	if !slices.Contains(params.PartyMembers, *params.CharacterID) {
		return errs.NewHTTPError(http.StatusBadRequest, "character must be in the party to be eligible for receiving equipment", nil)
	}

	return nil
}
