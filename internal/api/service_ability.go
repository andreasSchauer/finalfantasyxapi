package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func quickAssembleAbility(cfg *Config, abilityID int32, interactionPtr *int32) (seeding.Ability, seeding.BattleInteraction, error) {
	ability, _ := seeding.GetResourceByID(abilityID, cfg.L.AbilitiesID)
	abilityType := string(ability.Type)
	abilityString := h.NameToString(ability.Name, ability.Version, &abilityType)
	interactionsAmt := h.Len32(ability.BattleInteractions)

	if interactionPtr == nil {
		if interactionsAmt > 1 {
			return seeding.Ability{}, seeding.BattleInteraction{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("ability '%s' has multiple battle interactions. use '%s.%s.%s' to specify the interaction.", abilityString, PfnAbility, PfnLookup, PfnInteraction), nil)
		}
		return ability, ability.BattleInteractions[0], nil
	}

	interactionID := *interactionPtr

	if interactionID > interactionsAmt || interactionID <= 0 {
		return seeding.Ability{}, seeding.BattleInteraction{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided interaction id '%d' used for ability '%s' is out of range. max id: %d.", interactionID, abilityString, interactionsAmt), nil)
	}

	wantedBI := ability.BattleInteractions[interactionID - 1]

	return ability, wantedBI, nil
}
