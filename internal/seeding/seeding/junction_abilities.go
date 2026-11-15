package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getAbilities() []seedtypes.Ability {
	abilities := []seedtypes.Ability{}

	for _, ability := range seedhelpers.RawJSON.PlayerAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, ability := range seedhelpers.RawJSON.OverdriveAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, item := range seedhelpers.RawJSON.Items {
		if len(item.BattleInteractions) == 0 {
			continue
		}
		abilities = append(abilities, item.Ability)
	}

	for _, command := range seedhelpers.RawJSON.TriggerCommands {
		abilities = append(abilities, command.Ability)
	}

	for _, ability := range seedhelpers.RawJSON.MiscAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, ability := range seedhelpers.RawJSON.EnemyAbilities {
		abilities = append(abilities, ability.Ability)
	}

	return abilities
}

func getAbilityBattleInteractions(a seedtypes.Ability) ([]seedtypes.BattleInteraction, error) {
	return a.BattleInteractions, nil
}

func SeedJuncAbilitiesBattleInteractions(qtx *database.Queries, ctx context.Context) error {
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
