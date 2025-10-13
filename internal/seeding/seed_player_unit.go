package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type PlayerUnit struct {
	ID		int32
	Name 	string 				`json:"name"`
	Type 	database.UnitType
}

func (pu PlayerUnit) ToHashFields() []any {
	return []any{
		pu.Name,
		pu.Type,
	}
}

func (pu PlayerUnit) ToKeyFields() []any {
	return []any{
		pu.Name,
		pu.Type,
	}
}


func (l *lookup) seedPlayerUnit(qtx *database.Queries, unit PlayerUnit) (database.PlayerUnit, error) {
	dbPlayerUnit, err := qtx.CreatePlayerUnit(context.Background(), database.CreatePlayerUnitParams{
		DataHash: generateDataHash(unit),
		Name:     unit.Name,
		Type:     unit.Type,
	})
	if err != nil {
		return database.PlayerUnit{}, fmt.Errorf("couldn't create Player Unit: %s: %v", unit.Name, err)
	}

	return dbPlayerUnit, nil
}