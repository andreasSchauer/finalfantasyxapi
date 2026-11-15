package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3SeedProperties(qtx *database.Queries, ctx context.Context) error {
	properties, err := extractProperties()
	if err != nil {
		return err
	}

	params := database.CreatePropertyBulkParams{
		DataHash:         make([]string, len(properties)),
		Name:             make([]string, len(properties)),
		Effect:           make([]string, len(properties)),
		NullifyArmored:   make([]database.NullNullifyArmored, len(properties)),
		ModifierChangeID: make([]sql.NullInt32, len(properties)),
	}

	for i, p := range properties {
		params.DataHash[i] = seedhelpers.GenerateDataHash(p)
		params.Name[i] = p.Name
		params.Effect[i] = p.Effect
		params.NullifyArmored[i] = database.ToNullNullifyArmored(p.NullifyArmored)
		params.ModifierChangeID[i] = h.ObjPtrToNullInt32ID(p.ModifierChange)
	}

	dbRows, err := qtx.CreatePropertyBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create properties: %v", err)
	}

	for i, row := range dbRows {
		properties[i].ID = row.ID
		seedhelpers.RawJSON.Properties[i].ID = row.ID
		lookup.L.Properties[seedtypes.Key(properties[i])] = properties[i]
		lookup.L.PropertiesID[row.ID] = properties[i]
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractProperties() ([]seedtypes.Property, error) {
	properties := []seedtypes.Property{}
	var err error

	for i := range seedhelpers.RawJSON.Properties {
		property := &seedhelpers.RawJSON.Properties[i]

		if property.ModifierChange != nil {
			property.ModifierChange.ID, err = seedhelpers.GetHashID(property.ModifierChange)
			if err != nil {
				return nil, err
			}
		}

		properties = append(properties, *property)
	}

	return dedupeRows(properties, seedhelpers.Hashes), nil
}
