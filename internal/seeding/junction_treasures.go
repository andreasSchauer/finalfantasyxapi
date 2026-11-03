package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getTreasures() []seedtypes.Treasure {
	treasures := []seedtypes.Treasure{}

	for _, list := range rawJSON.treasureLists {
		treasures = append(treasures, list.Treasures...)
	}

	return treasures
}

func getTreasureItems(t seedtypes.Treasure) ([]seedtypes.ItemAmount, error) {
	return t.Items, nil
}

func seedJuncTreasuresItems(qtx *database.Queries, ctx context.Context) error {
	const desc string = "treasures + items"
	jParams, err := processJunctions(desc, getTreasures(), getTreasureItems)
	if err != nil {
		return err
	}

	return qtx.CreateTreasuresItemsJunctionBulk(ctx, database.CreateTreasuresItemsJunctionBulkParams{
		DataHash:     jParams.DataHashes,
		TreasureID:   jParams.ParentIDs,
		ItemAmountID: jParams.ChildIDs,
	})
}

func getTreasureEquipment() []seedtypes.TreasureEquipment {
	treasureEquipment := []seedtypes.TreasureEquipment{}

	for _, list := range rawJSON.treasureLists {
		for _, treasure := range list.Treasures {
			if treasure.Equipment != nil {
				treasureEquipment = append(treasureEquipment, *treasure.Equipment)
			}
		}
	}

	return treasureEquipment
}

func getTreasureEquipmentAutoAbilities(te seedtypes.TreasureEquipment) ([]seedtypes.AutoAbility, error) {
	return GetResources(te.Abilities, L.AutoAbilities)
}

func seedJuncTreasureEquipmentAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "treasure equipment + auto-abilities"
	jParams, err := processJunctions(desc, getTreasureEquipment(), getTreasureEquipmentAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateTreasureEquipmentAbilitiesJunctionBulk(ctx, database.CreateTreasureEquipmentAbilitiesJunctionBulkParams{
		DataHash:            jParams.DataHashes,
		TreasureEquipmentID: jParams.ParentIDs,
		AutoAbilityID:       jParams.ChildIDs,
	})
}
