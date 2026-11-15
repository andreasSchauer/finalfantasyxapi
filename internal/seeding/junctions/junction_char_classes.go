package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func SeedJuncCharacterClassesPlayerUnits(qtx *database.Queries, ctx context.Context) error {
	const desc string = "character classes + player units"
	jParams, err := processJunctions(desc, sh.RawJSON.CharacterClasses, getCharacterClassPlayerUnits)
	if err != nil {
		return err
	}

	return qtx.CreateCharacterClassPlayerUnitsJunctionBulk(ctx, database.CreateCharacterClassPlayerUnitsJunctionBulkParams{
		DataHash: jParams.DataHashes,
		ClassID:  jParams.ParentIDs,
		UnitID:   jParams.ChildIDs,
	})
}

func getCharacterClassPlayerUnits(c seedtypes.CharacterClass) ([]seedtypes.PlayerUnit, error) {
	return lookup.GetResources(c.Members, lookup.L.PlayerUnits)
}
