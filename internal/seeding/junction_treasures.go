package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (l *Lookup) getTreasures() []seedtypes.Treasure {
	treasures := []seedtypes.Treasure{}

	for _, list := range l.json.treasureLists {
		treasures = append(treasures, list.Treasures...)
	}

	return treasures
}

func (l *Lookup) getTreasureItems(t seedtypes.Treasure) ([]seedtypes.ItemAmount, error) {
	return t.Items, nil
}

func (l *Lookup) seedJuncTreasuresItems(qtx *database.Queries, ctx context.Context) error {
	const desc string = "treasures + items"
	jParams, err := processJunctions(l, desc, l.getTreasures(), l.getTreasureItems)
	if err != nil {
		return err
	}

	return qtx.CreateTreasuresItemsJunctionBulk(ctx, database.CreateTreasuresItemsJunctionBulkParams{
		DataHash:     jParams.DataHashes,
		TreasureID:   jParams.ParentIDs,
		ItemAmountID: jParams.ChildIDs,
	})
}

func (l *Lookup) getTreasureEquipment() []seedtypes.TreasureEquipment {
	treasureEquipment := []seedtypes.TreasureEquipment{}

	for _, list := range l.json.treasureLists {
		for _, treasure := range list.Treasures {
			if treasure.Equipment != nil {
				treasureEquipment = append(treasureEquipment, *treasure.Equipment)
			}
		}
	}

	return treasureEquipment
}

func (l *Lookup) getTreasureEquipmentAutoAbilities(te seedtypes.TreasureEquipment) ([]seedtypes.AutoAbility, error) {
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
