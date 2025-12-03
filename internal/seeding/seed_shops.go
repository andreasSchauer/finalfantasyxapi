package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Shop struct {
	ID           int32
	Version      *int32       `json:"version"`
	LocationArea LocationArea `json:"location_area"`
	AreaID       int32
	Notes        *string  `json:"notes"`
	Category     string   `json:"category"`
	PreAirship   *SubShop `json:"pre_airship"`
	PostAirship  *SubShop `json:"post_airship"`
}

func (s Shop) ToHashFields() []any {
	return []any{
		h.DerefOrNil(s.Version),
		s.AreaID,
		h.DerefOrNil(s.Notes),
		s.Category,
	}
}

func (s Shop) ToKeyFields() []any {
	return []any{
		CreateLookupKey(s.LocationArea),
		h.DerefOrNil(s.Version),
	}
}

func (s Shop) GetID() int32 {
	return s.ID
}

func (s Shop) Error() string {
	return fmt.Sprintf("shop %s, %v", s.LocationArea, h.DerefOrNil(s.Version))
}

type SubShop struct {
	Items     []ShopItem      `json:"items"`
	Equipment []ShopEquipment `json:"equipment"`
	Type      database.ShopType
}

func (s SubShop) Error() string {
	return fmt.Sprintf("subshop type: %s", s.Type)
}

type ShopItem struct {
	ID     int32
	ItemID int32
	Name   string `json:"name"`
	Price  int32  `json:"price"`
}

func (s ShopItem) ToHashFields() []any {
	return []any{
		s.ItemID,
		s.Price,
	}
}

func (s ShopItem) ToKeyFields() []any {
	return []any{
		s.Name,
		s.Price,
	}
}

func (s ShopItem) GetID() int32 {
	return s.ID
}

func (s ShopItem) Error() string {
	return fmt.Sprintf("shop item %s, price %d", s.Name, s.Price)
}

type ShopEquipment struct {
	ID int32
	FoundEquipment
	Price int32 `json:"price"`
}

func (s ShopEquipment) ToHashFields() []any {
	return []any{
		s.FoundEquipment.ID,
		s.Price,
	}
}

func (s ShopEquipment) ToKeyFields() []any {
	return []any{
		s.Name,
		s.EmptySlotsAmount,
		s.Price,
	}
}

func (s ShopEquipment) GetID() int32 {
	return s.ID
}

func (s ShopEquipment) Error() string {
	return fmt.Sprintf("shop equipment %s, empty slots %d, price %d", s.Name, s.EmptySlotsAmount, s.Price)
}

type ShopJunction struct {
	Junction
	ShopType database.ShopType
}

func (sj ShopJunction) ToHashFields() []any {
	return []any{
		sj.ParentID,
		sj.ChildID,
		sj.ShopType,
	}
}

func (l *Lookup) seedShops(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/shops.json"

	var shops []Shop
	err := loadJSONFile(string(srcPath), &shops)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, shop := range shops {
			var err error

			locationArea := shop.LocationArea
			shop.AreaID, err = assignFK(locationArea, l.Areas)
			if err != nil {
				return h.GetErr(shop.Error(), err)
			}

			dbShop, err := qtx.CreateShop(context.Background(), database.CreateShopParams{
				DataHash: generateDataHash(shop),
				Version:  h.GetNullInt32(shop.Version),
				AreaID:   shop.AreaID,
				Notes:    h.GetNullString(shop.Notes),
				Category: database.ShopCategory(shop.Category),
			})
			if err != nil {
				return h.GetErr(shop.Error(), err, "couldn't create shop")
			}
			shop.ID = dbShop.ID
			key := CreateLookupKey(shop)
			l.Shops[key] = shop
		}
		return nil
	})
}

func (l *Lookup) seedShopsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/shops.json"

	var shops []Shop
	err := loadJSONFile(string(srcPath), &shops)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonShop := range shops {
			key := CreateLookupKey(jsonShop)
			shop, err := GetResource(key, l.Shops)
			if err != nil {
				return err
			}

			if shop.PreAirship != nil {
				shop.PreAirship.Type = database.ShopTypePreAirship
				err := l.seedSubShop(qtx, shop, shop.PreAirship)
				if err != nil {
					return h.GetErr(shop.Error(), err)
				}
			}

			if shop.PostAirship != nil {
				shop.PostAirship.Type = database.ShopTypePostAirship
				err := l.seedSubShop(qtx, shop, shop.PostAirship)
				if err != nil {
					return h.GetErr(shop.Error(), err)
				}
			}
		}
		return nil
	})
}

func (l *Lookup) seedSubShop(qtx *database.Queries, shop Shop, subShop *SubShop) error {
	err := l.seedShopItems(qtx, shop, subShop)
	if err != nil {
		return h.GetErr(subShop.Error(), err)
	}

	err = l.seedShopEquipmentPieces(qtx, shop, subShop)
	if err != nil {
		return h.GetErr(subShop.Error(), err)
	}

	return nil
}

func (l *Lookup) seedShopItems(qtx *database.Queries, shop Shop, subShop *SubShop) error {
	for _, shopItem := range subShop.Items {
		junction, err := createJunctionSeed(qtx, shop, shopItem, l.seedShopItem)
		if err != nil {
			return err
		}

		shopJunction := ShopJunction{
			Junction: junction,
			ShopType: subShop.Type,
		}

		err = qtx.CreateShopsItemsJunction(context.Background(), database.CreateShopsItemsJunctionParams{
			DataHash:   generateDataHash(shopJunction),
			ShopID:     shopJunction.ParentID,
			ShopItemID: shopJunction.ChildID,
			ShopType:   shopJunction.ShopType,
		})
		if err != nil {
			return h.GetErr(shopItem.Error(), err, "couldn't junction shop item")
		}
	}

	return nil
}

func (l *Lookup) seedShopItem(qtx *database.Queries, shopItem ShopItem) (ShopItem, error) {
	var err error

	shopItem.ItemID, err = assignFK(shopItem.Name, l.Items)
	if err != nil {
		return ShopItem{}, h.GetErr(shopItem.Error(), err)
	}

	dbShopItem, err := qtx.CreateShopItem(context.Background(), database.CreateShopItemParams{
		DataHash: generateDataHash(shopItem),
		ItemID:   shopItem.ItemID,
		Price:    shopItem.Price,
	})
	if err != nil {
		return ShopItem{}, h.GetErr(shopItem.Error(), err, "couldn't create shop item")
	}

	shopItem.ID = dbShopItem.ID

	return shopItem, nil
}

func (l *Lookup) seedShopEquipmentPieces(qtx *database.Queries, shop Shop, subShop *SubShop) error {
	for _, shopEquipment := range subShop.Equipment {
		junction, err := createJunctionSeed(qtx, shop, shopEquipment, l.seedShopEquipment)
		if err != nil {
			return err
		}

		shopJunction := ShopJunction{
			Junction: junction,
			ShopType: subShop.Type,
		}

		err = qtx.CreateShopsEquipmentJunction(context.Background(), database.CreateShopsEquipmentJunctionParams{
			DataHash:        generateDataHash(shopJunction),
			ShopID:          shopJunction.ParentID,
			ShopEquipmentID: shopJunction.ChildID,
			ShopType:        shopJunction.ShopType,
		})
		if err != nil {
			return h.GetErr(shopEquipment.Error(), err, "couldn't junction shop equipment")
		}
	}

	return nil
}

func (l *Lookup) seedShopEquipment(qtx *database.Queries, shopEquipment ShopEquipment) (ShopEquipment, error) {
	var err error

	shopEquipment.FoundEquipment, err = seedObjAssignID(qtx, shopEquipment.FoundEquipment, l.seedFoundEquipment)
	if err != nil {
		return ShopEquipment{}, h.GetErr(shopEquipment.Error(), err)
	}

	dbShopEquipment, err := qtx.CreateShopEquipmentPiece(context.Background(), database.CreateShopEquipmentPieceParams{
		DataHash:         generateDataHash(shopEquipment),
		FoundEquipmentID: shopEquipment.FoundEquipment.ID,
		Price:            shopEquipment.Price,
	})
	if err != nil {
		return ShopEquipment{}, h.GetErr(shopEquipment.Error(), err, "couldn't create shop equipment")
	}

	shopEquipment.ID = dbShopEquipment.ID

	return shopEquipment, nil
}
