package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func convertGetCelestialWeaponAutoAbilityIDs(srv *Server) DBQueryIntMany {
	return func(ctx context.Context, id int32) ([]int32, error) {
		cw, _ := lookup.GetResourceByID(id, lookup.L.CelestialWeaponsID)

		equipment, _ := lookup.GetResource(cw.Name, lookup.L.EquipmentNames)

		dbTableIDs, err := srv.Cfg.DB.GetEquipmentEquipmentTableIDs(ctx, equipment.ID)
		if err != nil {
			return nil, err
		}
		table, _ := lookup.GetResourceByID(dbTableIDs[0], lookup.L.EquipmentTablesID)

		return srv.Cfg.DB.GetCelestialWeaponAutoAbilityIDs(ctx, database.GetCelestialWeaponAutoAbilityIDsParams{
			CelestialWeaponID: cw.ID,
			EquipmentTableID:  table.ID,
		})
	}
}
