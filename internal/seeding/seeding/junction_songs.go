package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getSongBackgroundMusic(s seedtypes.Song) ([]seedtypes.BackgroundMusic, error) {
	return s.BackgroundMusic, nil
}

func getBackgroundMusicAreas(bm seedtypes.BackgroundMusic) ([]seedtypes.Area, error) {
	return lookup.GetResources(bm.LocationAreas, lookup.L.Areas)
}

func SeedJuncSongsBackgroundMusic(qtx *database.Queries, ctx context.Context) error {
	const desc string = "songs + background music"
	jParams, err := processThreewayJunctions(desc, seedhelpers.RawJSON.Songs, getSongBackgroundMusic, getBackgroundMusicAreas)
	if err != nil {
		return err
	}

	return qtx.CreateSongsBackgroundMusicJunctionBulk(ctx, database.CreateSongsBackgroundMusicJunctionBulkParams{
		DataHash: jParams.DataHashes,
		SongID:   jParams.GrandParentIDs,
		BmID:     jParams.ParentIDs,
		AreaID:   jParams.ChildIDs,
	})
}
