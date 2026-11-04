package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop1SeedModifiers(qtx *database.Queries, ctx context.Context) error {
	modifiers := dedupeRows(rawJSON.modifiers, hashes)

	params := database.CreateModifierBulkParams{
		DataHash:     make([]string, len(modifiers)),
		Name:         make([]string, len(modifiers)),
		Effect:       make([]string, len(modifiers)),
		Category:     make([]database.ModifierCategory, len(modifiers)),
		DefaultValue: make([]sql.NullFloat64, len(modifiers)),
	}

	for i, m := range modifiers {
		params.DataHash[i] = generateDataHash(m)
		params.Name[i] = m.Name
		params.Effect[i] = m.Effect
		params.Category[i] = database.ModifierCategory(m.Category)
		params.DefaultValue[i] = h.GetNullFloat64(m.DefaultValue)
	}

	dbRows, err := qtx.CreateModifierBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create modifiers: %v", err)
	}

	for i, row := range dbRows {
		modifiers[i].ID = row.ID
		rawJSON.modifiers[i].ID = row.ID
		lookup.L.Modifiers[seedtypes.Key(modifiers[i])] = modifiers[i]
		lookup.L.ModifiersID[row.ID] = modifiers[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}
