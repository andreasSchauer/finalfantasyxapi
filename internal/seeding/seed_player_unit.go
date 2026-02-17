package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type PlayerUnit struct {
	ID   int32
	Name string `json:"name"`
	Type database.UnitType
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

func (pu PlayerUnit) Error() string {
	return fmt.Sprintf("player unit %s, type %s", pu.Name, pu.Type)
}

func (l *Lookup) seedPlayerUnit(qtx *database.Queries, playerUnit PlayerUnit) (PlayerUnit, error) {
	dbPlayerUnit, err := qtx.CreatePlayerUnit(context.Background(), database.CreatePlayerUnitParams{
		DataHash: generateDataHash(playerUnit),
		Name:     playerUnit.Name,
		Type:     playerUnit.Type,
	})
	if err != nil {
		return PlayerUnit{}, h.NewErr(playerUnit.Error(), err, "couldn't create player unit")
	}

	playerUnit.ID = dbPlayerUnit.ID
	l.PlayerUnits[playerUnit.Name] = playerUnit
	l.PlayerUnitsID[playerUnit.ID] = playerUnit

	return playerUnit, nil
}
