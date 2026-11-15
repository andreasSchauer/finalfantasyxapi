package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getOverdriveModeActions(om seedtypes.OverdriveMode) ([]seedtypes.ActionToLearn, error) {
	return om.ActionsToLearn, nil
}

func SeedJuncOverdriveModeActions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrive modes + actions"
	jParams, err := processJunctions(desc, rawJSON.overdriveModes, getOverdriveModeActions)
	if err != nil {
		return err
	}

	return qtx.CreateOverdriveModesActionsToLearnJunctionBulk(ctx, database.CreateOverdriveModesActionsToLearnJunctionBulkParams{
		DataHash:        jParams.DataHashes,
		OverdriveModeID: jParams.ParentIDs,
		ActionID:        jParams.ChildIDs,
	})
}
