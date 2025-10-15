package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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
		m.ModifierID,
		m.CalculationType,
		m.Value,
	}
}

func (m ModifierChange) GetID() int32 {
	return m.ID
}

func (l *lookup) seedModifierChange(qtx *database.Queries, modifierChange ModifierChange) (ModifierChange, error) {
	var err error
	
	modifierChange.ModifierID, err = assignFK(modifierChange.ModifierName, l.getModifier)
	if err != nil {
		return ModifierChange{}, err
	}

	dbModifierChange, err := qtx.CreateModifierChange(context.Background(), database.CreateModifierChangeParams{
		DataHash:        generateDataHash(modifierChange),
		ModifierID:      modifierChange.ModifierID,
		CalculationType: database.CalculationType(modifierChange.CalculationType),
		Value:           modifierChange.Value,
	})
	if err != nil {
		return ModifierChange{}, err
	}
	modifierChange.ID = dbModifierChange.ID

	return modifierChange, nil
}
