package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop5SeedStatChanges(qtx *database.Queries, ctx context.Context) error {
	changes, err := l.extractStatChanges()
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
		params.DataHash[i] = generateDataHash(c)
		params.StatID[i] = c.StatID
		params.CalculationType[i] = database.CalculationType(c.CalculationType)
		params.Value[i] = c.Value
	}

	dbRows, err := qtx.CreateStatChangeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create stat changes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractStatChanges() ([]StatChange, error) {
	changes := []StatChange{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		newChanges, err := l.prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		newChanges, err := l.prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		newChanges, err := l.prepareAbilityStatChanges(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		newChanges, err := l.prepareAbilityStatChanges(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		newChanges, err := l.prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		newChanges, err := l.prepareAbilityStatChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		newChanges, err := l.prepareStatChanges(autoAbility.StatChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.properties {
		property := &l.json.properties[i]

		newChanges, err := l.prepareStatChanges(property.StatChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.statusConditions {
		status := &l.json.statusConditions[i]

		newChanges, err := l.prepareStatChanges(status.StatChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	return dedupeRows(changes, l.Hashes), nil
}

func (l *Lookup) prepareAbilityStatChanges(battleInteractions []BattleInteraction) ([]StatChange, error) {
	changes := []StatChange{}

	for i := range battleInteractions {
		bi := &battleInteractions[i]

		changesNew, err := l.prepareStatChanges(bi.StatChanges)
		if err != nil {
			return nil, err
		}
		changes = append(changes, changesNew...)
	}

	return changes, nil
}

func (l *Lookup) prepareStatChanges(changes []StatChange) ([]StatChange, error) {
	changesNew := []StatChange{}
	var err error

	for i := range changes {
		change := &changes[i]

		change.StatID, err = assignFK(change.StatName, l.Stats)
		if err != nil {
			return nil, err
		}

		changesNew = append(changesNew, *change)
	}

	return changesNew, nil
}
