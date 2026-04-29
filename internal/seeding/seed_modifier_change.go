package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type ModifierChange struct {
	ID              int32
	ModifierID      int32
	ModifierName    string  `json:"name"`
	CalculationType string  `json:"calculation_type"`
	Value           float32 `json:"value"`
}

func (m ModifierChange) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.ModifierID,
		m.CalculationType,
		m.Value,
	}
}

func (m ModifierChange) GetID() int32 {
	return m.ID
}

func (m ModifierChange) Error() string {
	return fmt.Sprintf("modifier change with modifier: %s, calc type: %s, value %f", m.ModifierName, m.CalculationType, m.Value)
}

func (l *Lookup) seedModifierChange(qtx *database.Queries, modifierChange ModifierChange) (ModifierChange, error) {
	var err error

	modifierChange.ModifierID, err = assignFK(modifierChange.ModifierName, l.Modifiers)
	if err != nil {
		return ModifierChange{}, h.NewErr(modifierChange.Error(), err)
	}

	dbModifierChange, err := qtx.CreateModifierChange(context.Background(), database.CreateModifierChangeParams{
		DataHash:        generateDataHash(modifierChange),
		ModifierID:      modifierChange.ModifierID,
		CalculationType: database.CalculationType(modifierChange.CalculationType),
		Value:           modifierChange.Value,
	})
	if err != nil {
		return ModifierChange{}, h.NewErr(modifierChange.Error(), err, "couldn't create modifier change")
	}
	modifierChange.ID = dbModifierChange.ID

	return modifierChange, nil
}


func (l *Lookup) loop2SeedModifierChanges(qtx *database.Queries, ctx context.Context) error {
	changes, err := l.extractModifierChanges()
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
		params.DataHash[i] = generateDataHash(c)
		params.ModifierID[i] = c.ModifierID
		params.CalculationType[i] = database.CalculationType(c.CalculationType)
		params.Value[i] = c.Value
	}

	dbRows, err := qtx.CreateModifierChangeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create modifier changes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractModifierChanges() ([]ModifierChange, error) {
	changes := []ModifierChange{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		newChanges, err := l.prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		newChanges, err := l.prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		newChanges, err := l.prepareAbilityModifierChanges(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		newChanges, err := l.prepareAbilityModifierChanges(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		newChanges, err := l.prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		newChanges, err := l.prepareAbilityModifierChanges(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		newChanges, err := l.prepareModifierChanges(autoAbility.ModifierChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.properties {
		property := &l.json.properties[i]

		newChanges, err := l.prepareModifierChanges(property.ModifierChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	for i := range l.json.statusConditions {
		status := &l.json.statusConditions[i]

		newChanges, err := l.prepareModifierChanges(status.ModifierChanges)
		if err != nil {
			return nil, err
		}

		changes = append(changes, newChanges...)
	}

	return dedupeRows(changes, l.Hashes), nil
}

func (l *Lookup) prepareAbilityModifierChanges(battleInteractions []BattleInteraction) ([]ModifierChange, error) {
	changes := []ModifierChange{}

	for i := range battleInteractions {
		bi := &battleInteractions[i]

		changesNew, err := l.prepareModifierChanges(bi.ModifierChanges)
		if err != nil {
			return nil, err
		}
		changes = append(changes, changesNew...)
	}

	return changes, nil
}

func (l *Lookup) prepareModifierChanges(changes []ModifierChange) ([]ModifierChange, error) {
	changesNew := []ModifierChange{}
	var err error

	for i := range changes {
		change := &changes[i]

		change.ModifierID, err = assignFK(change.ModifierName, l.Modifiers)
		if err != nil {
			return nil, err
		}

		changesNew = append(changesNew, *change)
	}

	return changesNew, nil
}