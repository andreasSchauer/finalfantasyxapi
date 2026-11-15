package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func SeedJuncCuesIncludedAreas(qtx *database.Queries, ctx context.Context) error {
	const desc string = "cues + included areas"
	jParams, err := processJunctions(desc, getCues(), getCueIncludedAreas)
	if err != nil {
		return err
	}

	return qtx.CreateCuesIncludedAreasJunctionBulk(ctx, database.CreateCuesIncludedAreasJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		CueID:          jParams.ParentIDs,
		IncludedAreaID: jParams.ChildIDs,
	})
}

func getCues() []seedtypes.Cue {
	cues := []seedtypes.Cue{}

	for _, song := range sh.RawJSON.Songs {
		cues = append(cues, song.Cues...)
	}

	return cues
}

func getCueIncludedAreas(c seedtypes.Cue) ([]seedtypes.Area, error) {
	return lookup.GetResources(c.IncludedAreas, lookup.L.Areas)
}
