package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type TreasureEquipment struct {
	ID               int32
	TreasureID       int32
	EquipmentNameID  int32
	Name             string   `json:"name"`
	Abilities        []string `json:"abilities"`
	EmptySlotsAmount int32    `json:"empty_slots_amount"`
}

func (te TreasureEquipment) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", te),
		te.TreasureID,
		te.EquipmentNameID,
		te.EmptySlotsAmount,
	}
}

func (te TreasureEquipment) GetID() int32 {
	return te.ID
}

func (te *TreasureEquipment) SetID(id int32) {
	te.ID = id
}

func (te TreasureEquipment) Error() string {
	return fmt.Sprintf("treasure equipment with name: %s, empty slots: %d", te.Name, te.EmptySlotsAmount)
}

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

func (l *Lookup) getTreasureEquipment() []TreasureEquipment {
	treasureEquipment := []TreasureEquipment{}

	for _, list := range l.json.treasureLists {
		for _, treasure := range list.Treasures {
			if treasure.Equipment != nil {
				treasureEquipment = append(treasureEquipment, *treasure.Equipment)
			}
		}
	}

	return treasureEquipment
}

func (l *Lookup) getTreasureEquipmentAutoAbilities(te TreasureEquipment) ([]AutoAbility, error) {
	return getResources(te.Abilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncTreasureEquipmentAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "treasure equipment + auto-abilities"
	jParams, err := processJunctions(l, desc, l.getTreasureEquipment(), l.getTreasureEquipmentAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateTreasureEquipmentAbilitiesJunctionBulk(ctx, database.CreateTreasureEquipmentAbilitiesJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		TreasureEquipmentID: jParams.ParentIDs,
		AutoAbilityID:       jParams.ChildIDs,
	})
}
