package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop6SeedAeonEquipment(qtx *database.Queries, ctx context.Context) error {
	equipment, err := l.extractAeonEquipment()
	if err != nil {
		return err
	}

	params := database.CreateAeonEquipmentBulkParams{
		DataHash:      make([]string, len(equipment)),
		AutoAbilityID: make([]int32, len(equipment)),
		CelestialWpn:  make([]bool, len(equipment)),
		EquipType:     make([]database.EquipType, len(equipment)),
	}

	for i, e := range equipment {
		params.DataHash[i] = generateDataHash(e)
		params.AutoAbilityID[i] = e.AutoAbilityID
		params.CelestialWpn[i] = e.CelestialWeapon
		params.EquipType[i] = database.EquipType(e.EquipType)
	}

	dbRows, err := qtx.CreateAeonEquipmentBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create aeon equipment: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAeonEquipment() ([]AeonEquipment, error) {
	equipment := []AeonEquipment{}
	var err error

	for i := range l.json.aeons {
		aeon := &l.json.aeons[i]

		for j := range aeon.Weapon {
			ae := &aeon.Weapon[j]
			ae.EquipType = string(database.EquipTypeWeapon)

			ae.AutoAbilityID, err = assignFK(ae.AutoAbility, l.AutoAbilities)
			if err != nil {
				return nil, err
			}

			equipment = append(equipment, *ae)
		}

		for j := range aeon.Armor {
			ae := &aeon.Armor[j]
			ae.EquipType = string(database.EquipTypeArmor)

			ae.AutoAbilityID, err = assignFK(ae.AutoAbility, l.AutoAbilities)
			if err != nil {
				return nil, err
			}

			equipment = append(equipment, *ae)
		}
	}

	return dedupeRows(equipment, l.Hashes), nil
}
