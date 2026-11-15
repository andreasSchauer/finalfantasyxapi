package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
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
		params.DataHash[i] = sh.GenerateDataHash(c)
		params.StatID[i] = c.StatID
		params.CalculationType[i] = database.CalculationType(c.CalculationType)
		params.Value[i] = c.Value
	}

	dbRows, err := qtx.CreateStatChangeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create stat changes: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractStatChanges() ([]seedtypes.StatChange, error) {
	changes := []seedtypes.StatChange{}

	for i := range sh.RawJSON.PlayerAbilities {
		ability := &sh.RawJSON.PlayerAbilities[i]

		newChanges, err := prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.OverdriveAbilities {
		ability := &sh.RawJSON.OverdriveAbilities[i]

		newChanges, err := prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.Items {
		item := &sh.RawJSON.Items[i]

		newChanges, err := prepareAbilityStatChanges(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.TriggerCommands {
		command := &sh.RawJSON.TriggerCommands[i]

		newChanges, err := prepareAbilityStatChanges(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.MiscAbilities {
		ability := &sh.RawJSON.MiscAbilities[i]

		newChanges, err := prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.EnemyAbilities {
		ability := &sh.RawJSON.EnemyAbilities[i]

		newChanges, err := prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.AutoAbilities {
		autoAbility := &sh.RawJSON.AutoAbilities[i]

		newChanges, err := prepareStatChanges(autoAbility.StatChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.StatusConditions {
		status := &sh.RawJSON.StatusConditions[i]

		newChanges, err := prepareStatChanges(status.StatChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	return sh.DedupeRows(changes, sh.Hashes), nil
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

		change.StatID, err = sh.AssignFK(change.StatName, lookup.L.Stats)
		if err != nil {
			return nil, err
		}

		changesNew = append(changesNew, *change)
	}

	return changesNew, nil
}
