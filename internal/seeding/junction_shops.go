package seeding

import (
	"context"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type ShopJunction struct {
	StdJunction
	ShopType database.ShopType
}

func (j ShopJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		j.ShopType,
	}
}

func (j ShopJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
}

func (l *Lookup) seedJuncShopsShopItems(qtx *database.Queries, ctx context.Context) error {
	const desc string = "shops + shop items"
	params := database.CreateShopsItemsJunctionBulkParams{
		DataHash:   make([]string, 0),
		ShopID:     make([]int32, 0),
		ShopItemID: make([]int32, 0),
		ShopType:   make([]database.ShopType, 0),
	}

	for _, shop := range l.json.shops {
		if shop.PreAirship != nil {
			for _, item := range shop.PreAirship.Items {
				j := ShopJunction{}
				j.ParentID = shop.ID
				j.ChildID = item.ID
				j.ShopType = shop.PreAirship.Type
				dataHash := generateJunctionHash(j, desc)

				params.DataHash = append(params.DataHash, dataHash)
				params.ShopID = append(params.ShopID, shop.ID)
				params.ShopItemID = append(params.ShopItemID, item.ID)
				params.ShopType = append(params.ShopType, shop.PreAirship.Type)
			}
		}

		if shop.PostAirship != nil {
			for _, item := range shop.PostAirship.Items {
				j := ShopJunction{}
				j.ParentID = shop.ID
				j.ChildID = item.ID
				j.ShopType = shop.PostAirship.Type
				dataHash := generateJunctionHash(j, desc)

				params.DataHash = append(params.DataHash, dataHash)
				params.ShopID = append(params.ShopID, shop.ID)
				params.ShopItemID = append(params.ShopItemID, item.ID)
				params.ShopType = append(params.ShopType, shop.PostAirship.Type)
			}
		}
	}

	return qtx.CreateShopsItemsJunctionBulk(ctx, params)
}

func (l *Lookup) getShopEquipment() []ShopEquipment {
	shopEquipment := []ShopEquipment{}

	for _, shop := range l.json.shops {
		if shop.PreAirship != nil {
			shopEquipment = append(shopEquipment, shop.PreAirship.Equipment...)
		}

		if shop.PostAirship != nil {
			shopEquipment = append(shopEquipment, shop.PostAirship.Equipment...)
		}
	}

	return shopEquipment
}

func (l *Lookup) getShopEquipmentAutoAbilities(se ShopEquipment) ([]AutoAbility, error) {
	return getResources(se.Abilities, l.AutoAbilities)
}

func (l *Lookup) seedJuncShopEquipmentAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "shop equipment + auto-abilities"
	jParams, err := processJunctions(l, desc, l.getShopEquipment(), l.getShopEquipmentAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateShopEquipmentAbilitiesJunctionBulk(ctx, database.CreateShopEquipmentAbilitiesJunctionBulkParams{
		DataHash:        jParams.DataHashes,
		ShopEquipmentID: jParams.ParentIDs,
		AutoAbilityID:   jParams.ChildIDs,
	})
}
