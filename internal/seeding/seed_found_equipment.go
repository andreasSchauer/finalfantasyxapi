package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type FoundEquipment struct {
	ID					int32
	EquipmentNameID		int32
	Name				string		`json:"name"`
	Abilities			[]string	`json:"abilities"`
	EmptySlotsAmount	int32		`json:"empty_slots_amount"`
}


func (f FoundEquipment) ToHashFields() []any {
	return []any{
		f.EquipmentNameID,
		f.EmptySlotsAmount,
	}
}


func (f FoundEquipment) GetID() int32 {
	return f.ID
}

func (f FoundEquipment) Error() string {
	return fmt.Sprintf("found equipment with name: %s, empty slots: %d", f.Name, f.EmptySlotsAmount)
}


func (l *lookup) seedFoundEquipment(qtx *database.Queries, foundEquipment FoundEquipment) (FoundEquipment, error) {
	var err error

	foundEquipment.EquipmentNameID, err = assignFK(foundEquipment.Name, l.getEquipmentName)
	if err != nil {
		return FoundEquipment{}, err
	}

	dbFoundEquipment, err := qtx.CreateFoundEquipmentPiece(context.Background(), database.CreateFoundEquipmentPieceParams{
		DataHash: generateDataHash(foundEquipment),
		EquipmentNameID: foundEquipment.EquipmentNameID,
		EmptySlotsAmount: foundEquipment.EmptySlotsAmount,
	})
	if err != nil {
		return FoundEquipment{}, fmt.Errorf("couldn't create found equipment %s: %v", foundEquipment.Name, err)
	}

	foundEquipment.ID = dbFoundEquipment.ID

	err = l.seedFoundEquipmentAbilities(qtx, foundEquipment)
	if err != nil {
		return FoundEquipment{}, fmt.Errorf("found equipment: %s: %v", foundEquipment.Name, err)
	}

	return foundEquipment, nil
}


func (l *lookup) seedFoundEquipmentAbilities(qtx *database.Queries, foundEquipment FoundEquipment) error {
	for _, autoAbility := range foundEquipment.Abilities {
		junction, err := createJunction(foundEquipment, autoAbility, l.getAutoAbility)
		if err != nil {
			return fmt.Errorf("couldn't create junction with auto ability %s: %v", autoAbility, err)
		}

		err = qtx.CreateFoundEquipmentAbilitiesJunction(context.Background(), database.CreateFoundEquipmentAbilitiesJunctionParams{
			DataHash: generateDataHash(junction),
			FoundEquipmentID: junction.ParentID,
			AutoAbilityID: junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't seed junction with auto ability %s: %v", autoAbility, err)
		}
	}

	return nil
}