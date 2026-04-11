package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func convertGetCelestialWeaponAutoAbilityIDs(cfg *Config) DbQueryIntMany {
	return func (ctx context.Context, id int32) ([]int32, error) {
		cw, _ := seeding.GetResourceByID(id, cfg.l.CelestialWeaponsID)

		equipment, _ := seeding.GetResource(cw.Name, cfg.l.EquipmentNames)

		dbTableIDs, err := cfg.db.GetEquipmentEquipmentTableIDs(ctx, equipment.ID)
		if err != nil {
			return nil, err
		}
		table, _ := seeding.GetResourceByID(dbTableIDs[0], cfg.l.EquipmentTablesID)

		return cfg.db.GetCelestialWeaponAutoAbilityIDs(ctx, database.GetCelestialWeaponAutoAbilityIDsParams{
			CelestialWeaponID: 	cw.ID,
			EquipmentTableID: 	table.ID,
		})
	}
}