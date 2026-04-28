package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type IsPlayerUnit interface {
	GetPlayerUnitParams() PlayerUnitParams
}

type PlayerUnitParams struct {
	ID		int32
	Name	string
	Type	string
	
}

type PlayerUnit struct {
	ID   int32
	Name string `json:"name"`
	Type database.UnitType
}

func (pu PlayerUnit) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", pu),
		pu.Name,
		pu.Type,
	}
}

func (pu PlayerUnit) ToKeyFields() []any {
	return []any{
		pu.Name,
	}
}

func (pu PlayerUnit) GetID() int32 {
	return pu.ID
}

func (pu PlayerUnit) Error() string {
	return fmt.Sprintf("player unit %s, type %s", pu.Name, pu.Type)
}

func (pu PlayerUnit) GetResParamsTyped() h.ResParamsTyped {
	return h.ResParamsTyped{
		ID: 	pu.ID,
		Name: 	pu.Name,
		Type: 	string(pu.Type),
	}
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

func (l *Lookup) loop1SeedPlayerUnits(qtx *database.Queries, ctx context.Context) error {
	units := l.extractPlayerUnits()

	params := database.CreatePlayerUnitBulkParams{
		DataHash: make([]string, len(units)),
		Name:     make([]string, len(units)),
		Type:     make([]database.UnitType, len(units)),
	}

	for i, m := range units {
		params.DataHash[i] = generateDataHash(m)
		params.Name[i] = m.Name
		params.Type[i] = m.Type
	}

	dbRows, err := qtx.CreatePlayerUnitBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create player units: %v", err)
	}

	for i, row := range dbRows {
		units[i].ID = row.ID
		key := CreateLookupKey(units[i])
		l.PlayerUnits[key] = units[i]
		l.PlayerUnitsID[row.ID] = units[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractPlayerUnits() []PlayerUnit {
	playerUnits := []PlayerUnit{}

	for _, c := range l.json.characters {
		c.PlayerUnit.Type = database.UnitTypeCharacter
		playerUnits = append(playerUnits, c.PlayerUnit)
	}

	for _, a := range l.json.aeons {
		a.PlayerUnit.Type = database.UnitTypeAeon
		playerUnits = append(playerUnits, a.PlayerUnit)
	}

	return dedupeRows(playerUnits, l.Hashes)
}