// Package servicehelpers asjdasd
package servicehelpers

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func QuickAssembleAbility(abilityID int32, interactionPtr *int32) (seedtypes.Ability, seedtypes.BattleInteraction, error) {
	ability, _ := lookup.GetResourceByID(abilityID, lookup.L.AbilitiesID)
	abilityType := string(ability.Type)
	abilityString := h.NameToString(ability.Name, ability.Version, &abilityType)
	interactionsAmt := h.Len32(ability.BattleInteractions)

	if interactionPtr == nil {
		if interactionsAmt > 1 {
			return seedtypes.Ability{}, seedtypes.BattleInteraction{}, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ability '%s' has multiple battle interactions. use '%s.%s.%s' to specify the interaction", abilityString, ts.PfnAbility, ts.PfnLookup, ts.PfnInteraction), nil)
		}
		return ability, ability.BattleInteractions[0], nil
	}

	interactionID := *interactionPtr

	if interactionID > interactionsAmt || interactionID <= 0 {
		return seedtypes.Ability{}, seedtypes.BattleInteraction{}, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("provided interaction id '%d' used for ability '%s' is out of range. max id: %d", interactionID, abilityString, interactionsAmt), nil)
	}

	wantedBI := ability.BattleInteractions[interactionID-1]

	return ability, wantedBI, nil
}
