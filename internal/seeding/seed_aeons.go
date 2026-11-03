package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop4SeedAeons(qtx *database.Queries, ctx context.Context) error {
	aeons, err := extractAeons()
	if err != nil {
		return err
	}

	params := database.CreateAeonBulkParams{
		DataHash:              make([]string, len(aeons)),
		UnitID:                make([]int32, len(aeons)),
		UnlockCondition:       make([]string, len(aeons)),
		IsOptional:            make([]bool, len(aeons)),
		BattlesToRegenerate:   make([]int32, len(aeons)),
		PhysAtkDamageConstant: make([]sql.NullInt32, len(aeons)),
		PhysAtkRange:          make([]sql.NullInt32, len(aeons)),
		PhysAtkShatterRate:    make([]sql.NullInt32, len(aeons)),
		AreaID:                make([]sql.NullInt32, len(aeons)),
		AccuracyID:            make([]sql.NullInt32, len(aeons)),
	}

	for i, a := range aeons {
		params.DataHash[i] = generateDataHash(a)
		params.UnitID[i] = a.PlayerUnit.ID
		params.UnlockCondition[i] = a.UnlockCondition
		params.IsOptional[i] = a.IsOptional
		params.BattlesToRegenerate[i] = a.BattlesToRegenerate
		params.PhysAtkDamageConstant[i] = h.GetNullInt32(a.PhysAtkDmgConstant)
		params.PhysAtkRange[i] = h.GetNullInt32(a.PhysAtkRange)
		params.PhysAtkShatterRate[i] = h.GetNullInt32(a.PhysAtkShatterRate)
		params.AreaID[i] = h.GetNullInt32(a.AreaID)
		params.AccuracyID[i] = h.ObjPtrToNullInt32ID(a.PhysAtkAccuracy)
	}

	dbRows, err := qtx.CreateAeonBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create aeons: %v", err)
	}

	for i, row := range dbRows {
		aeons[i].ID = row.ID
		rawJSON.aeons[i].ID = row.ID
		L.Aeons[seedtypes.Key(aeons[i])] = aeons[i]
		L.AeonsID[row.ID] = aeons[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAeons() ([]seedtypes.Aeon, error) {
	aeons := []seedtypes.Aeon{}
	var err error

	for i := range rawJSON.aeons {
		aeon := &rawJSON.aeons[i]

		aeon.PlayerUnit.ID, err = GetHashID(aeon.PlayerUnit)
		if err != nil {
			return nil, err
		}

		aeon.AreaID, err = assignFKPtr(&aeon.LocationArea, L.Areas)
		if err != nil {
			return nil, err
		}

		if aeon.PhysAtkAccuracy != nil {
			aeon.PhysAtkAccuracy.ID, err = GetHashID(aeon.PhysAtkAccuracy)
			if err != nil {
				return nil, err
			}
		}

		aeons = append(aeons, *aeon)
	}

	return dedupeRows(aeons, hashes), nil
}

func completeAeons() error {
	for i := range rawJSON.aeons {
		aeon := &rawJSON.aeons[i]
		err := assignIDs(aeon.Weapon)
		if err != nil {
			return err
		}

		err = assignIDs(aeon.Armor)
		if err != nil {
			return err
		}

		L.Aeons[seedtypes.Key(aeon)] = *aeon
		L.AeonsID[aeon.ID] = *aeon
	}

	return nil
}

func getAeonAeonEquipment(a seedtypes.Aeon) ([]seedtypes.AeonEquipment, error) {
	return slices.Concat(a.Weapon, a.Armor), nil
}

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
