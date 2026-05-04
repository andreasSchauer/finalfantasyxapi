package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop6SeedEquipmentDrops(qtx *database.Queries, ctx context.Context) error {
	drops, err := l.extractEquipmentDrops()
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
		params.DataHash[i] = generateDataHash(d)
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractEquipmentDrops() ([]EquipmentDrop, error) {
	drops := []EquipmentDrop{}
	var err error

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		if mon.Equipment == nil {
			continue
		}

		for j := range mon.Equipment.WeaponAbilities {
			drop := &mon.Equipment.WeaponAbilities[j]
			drop.Type = database.EquipTypeWeapon

			drop.AutoAbilityID, err = assignFK(drop.Ability, l.AutoAbilities)
			if err != nil {
				return nil, err
			}

			drops = append(drops, *drop)
		}

		for j := range mon.Equipment.ArmorAbilities {
			drop := &mon.Equipment.ArmorAbilities[j]
			drop.Type = database.EquipTypeArmor

			drop.AutoAbilityID, err = assignFK(drop.Ability, l.AutoAbilities)
			if err != nil {
				return nil, err
			}

			drops = append(drops, *drop)
		}
	}

	return dedupeRows(drops, l.Hashes), nil
}
