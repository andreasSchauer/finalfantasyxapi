package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)


func getOverdriveAbilityRelatedStats(oa seedtypes.OverdriveAbility) ([]seedtypes.Stat, error) {
	return lookup.GetResources(oa.RelatedStats, lookup.L.Stats)
}

func seedJuncOverdriveAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrive abilities + related stats"
	jParams, err := processJunctions(desc, rawJSON.overdriveAbilities, getOverdriveAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateOverdriveAbilitiesRelatedStatsJunctionBulk(ctx, database.CreateOverdriveAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		OverdriveAbilityID: jParams.ParentIDs,
		StatID:             jParams.ChildIDs,
	})
}
