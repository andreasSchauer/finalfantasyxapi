package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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

func (s StatChange) Error() string {
	return fmt.Sprintf("stat change with stat: %s, calc type: %s, value %f", s.StatName, s.CalculationType, s.Value)
}

func (l *Lookup) seedStatChange(qtx *database.Queries, statChange StatChange) (StatChange, error) {
	var err error

	statChange.StatID, err = assignFK(statChange.StatName, l.Stats)
	if err != nil {
		return StatChange{}, h.NewErr(statChange.Error(), err)
	}

	dbStatChange, err := qtx.CreateStatChange(context.Background(), database.CreateStatChangeParams{
		DataHash:        generateDataHash(statChange),
		StatID:          statChange.StatID,
		CalculationType: database.CalculationType(statChange.CalculationType),
		Value:           statChange.Value,
	})
	if err != nil {
		return StatChange{}, h.NewErr(statChange.Error(), err, "couldn't create stat change")
	}
	statChange.ID = dbStatChange.ID

	return statChange, nil
}
