package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop1SeedPlayerUnits(qtx *database.Queries, ctx context.Context) error {
	units := extractPlayerUnits()

	params := database.CreatePlayerUnitBulkParams{
		DataHash: make([]string, len(units)),
		Name:     make([]string, len(units)),
		Type:     make([]database.UnitType, len(units)),
	}

	for i, m := range units {
		params.DataHash[i] = sh.GenerateDataHash(m)
		params.Name[i] = m.Name
		params.Type[i] = m.Type
	}

	dbRows, err := qtx.CreatePlayerUnitBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create player units: %v", err)
	}

	for i, row := range dbRows {
		units[i].ID = row.ID
		lookup.L.PlayerUnits[seedtypes.Key(units[i])] = units[i]
		lookup.L.PlayerUnitsID[row.ID] = units[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractPlayerUnits() []seedtypes.PlayerUnit {
	playerUnits := []seedtypes.PlayerUnit{}

	for i := range sh.RawJSON.Characters {
		c := &sh.RawJSON.Characters[i]
		c.PlayerUnit.Type = database.UnitTypeCharacter
		playerUnits = append(playerUnits, c.PlayerUnit)
	}

	for i := range sh.RawJSON.Aeons {
		a := &sh.RawJSON.Aeons[i]
		a.PlayerUnit.Type = database.UnitTypeAeon
		playerUnits = append(playerUnits, a.PlayerUnit)
	}

	return sh.DedupeRows(playerUnits, sh.Hashes)
}
