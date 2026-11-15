package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)


func getPlayerAbilityLearnedBy(pa seedtypes.PlayerAbility) ([]seedtypes.CharacterClass, error) {
	return lookup.GetResources(pa.LearnedBy, lookup.L.CharClasses)
}

func seedJuncPlayerAbilitiesLearnedBy(qtx *database.Queries, ctx context.Context) error {
	const desc string = "player abilities + learned by"
	jParams, err := processJunctions(desc, rawJSON.playerAbilities, getPlayerAbilityLearnedBy)
	if err != nil {
		return err
	}

	return qtx.CreatePlayerAbilitiesLearnedByJunctionBulk(ctx, database.CreatePlayerAbilitiesLearnedByJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		PlayerAbilityID:  jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}

func getPlayerAbilityRelatedStats(pa seedtypes.PlayerAbility) ([]seedtypes.Stat, error) {
	return lookup.GetResources(pa.RelatedStats, lookup.L.Stats)
}

func seedJuncPlayerAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "player abilites + related stats"
	jParams, err := processJunctions(desc, rawJSON.playerAbilities, getPlayerAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreatePlayerAbilitiesRelatedStatsJunctionBulk(ctx, database.CreatePlayerAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:        jParams.DataHashes,
		PlayerAbilityID: jParams.ParentIDs,
		StatID:          jParams.ChildIDs,
	})
}
