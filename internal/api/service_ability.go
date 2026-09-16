package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


// can reuse for other endpoints
func getBattleInteraction(cfg *Config, abilityID int32, interactionPtr *int32) (seeding.BattleInteraction, error) {
	ability, _ := seeding.GetResourceByID(abilityID, cfg.l.AbilitiesID)
	interactionsAmt := h.Len32(ability.BattleInteractions)
	abilityType := string(ability.Type)
	abilityString := h.NameToString(ability.Name, ability.Version, &abilityType)

	if interactionPtr == nil && interactionsAmt > 1 {
		return seeding.BattleInteraction{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("no interaction id provided, when ability '%s' has more battle interactions than one.", abilityString), nil)
	}

	if interactionPtr == nil {
		return ability.BattleInteractions[0], nil
	}

	interactionID := *interactionPtr

	if interactionID > interactionsAmt || interactionID <= 0 {
		return seeding.BattleInteraction{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided interaction id '%d' used for ability '%s' is out of range. max id: %d.", interactionID, abilityString, interactionsAmt), nil)
	}

	interactionIdx := interactionID-1

	return ability.BattleInteractions[interactionIdx], nil
}