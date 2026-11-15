package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func SeedJuncAreaConnectedAreas(qtx *database.Queries, ctx context.Context) error {
	const desc string = "area + area connections"
	jParams, err := processJunctions(desc, getAreas(), GetAreaConnectedAreas)
	if err != nil {
		return err
	}

	return qtx.CreateAreaConnectedAreasJunctionBulk(ctx, database.CreateAreaConnectedAreasJunctionBulkParams{
		DataHash:     jParams.DataHashes,
		AreaID:       jParams.ParentIDs,
		ConnectionID: jParams.ChildIDs,
	})
}

func getAreas() []seedtypes.Area {
	areas := []seedtypes.Area{}

	for _, location := range sh.RawJSON.Locations {
		for _, sublocation := range location.Sublocations {
			areas = append(areas, sublocation.Areas...)
		}
	}

	return areas
}

func GetAreaConnectedAreas(a seedtypes.Area) ([]seedtypes.AreaConnection, error) {
	return a.ConnectedAreas, nil
}
