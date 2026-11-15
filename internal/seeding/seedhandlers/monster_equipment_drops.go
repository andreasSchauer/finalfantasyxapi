package seedhandlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop6EquipmentDrops(qtx *database.Queries, ctx context.Context) error {
	drops, err := extractEquipmentDrops()
	if err != nil {
		return err
	}

	params := database.CreateEquipmentDropBulkParams{
		DataHash:      make([]string, len(drops)),
		AutoAbilityID: make([]int32, len(drops)),
		IsForced:      make([]bool, len(drops)),
		Probability:   make([]sql.NullInt32, len(drops)),
		Type:          make([]database.EquipType, len(drops)),
	}

	for i, d := range drops {
		params.DataHash[i] = sh.GenerateDataHash(d)
		params.AutoAbilityID[i] = d.AutoAbilityID
		params.IsForced[i] = d.IsForced
		params.Probability[i] = h.GetNullInt32(d.Probability)
		params.Type[i] = d.Type
	}

	dbRows, err := qtx.CreateEquipmentDropBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create equipment drops: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractEquipmentDrops() ([]seedtypes.EquipmentDrop, error) {
	drops := []seedtypes.EquipmentDrop{}
	var err error

	for i := range sh.RawJSON.Monsters {
		mon := &sh.RawJSON.Monsters[i]

		if mon.Equipment == nil {
			continue
		}

		for j := range mon.Equipment.WeaponAbilities {
			drop := &mon.Equipment.WeaponAbilities[j]
			drop.Type = database.EquipTypeWeapon

			drop.AutoAbilityID, err = sh.AssignFK(drop.Ability, lookup.L.AutoAbilities)
			if err != nil {
				return nil, err
			}

			drops = append(drops, *drop)
		}

		for j := range mon.Equipment.ArmorAbilities {
			drop := &mon.Equipment.ArmorAbilities[j]
			drop.Type = database.EquipTypeArmor

			drop.AutoAbilityID, err = sh.AssignFK(drop.Ability, lookup.L.AutoAbilities)
			if err != nil {
				return nil, err
			}

			drops = append(drops, *drop)
		}
	}

	return sh.DedupeRows(drops, sh.Hashes), nil
}
