package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop1SeedCharacterClasses(qtx *database.Queries, ctx context.Context) error {
	classes := dedupeRows(rawJSON.characterClasses, hashes)

	params := database.CreateCharacterClassBulkParams{
		DataHash: make([]string, len(classes)),
		Name:     make([]string, len(classes)),
		Category: make([]database.CharacterClassCategory, len(classes)),
	}

	for i, c := range classes {
		params.DataHash[i] = generateDataHash(c)
		params.Name[i] = c.Name
		params.Category[i] = database.CharacterClassCategory(c.Category)
	}

	dbRows, err := qtx.CreateCharacterClassBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create character classes: %v", err)
	}

	for i, row := range dbRows {
		classes[i].ID = row.ID
		rawJSON.characterClasses[i].ID = row.ID
		L.CharClasses[seedtypes.Key(classes[i])] = classes[i]
		L.CharClassesID[row.ID] = classes[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func getCharacterClassPlayerUnits(c seedtypes.CharacterClass) ([]seedtypes.PlayerUnit, error) {
	return GetResources(c.Members, L.PlayerUnits)
}

func seedJuncCharacterClassesPlayerUnits(qtx *database.Queries, ctx context.Context) error {
	const desc string = "character classes + player units"
	jParams, err := processJunctions(desc, rawJSON.characterClasses, getCharacterClassPlayerUnits)
	if err != nil {
		return err
	}

	return qtx.CreateCharacterClassPlayerUnitsJunctionBulk(ctx, database.CreateCharacterClassPlayerUnitsJunctionBulkParams{
		DataHash: jParams.DataHashes,
		ClassID:  jParams.ParentIDs,
		UnitID:   jParams.ChildIDs,
	})
}
