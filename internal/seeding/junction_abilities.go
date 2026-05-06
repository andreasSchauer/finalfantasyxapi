package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getAbilities() []Ability {
	abilities := []Ability{}

	for _, ability := range l.json.playerAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, ability := range l.json.overdriveAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, item := range l.json.items {
		if len(item.BattleInteractions) == 0 {
			continue
		}
		abilities = append(abilities, item.Ability)
	}

	for _, command := range l.json.triggerCommands {
		abilities = append(abilities, command.Ability)
	}

	for _, ability := range l.json.miscAbilities {
		abilities = append(abilities, ability.Ability)
	}

	for _, ability := range l.json.enemyAbilities {
		abilities = append(abilities, ability.Ability)
	}

	return abilities
}

func (l *Lookup) getAbilityBattleInteractions(a Ability) ([]BattleInteraction, error) {
	return a.BattleInteractions, nil
}

func (l *Lookup) seedJuncAbilitiesBattleInteractions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "abilities + battle interactions"
	jParams, err := processJunctions(l, desc, l.getAbilities(), l.getAbilityBattleInteractions)
	if err != nil {
		return err
	}

	return qtx.CreateAbilitiesBattleInteractionsJunctionBulk(ctx, database.CreateAbilitiesBattleInteractionsJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		AbilityID:           jParams.ParentIDs,
		BattleInteractionID: jParams.ChildIDs,
	})
}
