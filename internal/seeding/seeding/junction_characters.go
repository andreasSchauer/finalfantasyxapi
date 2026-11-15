package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func SeedJuncCharactersBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "characters + base stats"
	jParams, err := processJunctions(desc, seedhelpers.RawJSON.Characters, getCharacterBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateCharactersBaseStatsJunctionBulk(ctx, database.CreateCharactersBaseStatsJunctionBulkParams{
		DataHash:    jParams.DataHashes,
		CharacterID: jParams.ParentIDs,
		BaseStatID:  jParams.ChildIDs,
	})
}

func getCharacterBaseStats(c seedtypes.Character) ([]seedtypes.BaseStat, error) {
	return c.BaseStats, nil
}
