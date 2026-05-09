package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedProperties(qtx *database.Queries, ctx context.Context) error {
	properties, err := l.extractProperties()
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
		params.DataHash[i] = generateDataHash(p)
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
		l.json.properties[i].ID = row.ID
		l.Properties[Key(properties[i])] = properties[i]
		l.PropertiesID[row.ID] = properties[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractProperties() ([]Property, error) {
	properties := []Property{}
	var err error

	for i := range l.json.properties {
		property := &l.json.properties[i]

		if property.ModifierChange != nil {
			property.ModifierChange.ID, err = l.GetHashID(property.ModifierChange)
			if err != nil {
				return nil, err
			}
		}

		properties = append(properties, *property)
	}

	return dedupeRows(properties, l.Hashes), nil
}

func (l *Lookup) getPropertyRelatedStats(p Property) ([]Stat, error) {
	return getResources(p.RelatedStats, l.Stats)
}

func (l *Lookup) seedJuncPropertiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "properties + related stats"
	jParams, err := processJunctions(l, desc, l.json.properties, l.getPropertyRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreatePropertiesRelatedStatsJunctionBulk(ctx, database.CreatePropertiesRelatedStatsJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		PropertyID: jParams.ParentIDs,
		StatID:     jParams.ChildIDs,
	})
}
