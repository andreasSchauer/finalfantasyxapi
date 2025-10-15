package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type StatChange struct {
	ID              int32
	StatID          int32
	StatName        string  `json:"name"`
	CalculationType string  `json:"calculation_type"`
	Value           float32 `json:"value"`
}

func (s StatChange) ToHashFields() []any {
	return []any{
		s.StatID,
		s.CalculationType,
		s.Value,
	}
}

func (s StatChange) GetID() int32 {
	return s.ID
}

func (l *lookup) seedStatChange(qtx *database.Queries, statChange StatChange) (StatChange, error) {
	var err error
	
	statChange.StatID, err = assignFK(statChange.StatName, l.getStat)
	if err != nil {
		return StatChange{}, err
	}

	dbStatChange, err := qtx.CreateStatChange(context.Background(), database.CreateStatChangeParams{
		DataHash:        generateDataHash(statChange),
		StatID:          statChange.StatID,
		CalculationType: database.CalculationType(statChange.CalculationType),
		Value:           statChange.Value,
	})
	if err != nil {
		return StatChange{}, err
	}
	statChange.ID = dbStatChange.ID

	return statChange, nil
}
