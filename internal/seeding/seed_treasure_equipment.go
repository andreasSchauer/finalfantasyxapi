package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop6SeedTreasureEquipment(qtx *database.Queries, ctx context.Context) error {
	equipment, err := l.extractTreasureEquipment()
	if err != nil {
		return err
	}

	params := database.CreateTreasureEquipmentPieceBulkParams{
		DataHash:         make([]string, len(equipment)),
		TreasureID:       make([]int32, len(equipment)),
		EquipmentNameID:  make([]int32, len(equipment)),
		EmptySlotsAmount: make([]int32, len(equipment)),
	}

	for i, e := range equipment {
		params.DataHash[i] = generateDataHash(e)
		params.TreasureID[i] = e.TreasureID
		params.EquipmentNameID[i] = e.EquipmentNameID
		params.EmptySlotsAmount[i] = e.EmptySlotsAmount
	}

	dbRows, err := qtx.CreateTreasureEquipmentPieceBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create treasure equipment: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractTreasureEquipment() ([]TreasureEquipment, error) {
	equipment := []TreasureEquipment{}
	var err error

	for i := range l.json.treasureLists {
		list := &l.json.treasureLists[i]

		for j := range list.Treasures {
			treasure := &list.Treasures[j]

			if treasure.Equipment == nil {
				continue
			}

			treasure.Equipment.TreasureID, err = l.getHashID(treasure)
			if err != nil {
				return nil, err
			}

			treasure.Equipment.EquipmentNameID, err = assignFK(treasure.Equipment.Name, l.EquipmentNames)
			if err != nil {
				return nil, err
			}

			equipment = append(equipment, *treasure.Equipment)
		}
	}

	return dedupeRows(equipment, l.Hashes), nil
}
