package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getTriggerCommandRelatedStats(tc seedtypes.TriggerCommand) ([]seedtypes.Stat, error) {
	return lookup.GetResources(tc.RelatedStats, lookup.L.Stats)
}

func SeedJuncTriggerCommandsRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "trigger commands + related stats"
	jParams, err := processJunctions(desc, seedhelpers.RawJSON.TriggerCommands, getTriggerCommandRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateTriggerCommandsRelatedStatsJunctionBulk(ctx, database.CreateTriggerCommandsRelatedStatsJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		TriggerCommandID: jParams.ParentIDs,
		StatID:           jParams.ChildIDs,
	})
}
