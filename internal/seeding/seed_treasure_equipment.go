package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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

func (l *Lookup) seedFoundEquipment(qtx *database.Queries, foundEquipment TreasureEquipment) (TreasureEquipment, error) {
	var err error

	foundEquipment.EquipmentNameID, err = assignFK(foundEquipment.Name, l.EquipmentNames)
	if err != nil {
		return TreasureEquipment{}, h.NewErr(foundEquipment.Error(), err)
	}

	dbFoundEquipment, err := qtx.CreateTreasureEquipmentPiece(context.Background(), database.CreateTreasureEquipmentPieceParams{
		DataHash:         generateDataHash(foundEquipment),
		TreasureID:       foundEquipment.TreasureID,
		EquipmentNameID:  foundEquipment.EquipmentNameID,
		EmptySlotsAmount: foundEquipment.EmptySlotsAmount,
	})
	if err != nil {
		return TreasureEquipment{}, h.NewErr(foundEquipment.Error(), err, "couldn't create found equipment")
	}

	foundEquipment.ID = dbFoundEquipment.ID

	err = l.seedFoundEquipmentAbilities(qtx, foundEquipment)
	if err != nil {
		return TreasureEquipment{}, h.NewErr(foundEquipment.Error(), err)
	}

	return foundEquipment, nil
}

func (l *Lookup) seedFoundEquipmentAbilities(qtx *database.Queries, foundEquipment TreasureEquipment) error {
	for _, autoAbility := range foundEquipment.Abilities {
		junction, err := createJunction(foundEquipment, autoAbility, l.AutoAbilities)
		if err != nil {
			return h.NewErr(autoAbility, err)
		}

		err = qtx.CreateTreasureEquipmentAbilitiesJunction(context.Background(), database.CreateTreasureEquipmentAbilitiesJunctionParams{
			DataHash:            generateDataHash(junction),
			TreasureEquipmentID: junction.ParentID,
			AutoAbilityID:       junction.ChildID,
		})
		if err != nil {
			return h.NewErr(autoAbility, err, "couldn't junction auto-ability")
		}
	}

	return nil
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
