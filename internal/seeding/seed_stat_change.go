package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type StatChange struct {
	StatID			int32
	StatName		string	`json:"name"`
	CalculationType	string	`json:"calculation_type"`
	Value			float32	`json:"value"`
}


func (s StatChange) ToHashFields() []any {
	return []any{
		s.StatID,
		s.CalculationType,
		s.Value,
	}
}


func (l *lookup) seedStatChange(qtx *database.Queries, statChange StatChange) (database.StatChange, error) {
	stat, err := l.getStat(statChange.StatName)
	if err != nil {
		return database.StatChange{}, err
	}

	statChange.StatID = stat.ID

	dbStatChange, err := qtx.CreateStatChange(context.Background(), database.CreateStatChangeParams{
		DataHash: 			generateDataHash(statChange),
		StatID: 			statChange.StatID,
		CalculationType: 	database.CalculationType(statChange.CalculationType),
		Value: 				statChange.Value,
	})
	if err != nil {
		return database.StatChange{}, err
	}

	return dbStatChange, nil
}