package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getPropertyRelatedStats(p seedtypes.Property) ([]seedtypes.Stat, error) {
	return lookup.GetResources(p.RelatedStats, lookup.L.Stats)
}

func SeedJuncPropertiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "properties + related stats"
	jParams, err := processJunctions(desc, rawJSON.properties, getPropertyRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreatePropertiesRelatedStatsJunctionBulk(ctx, database.CreatePropertiesRelatedStatsJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		PropertyID: jParams.ParentIDs,
		StatID:     jParams.ChildIDs,
	})
}
