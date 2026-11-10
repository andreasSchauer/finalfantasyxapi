package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getAbilities() []seedtypes.Ability {
	abilities := []seedtypes.Ability{}

	for _, ability := range rawJSON.playerAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, ability := range rawJSON.overdriveAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, item := range rawJSON.items {
		if len(item.BattleInteractions) == 0 {
			continue
		}
		abilities = append(abilities, item.Ability)
	}

	for _, command := range rawJSON.triggerCommands {
		abilities = append(abilities, command.Ability)
	}

	for _, ability := range rawJSON.miscAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, ability := range rawJSON.enemyAbilities {
		abilities = append(abilities, ability.Ability)
	}

	return abilities
}

func getAbilityBattleInteractions(a seedtypes.Ability) ([]seedtypes.BattleInteraction, error) {
	return a.BattleInteractions, nil
}

func seedJuncAbilitiesBattleInteractions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "abilities + battle interactions"
	jParams, err := processJunctions(desc, getAbilities(), getAbilityBattleInteractions)
	if err != nil {
		return err
	}

	return qtx.CreateAbilitiesBattleInteractionsJunctionBulk(ctx, database.CreateAbilitiesBattleInteractionsJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.ParentIDs,
		BattleInteractionID: jParams.ChildIDs,
	})
}
