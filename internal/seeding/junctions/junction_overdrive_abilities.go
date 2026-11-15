package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getOverdriveAbilityRelatedStats(oa seedtypes.OverdriveAbility) ([]seedtypes.Stat, error) {
	return lookup.GetResources(oa.RelatedStats, lookup.L.Stats)
}

func SeedJuncOverdriveAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrive abilities + related stats"
	jParams, err := processJunctions(desc, sh.RawJSON.OverdriveAbilities, getOverdriveAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateOverdriveAbilitiesRelatedStatsJunctionBulk(ctx, database.CreateOverdriveAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		OverdriveAbilityID: jParams.ParentIDs,
		StatID:             jParams.ChildIDs,
	})
}
