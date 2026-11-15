package seeding

import (
	"context"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func seedJuncAeonAeonEquipment(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + aeon equipment"
	jParams, err := processJunctions(desc, rawJSON.aeons, getAeonAeonEquipment)
	if err != nil {
		return err
	}

	return qtx.CreateAeonsWeaponArmorJunctionBulk(ctx, database.CreateAeonsWeaponArmorJunctionBulkParams{
		DataHash:        jParams.DataHashes,
		AeonID:          jParams.ParentIDs,
		AeonEquipmentID: jParams.ChildIDs,
	})
}


func getAeonAeonEquipment(a seedtypes.Aeon) ([]seedtypes.AeonEquipment, error) {
	return slices.Concat(a.Weapon, a.Armor), nil
}
