package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop2ModifierChanges(qtx *database.Queries, ctx context.Context) error {
	changes, err := extractModifierChanges()
	if err != nil {
		return err
	}

	params := database.CreateModifierChangeBulkParams{
		DataHash:        make([]string, len(changes)),
		ModifierID:      make([]int32, len(changes)),
		CalculationType: make([]database.CalculationType, len(changes)),
		Value:           make([]float32, len(changes)),
	}

	for i, c := range changes {
		params.DataHash[i] = sh.GenerateDataHash(c)
		params.ModifierID[i] = c.ModifierID
		params.CalculationType[i] = database.CalculationType(c.CalculationType)
		params.Value[i] = c.Value
	}

	dbRows, err := qtx.CreateModifierChangeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create modifier changes: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractModifierChanges() ([]seedtypes.ModifierChange, error) {
	changes := []seedtypes.ModifierChange{}

	for i := range sh.RawJSON.PlayerAbilities {
		ability := &sh.RawJSON.PlayerAbilities[i]

		newChanges, err := prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.OverdriveAbilities {
		ability := &sh.RawJSON.OverdriveAbilities[i]

		newChanges, err := prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.Items {
		item := &sh.RawJSON.Items[i]

		newChanges, err := prepareAbilityModifierChanges(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.TriggerCommands {
		command := &sh.RawJSON.TriggerCommands[i]

		newChanges, err := prepareAbilityModifierChanges(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.MiscAbilities {
		ability := &sh.RawJSON.MiscAbilities[i]

		newChanges, err := prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.EnemyAbilities {
		ability := &sh.RawJSON.EnemyAbilities[i]

		newChanges, err := prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.AutoAbilities {
		autoAbility := &sh.RawJSON.AutoAbilities[i]

		newChanges, err := prepareModifierChanges(autoAbility.ModifierChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range sh.RawJSON.Properties {
		property := &sh.RawJSON.Properties[i]

		if property.ModifierChange == nil {
			continue
		}

		change, err := prepareModifierChange(property.ModifierChange)
		if err != nil {
			return nil, err
		}

		changes = append(changes, *change)
	}

	for i := range sh.RawJSON.StatusConditions {
		status := &sh.RawJSON.StatusConditions[i]

		newChanges, err := prepareModifierChanges(status.ModifierChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	return sh.DedupeRows(changes, sh.Hashes), nil
}

func prepareAbilityModifierChanges(battleInteractions []seedtypes.BattleInteraction) ([]seedtypes.ModifierChange, error) {
	changes := []seedtypes.ModifierChange{}

	for i := range battleInteractions {
		bi := &battleInteractions[i]

		changesNew, err := prepareModifierChanges(bi.ModifierChanges)
		if err != nil {
			return nil, err
		}
		changes = append(changes, changesNew...)
	}

	return changes, nil
}

func prepareModifierChanges(changes []seedtypes.ModifierChange) ([]seedtypes.ModifierChange, error) {
	changesNew := []seedtypes.ModifierChange{}

	for i := range changes {
		change, err := prepareModifierChange(&changes[i])
		if err != nil {
			return nil, err
		}

		changesNew = append(changesNew, *change)
	}

	return changesNew, nil
}

func prepareModifierChange(mc *seedtypes.ModifierChange) (*seedtypes.ModifierChange, error) {
	var err error

	if mc == nil {
		return nil, nil
	}

	mc.ModifierID, err = sh.AssignFK(mc.ModifierName, lookup.L.Modifiers)
	if err != nil {
		return nil, err
	}

	return mc, nil
}
