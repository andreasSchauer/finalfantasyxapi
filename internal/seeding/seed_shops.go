package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Shop struct {
	ID				int32
	Version      	*int32       	`json:"version"`
	LocationArea 	LocationArea 	`json:"location_area"`
	AreaID       	int32
	Notes        	*string 		`json:"notes"`
	Category     	string  		`json:"category"`
	PreAirship		*SubShop		`json:"pre_airship"`
	PostAirship		*SubShop		`json:"post_airship"`
}

func (s Shop) ToHashFields() []any {
	return []any{
		derefOrNil(s.Version),
		s.AreaID,
		derefOrNil(s.Notes),
		s.Category,
	}
}

func (s Shop) ToKeyFields() []any {
	return []any{
		createLookupKey(s.LocationArea),
		derefOrNil(s.Version),
	}
}

func (s Shop) GetID() int32 {
	return s.ID
}

type SubShop struct {
	Items		[]ShopItem		`json:"items"`
	Equipment	[]ShopEquipment	`json:"equipment"`
	Type		database.ShopType
}


type ShopItem struct {
	ID		int32
	ItemID	int32
	Name	string	`json:"name"`
	Price	int32	`json:"price"`
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


type ShopEquipment struct {
	ID					int32
	FoundEquipment
	Price				int32	`json:"price"`
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


type ShopJunction struct {
	Junction
	ShopType	database.ShopType
}

func (sj ShopJunction) ToHashFields() []any {
	return []any{
		sj.ParentID,
		sj.ChildID,
		sj.ShopType,
	}
}


func (l *lookup) seedShops(db *database.Queries, dbConn *sql.DB) error {
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
			shop.AreaID, err = assignFK(locationArea, l.getArea)
			if err != nil {
				return fmt.Errorf("shops: %v", err)
			}

			dbShop, err := qtx.CreateShop(context.Background(), database.CreateShopParams{
				DataHash: generateDataHash(shop),
				Version:  getNullInt32(shop.Version),
				AreaID:   shop.AreaID,
				Notes:    getNullString(shop.Notes),
				Category: database.ShopCategory(shop.Category),
			})
			if err != nil {
				return fmt.Errorf("couldn't create shop: %s - shop version: %d: %v", createLookupKey(locationArea), derefOrNil(shop.Version), err)
			}
			shop.ID = dbShop.ID
			key := createLookupKey(shop)
			l.shops[key] = shop
		}
		return nil
	})
}



func (l *lookup) seedShopsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/shops.json"

	var shops []Shop
	err := loadJSONFile(string(srcPath), &shops)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonShop := range shops {
			key := createLookupKey(jsonShop)
			shop, err := l.getShop(key)
			if err != nil {
				return err
			}

			if shop.PreAirship != nil {
				shop.PreAirship.Type = database.ShopTypePreAirship
				err := l.seedSubShop(qtx, shop, shop.PreAirship)
				if err != nil {
					return err
				}
			}

			if shop.PostAirship != nil {
				shop.PostAirship.Type = database.ShopTypePostAirship
				err := l.seedSubShop(qtx, shop, shop.PostAirship)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}


func (l *lookup) seedSubShop (qtx *database.Queries, shop Shop, subShop *SubShop) error {
	err := l.seedShopItems(qtx, shop, subShop)
	if err != nil {
		return err
	}

	err = l.seedShopEquipmentPieces(qtx, shop, subShop)
	if err != nil {
		return err
	}

	return nil
}


func (l *lookup) seedShopItems (qtx *database.Queries, shop Shop, subShop *SubShop) error {
	for _, shopItem := range subShop.Items {
		junction, err := createJunctionSeed(qtx, shop, shopItem, l.seedShopItem)
		if err != nil {
			return fmt.Errorf("couldn't create junction for shop: %s: %v", createLookupKey(shop), err)
		}

		shopJunction := ShopJunction{
			Junction: junction,
			ShopType: subShop.Type,
		}

		err = qtx.CreateShopsItemsJunction(context.Background(), database.CreateShopsItemsJunctionParams{
			DataHash: 	generateDataHash(shopJunction),
			ShopID: 	shopJunction.ParentID,
			ShopItemID: shopJunction.ChildID,
			ShopType: 	shopJunction.ShopType,
		})
		if err != nil {
			return fmt.Errorf("couldn't seed junction for shop: %s: %v", createLookupKey(shop), err)
		}
	}

	return nil
}


func (l *lookup) seedShopItem (qtx *database.Queries, shopItem ShopItem) (ShopItem, error) {
	var err error

	shopItem.ItemID, err = assignFK(shopItem.Name, l.getItem)
	if err != nil {
		return ShopItem{}, err
	}

	dbShopItem, err := qtx.CreateShopItem(context.Background(), database.CreateShopItemParams{
		DataHash: 	generateDataHash(shopItem),
		ItemID: 	shopItem.ItemID,
		Price: 	shopItem.Price,
	})
	if err != nil {
		return ShopItem{}, fmt.Errorf("couldn't create shop item %s: %v", createLookupKey(shopItem), err)
	}

	shopItem.ID = dbShopItem.ID

	return shopItem, nil
}


func (l *lookup) seedShopEquipmentPieces (qtx *database.Queries, shop Shop, subShop *SubShop) error {
	for _, shopEquipment := range subShop.Equipment {
		junction, err := createJunctionSeed(qtx, shop, shopEquipment, l.seedShopEquipment)
		if err != nil {
			return fmt.Errorf("couldn't create junction for shop: %s: %v", createLookupKey(shop), err)
		}

		shopJunction := ShopJunction{
			Junction: junction,
			ShopType: subShop.Type,
		}

		err = qtx.CreateShopsEquipmentJunction(context.Background(), database.CreateShopsEquipmentJunctionParams{
			DataHash: 			generateDataHash(shopJunction),
			ShopID: 			shopJunction.ParentID,
			ShopEquipmentID: 	shopJunction.ChildID,
			ShopType: 			shopJunction.ShopType,
		})
		if err != nil {
			return fmt.Errorf("couldn't seed junction for shop: %s: %v", createLookupKey(shop), err)
		}
	}

	return nil
}


func (l *lookup) seedShopEquipment (qtx *database.Queries, shopEquipment ShopEquipment) (ShopEquipment, error) {
	var err error

	shopEquipment.FoundEquipment, err = seedObjAssignID(qtx, shopEquipment.FoundEquipment, l.seedFoundEquipment)
	if err != nil {
		return ShopEquipment{}, err
	}

	dbShopEquipment, err := qtx.CreateShopEquipmentPiece(context.Background(), database.CreateShopEquipmentPieceParams{
		DataHash: 			generateDataHash(shopEquipment),
		FoundEquipmentID: 	shopEquipment.FoundEquipment.ID,
		Price: 				shopEquipment.Price,
	})
	if err != nil {
		return ShopEquipment{}, fmt.Errorf("couldn't create shop item %s: %v", createLookupKey(shopEquipment), err)
	}

	shopEquipment.ID = dbShopEquipment.ID

	return shopEquipment, nil
}