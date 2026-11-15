package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop6SeedTreasureEquipment(qtx *database.Queries, ctx context.Context) error {
	equipment, err := extractTreasureEquipment()
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
		params.DataHash[i] = sh.GenerateDataHash(e)
		params.TreasureID[i] = e.TreasureID
		params.EquipmentNameID[i] = e.EquipmentNameID
		params.EmptySlotsAmount[i] = e.EmptySlotsAmount
	}

	dbRows, err := qtx.CreateTreasureEquipmentPieceBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create treasure equipment: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractTreasureEquipment() ([]seedtypes.TreasureEquipment, error) {
	equipment := []seedtypes.TreasureEquipment{}
	var err error

	for i := range sh.RawJSON.TreasureLists {
		list := &sh.RawJSON.TreasureLists[i]

		for j := range list.Treasures {
			treasure := &list.Treasures[j]

			if treasure.Equipment == nil {
				continue
			}

			treasure.Equipment.TreasureID, err = sh.GetHashID(treasure)
			if err != nil {
				return nil, err
			}

			treasure.Equipment.EquipmentNameID, err = sh.AssignFK(treasure.Equipment.Name, lookup.L.EquipmentNames)
			if err != nil {
				return nil, err
			}

			equipment = append(equipment, *treasure.Equipment)
		}
	}

	return sh.DedupeRows(equipment, sh.Hashes), nil
}
