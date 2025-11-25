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

	modifierChange.ModifierID, err = assignFK(modifierChange.ModifierName, l.getModifier)
	if err != nil {
		return ModifierChange{}, h.GetErr(modifierChange.Error(), err)
	}

	dbModifierChange, err := qtx.CreateModifierChange(context.Background(), database.CreateModifierChangeParams{
		DataHash:        generateDataHash(modifierChange),
		ModifierID:      modifierChange.ModifierID,
		CalculationType: database.CalculationType(modifierChange.CalculationType),
		Value:           modifierChange.Value,
	})
	if err != nil {
		return ModifierChange{}, h.GetErr(modifierChange.Error(), err, "couldn't create modifier change")
	}
	modifierChange.ID = dbModifierChange.ID

	return modifierChange, nil
}
