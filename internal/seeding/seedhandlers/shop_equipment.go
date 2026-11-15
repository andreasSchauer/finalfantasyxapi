package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop6ShopEquipment(qtx *database.Queries, ctx context.Context) error {
	equipment, err := extractShopEquipment()
	if err != nil {
		return err
	}

	params := database.CreateShopEquipmentPieceBulkParams{
		DataHash:         make([]string, len(equipment)),
		ShopID:           make([]int32, len(equipment)),
		EquipmentNameID:  make([]int32, len(equipment)),
		ShopType:         make([]database.ShopType, len(equipment)),
		EmptySlotsAmount: make([]int32, len(equipment)),
		Price:            make([]int32, len(equipment)),
	}

	for i, se := range equipment {
		params.DataHash[i] = sh.GenerateDataHash(se)
		params.ShopID[i] = se.ShopID
		params.EquipmentNameID[i] = se.EquipmentNameID
		params.ShopType[i] = se.ShopType
		params.EmptySlotsAmount[i] = se.EmptySlotsAmount
		params.Price[i] = se.Price
	}

	dbRows, err := qtx.CreateShopEquipmentPieceBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create shop equipment: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractShopEquipment() ([]seedtypes.ShopEquipment, error) {
	shopEquipment := []seedtypes.ShopEquipment{}
	var err error

	for i := range sh.RawJSON.Shops {
		shop := &sh.RawJSON.Shops[i]

		if shop.PreAirship != nil {
			for j := range shop.PreAirship.Equipment {
				equipment := &shop.PreAirship.Equipment[j]

				equipment.ShopID = shop.ID
				equipment.ShopType = database.ShopTypePreAirship

				equipment.EquipmentNameID, err = sh.AssignFK(equipment.Name, lookup.L.EquipmentNames)
				if err != nil {
					return nil, err
				}

				shopEquipment = append(shopEquipment, *equipment)
			}
		}

		if shop.PostAirship != nil {
			for j := range shop.PostAirship.Equipment {
				equipment := &shop.PostAirship.Equipment[j]

				equipment.ShopID = shop.ID
				equipment.ShopType = database.ShopTypePostAirship

				equipment.EquipmentNameID, err = sh.AssignFK(equipment.Name, lookup.L.EquipmentNames)
				if err != nil {
					return nil, err
				}

				shopEquipment = append(shopEquipment, *equipment)
			}
		}
	}

	return sh.DedupeRows(shopEquipment, sh.Hashes), nil
}
