package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)



type ModifierChange struct {
	ModifierID		int32
	ModifierName	string	`json:"name"`
	CalculationType	string	`json:"calculation_type"`
	Value			float32	`json:"value"`
}


func (m ModifierChange) ToHashFields() []any {
	return []any{
		m.ModifierID,
		m.CalculationType,
		m.Value,
	}
}


func (l *lookup) seedModifierChange(qtx *database.Queries, modifierChange ModifierChange) (database.ModifierChange, error) {
	modifier, err := l.getModifier(modifierChange.ModifierName)
	if err != nil {
		return database.ModifierChange{}, err
	}
	
	modifierChange.ModifierID = modifier.ID

	dbModifierChange, err := qtx.CreateModifierChange(context.Background(), database.CreateModifierChangeParams{
		DataHash: 			generateDataHash(modifierChange),
		ModifierID: 		modifierChange.ModifierID,
		CalculationType: 	database.CalculationType(modifierChange.CalculationType),
		Value: 				modifierChange.Value,
	})
	if err != nil {
		return database.ModifierChange{}, err
	}

	return dbModifierChange, nil
}