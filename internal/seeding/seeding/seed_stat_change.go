package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop5SeedStatChanges(qtx *database.Queries, ctx context.Context) error {
	changes, err := extractStatChanges()
	if err != nil {
		return err
	}

	params := database.CreateStatChangeBulkParams{
		DataHash:        make([]string, len(changes)),
		StatID:          make([]int32, len(changes)),
		CalculationType: make([]database.CalculationType, len(changes)),
		Value:           make([]float32, len(changes)),
	}

	for i, c := range changes {
		params.DataHash[i] = seedhelpers.GenerateDataHash(c)
		params.StatID[i] = c.StatID
		params.CalculationType[i] = database.CalculationType(c.CalculationType)
		params.Value[i] = c.Value
	}

	dbRows, err := qtx.CreateStatChangeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create stat changes: %v", err)
	}

	for _, row := range dbRows {
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractStatChanges() ([]seedtypes.StatChange, error) {
	changes := []seedtypes.StatChange{}

	for i := range seedhelpers.RawJSON.PlayerAbilities {
		ability := &seedhelpers.RawJSON.PlayerAbilities[i]

		newChanges, err := prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range seedhelpers.RawJSON.OverdriveAbilities {
		ability := &seedhelpers.RawJSON.OverdriveAbilities[i]

		newChanges, err := prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range seedhelpers.RawJSON.Items {
		item := &seedhelpers.RawJSON.Items[i]

		newChanges, err := prepareAbilityStatChanges(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range seedhelpers.RawJSON.TriggerCommands {
		command := &seedhelpers.RawJSON.TriggerCommands[i]

		newChanges, err := prepareAbilityStatChanges(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range seedhelpers.RawJSON.MiscAbilities {
		ability := &seedhelpers.RawJSON.MiscAbilities[i]

		newChanges, err := prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range seedhelpers.RawJSON.EnemyAbilities {
		ability := &seedhelpers.RawJSON.EnemyAbilities[i]

		newChanges, err := prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range seedhelpers.RawJSON.AutoAbilities {
		autoAbility := &seedhelpers.RawJSON.AutoAbilities[i]

		newChanges, err := prepareStatChanges(autoAbility.StatChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range seedhelpers.RawJSON.StatusConditions {
		status := &seedhelpers.RawJSON.StatusConditions[i]

		newChanges, err := prepareStatChanges(status.StatChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	return dedupeRows(changes, seedhelpers.Hashes), nil
}

func prepareAbilityStatChanges(battleInteractions []seedtypes.BattleInteraction) ([]seedtypes.StatChange, error) {
	changes := []seedtypes.StatChange{}

	for i := range battleInteractions {
		bi := &battleInteractions[i]

		changesNew, err := prepareStatChanges(bi.StatChanges)
		if err != nil {
			return nil, err
		}
		changes = append(changes, changesNew...)
	}

	return changes, nil
}

func prepareStatChanges(changes []seedtypes.StatChange) ([]seedtypes.StatChange, error) {
	changesNew := []seedtypes.StatChange{}
	var err error

	for i := range changes {
		change := &changes[i]

		change.StatID, err = seedhelpers.AssignFK(change.StatName, lookup.L.Stats)
		if err != nil {
			return nil, err
		}

		changesNew = append(changesNew, *change)
	}

	return changesNew, nil
}
