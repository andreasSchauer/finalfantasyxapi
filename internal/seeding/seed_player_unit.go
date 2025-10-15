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

func (pu PlayerUnit) GetID() int32 {
	return pu.ID
}


func (l *lookup) seedPlayerUnit(qtx *database.Queries, playerUnit PlayerUnit) (PlayerUnit, error) {
	dbPlayerUnit, err := qtx.CreatePlayerUnit(context.Background(), database.CreatePlayerUnitParams{
		DataHash: generateDataHash(playerUnit),
		Name:     playerUnit.Name,
		Type:     playerUnit.Type,
	})
	if err != nil {
		return PlayerUnit{}, fmt.Errorf("couldn't create Player Unit: %s: %v", playerUnit.Name, err)
	}

	playerUnit.ID = dbPlayerUnit.ID

	return playerUnit, nil
}