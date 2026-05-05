package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop1SeedCharacterClasses(qtx *database.Queries, ctx context.Context) error {
	classes := dedupeRows(l.json.characterClasses, l.Hashes)

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
		l.json.characterClasses[i].ID = row.ID
		l.CharClasses[Key(classes[i])] = classes[i]
		l.CharClassesID[row.ID] = classes[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getCharacterClassPlayerUnits(c CharacterClass) ([]PlayerUnit, error) {
	return getResources(c.Members, l.PlayerUnits)
}

func (l *Lookup) seedJuncCharacterClassesPlayerUnits(qtx *database.Queries, ctx context.Context) error {
	const desc string = "character classes + player units"
	jParams, err := processJunctions(l, desc, l.json.characterClasses, l.getCharacterClassPlayerUnits)
	if err != nil {
		return err
	}

	return qtx.CreateCharacterClassPlayerUnitsJunctionBulk(ctx, database.CreateCharacterClassPlayerUnitsJunctionBulkParams{
		DataHash: jParams.DataHashes,
		ClassID:  jParams.ParentIDs,
		UnitID:   jParams.ChildIDs,
	})
}
