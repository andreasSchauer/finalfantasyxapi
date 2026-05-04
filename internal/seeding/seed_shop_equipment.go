package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop6SeedShopEquipment(qtx *database.Queries, ctx context.Context) error {
	equipment, err := l.extractShopEquipment()
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
		params.DataHash[i] = generateDataHash(se)
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractShopEquipment() ([]ShopEquipment, error) {
	shopEquipment := []ShopEquipment{}
	var err error

	for i := range l.json.shops {
		shop := &l.json.shops[i]

		if shop.PreAirship != nil {
			for j := range shop.PreAirship.Equipment {
				equipment := &shop.PreAirship.Equipment[j]

				equipment.ShopID = shop.ID
				equipment.ShopType = database.ShopTypePreAirship

				equipment.EquipmentNameID, err = assignFK(equipment.Name, l.EquipmentNames)
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

				equipment.EquipmentNameID, err = assignFK(equipment.Name, l.EquipmentNames)
				if err != nil {
					return nil, err
				}

				shopEquipment = append(shopEquipment, *equipment)
			}
		}
	}

	return dedupeRows(shopEquipment, l.Hashes), nil
}
