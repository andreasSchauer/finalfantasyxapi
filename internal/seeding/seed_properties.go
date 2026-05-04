package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop1SeedProperties(qtx *database.Queries, ctx context.Context) error {
	properties := dedupeRows(l.json.properties, l.Hashes)

	params := database.CreatePropertyBulkParams{
		DataHash:       make([]string, len(properties)),
		Name:           make([]string, len(properties)),
		Effect:         make([]string, len(properties)),
		NullifyArmored: make([]database.NullNullifyArmored, len(properties)),
	}

	for i, p := range properties {
		params.DataHash[i] = generateDataHash(p)
		params.Name[i] = p.Name
		params.Effect[i] = p.Effect
		params.NullifyArmored[i] = database.ToNullNullifyArmored(p.NullifyArmored)
	}

	dbRows, err := qtx.CreatePropertyBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create properties: %v", err)
	}

	for i, row := range dbRows {
		properties[i].ID = row.ID
		l.json.properties[i].ID = row.ID
		l.Properties[properties[i].Name] = properties[i]
		l.PropertiesID[row.ID] = properties[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) completeProperties() error {
	for i := range l.json.properties {
		property := &l.json.properties[i]

		err := assignIDs(l, property.ModifierChanges)
		if err != nil {
			return err
		}

		l.Properties[property.Name] = *property
		l.PropertiesID[property.ID] = *property
	}

	return nil
}

func (l *Lookup) getPropertyModifierChanges(p Property) ([]ModifierChange, error) {
	return p.ModifierChanges, nil
}

func (l *Lookup) seedJuncPropertiesModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "properties + modifier changes"
	jParams, err := processJunctions(l, desc, l.json.properties, l.getPropertyModifierChanges)
	if err != nil {
		return err
	}

	return qtx.CreatePropertiesModifierChangesJunctionBulk(ctx, database.CreatePropertiesModifierChangesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		PropertyID:       jParams.ParentIDs,
		ModifierChangeID: jParams.ChildIDs,
	})
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
