package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"slices"

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
	Availability string   `json:"availability"`
	PreAirship   *SubShop `json:"pre_airship"`
	PostAirship  *SubShop `json:"post_airship"`
}

func (s Shop) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		h.DerefOrNil(s.Version),
		s.AreaID,
		h.DerefOrNil(s.Notes),
		s.Category,
		s.Availability,
	}
}

func (s Shop) ToKeyFields() []any {
	return []any{
		Key(s.LocationArea),
		h.DerefOrNil(s.Version),
	}
}

func (s Shop) GetID() int32 {
	return s.ID
}

func (s Shop) Error() string {
	return fmt.Sprintf("shop %s, %v", s.LocationArea, h.PtrToString(s.Version))
}

func (s Shop) GetResParamsUnnamed() h.ResParamsUnnamed {
	return h.ResParamsUnnamed{
		ID: s.ID,
	}
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
		fmt.Sprintf("%T", s),
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

func (s *ShopItem) SetID(id int32) {
	s.ID = id
}

func (s ShopItem) Error() string {
	return fmt.Sprintf("shop item %s, price %d", s.Name, s.Price)
}

type ShopEquipment struct {
	ID       int32
	ShopID   int32
	ShopType database.ShopType
	TreasureEquipment
	Price int32 `json:"price"`
}

func (s ShopEquipment) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.ShopID,
		s.EquipmentNameID,
		s.ShopType,
		s.EmptySlotsAmount,
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

func (s *ShopEquipment) SetID(id int32) {
	s.ID = id
}

func (s ShopEquipment) Error() string {
	return fmt.Sprintf("shop equipment %s, empty slots %d, price %d", s.Name, s.EmptySlotsAmount, s.Price)
}

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

func (l *Lookup) seedShops(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/shops.json"

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
				return h.NewErr(shop.Error(), err)
			}

			dbShop, err := qtx.CreateShop(context.Background(), database.CreateShopParams{
				DataHash:     generateDataHash(shop),
				Version:      h.GetNullInt32(shop.Version),
				AreaID:       shop.AreaID,
				Notes:        h.GetNullString(shop.Notes),
				Category:     database.ShopCategory(shop.Category),
				Availability: database.AvailabilityType(shop.Availability),
			})
			if err != nil {
				return h.NewErr(shop.Error(), err, "couldn't create shop")
			}
			shop.ID = dbShop.ID
			key := Key(shop)
			l.Shops[key] = shop
			l.ShopsID[shop.ID] = shop
		}
		return nil
	})
}

func (l *Lookup) seedShopsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/shops.json"

	var shops []Shop
	err := loadJSONFile(string(srcPath), &shops)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonShop := range shops {
			key := Key(jsonShop)
			shop, err := GetResource(key, l.Shops)
			if err != nil {
				return err
			}

			if shop.PreAirship != nil {
				shop.PreAirship.Type = database.ShopTypePreAirship
				err := l.seedSubShop(qtx, shop, shop.PreAirship)
				if err != nil {
					return h.NewErr(shop.Error(), err)
				}
			}

			if shop.PostAirship != nil {
				shop.PostAirship.Type = database.ShopTypePostAirship
				err := l.seedSubShop(qtx, shop, shop.PostAirship)
				if err != nil {
					return h.NewErr(shop.Error(), err)
				}
			}
		}
		return nil
	})
}

func (l *Lookup) seedSubShop(qtx *database.Queries, shop Shop, subShop *SubShop) error {
	err := l.seedShopItems(qtx, shop, subShop)
	if err != nil {
		return h.NewErr(subShop.Error(), err)
	}

	err = l.seedShopEquipmentPieces(qtx, shop, subShop)
	if err != nil {
		return h.NewErr(subShop.Error(), err)
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
			StdJunction: junction,
			ShopType:    subShop.Type,
		}

		err = qtx.CreateShopsItemsJunction(context.Background(), database.CreateShopsItemsJunctionParams{
			DataHash:   generateDataHash(shopJunction),
			ShopID:     shopJunction.ParentID,
			ShopItemID: shopJunction.ChildID,
			ShopType:   shopJunction.ShopType,
		})
		if err != nil {
			return h.NewErr(shopItem.Error(), err, "couldn't junction shop item")
		}
	}

	return nil
}

func (l *Lookup) seedShopItem(qtx *database.Queries, shopItem ShopItem) (ShopItem, error) {
	var err error

	shopItem.ItemID, err = assignFK(shopItem.Name, l.Items)
	if err != nil {
		return ShopItem{}, h.NewErr(shopItem.Error(), err)
	}

	dbShopItem, err := qtx.CreateShopItem(context.Background(), database.CreateShopItemParams{
		DataHash: generateDataHash(shopItem),
		ItemID:   shopItem.ItemID,
		Price:    shopItem.Price,
	})
	if err != nil {
		return ShopItem{}, h.NewErr(shopItem.Error(), err, "couldn't create shop item")
	}

	shopItem.ID = dbShopItem.ID

	return shopItem, nil
}

func (l *Lookup) seedShopEquipmentPieces(qtx *database.Queries, shop Shop, subShop *SubShop) error {
	for _, shopEquipment := range subShop.Equipment {
		var err error
		shopEquipment.ShopID = shop.ID
		shopEquipment.ShopType = subShop.Type

		shopEquipment.EquipmentNameID, err = assignFK(shopEquipment.Name, l.EquipmentNames)
		if err != nil {
			return h.NewErr(shopEquipment.Error(), err)
		}

		dbShopEquipment, err := qtx.CreateShopEquipmentPiece(context.Background(), database.CreateShopEquipmentPieceParams{
			DataHash:         generateDataHash(shopEquipment),
			ShopID:           shopEquipment.ShopID,
			EquipmentNameID:  shopEquipment.EquipmentNameID,
			ShopType:         shopEquipment.ShopType,
			EmptySlotsAmount: shopEquipment.EmptySlotsAmount,
			Price:            shopEquipment.Price,
		})
		if err != nil {
			return h.NewErr(shopEquipment.Error(), err, "couldn't create shop equipment")
		}

		shopEquipment.ID = dbShopEquipment.ID

		err = l.seedShopEquipmentAbilities(qtx, shopEquipment)
		if err != nil {
			return h.NewErr(shopEquipment.Error(), err)
		}
	}

	return nil
}

func (l *Lookup) seedShopEquipmentAbilities(qtx *database.Queries, shopEquipment ShopEquipment) error {
	for _, autoAbility := range shopEquipment.Abilities {
		junction, err := createJunction(shopEquipment, autoAbility, l.AutoAbilities)
		if err != nil {
			return h.NewErr(autoAbility, err)
		}

		err = qtx.CreateShopEquipmentAbilitiesJunction(context.Background(), database.CreateShopEquipmentAbilitiesJunctionParams{
			DataHash:        generateDataHash(junction),
			ShopEquipmentID: junction.ParentID,
			AutoAbilityID:   junction.ChildID,
		})
		if err != nil {
			return h.NewErr(autoAbility, err, "couldn't junction auto-ability")
		}
	}

	return nil
}

func (l *Lookup) loop4SeedShops(qtx *database.Queries, ctx context.Context) error {
	shops, err := l.extractShops()
	if err != nil {
		return err
	}

	params := database.CreateShopBulkParams{
		DataHash:     make([]string, len(shops)),
		Version:      make([]sql.NullInt32, len(shops)),
		AreaID:       make([]int32, len(shops)),
		Notes:        make([]sql.NullString, len(shops)),
		Category:     make([]database.ShopCategory, len(shops)),
		Availability: make([]database.AvailabilityType, len(shops)),
	}

	for i, s := range shops {
		params.DataHash[i] = generateDataHash(s)
		params.Version[i] = h.GetNullInt32(s.Version)
		params.AreaID[i] = s.AreaID
		params.Notes[i] = h.GetNullString(s.Notes)
		params.Category[i] = database.ShopCategory(s.Category)
		params.Availability[i] = database.AvailabilityType(s.Availability)
	}

	dbRows, err := qtx.CreateShopBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create shops: %v", err)
	}

	for i, row := range dbRows {
		shops[i].ID = row.ID
		l.json.shops[i].ID = row.ID
		key := Key(shops[i])
		l.Shops[key] = shops[i]
		l.ShopsID[row.ID] = shops[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractShops() ([]Shop, error) {
	shops := []Shop{}
	var err error

	for i := range l.json.shops {
		shop := &l.json.shops[i]

		if shop.PreAirship != nil {
			shop.PreAirship.Type = database.ShopTypePreAirship
		}

		if shop.PostAirship != nil {
			shop.PostAirship.Type = database.ShopTypePostAirship
		}

		shop.AreaID, err = assignFK(shop.LocationArea, l.Areas)
		if err != nil {
			return nil, err
		}

		shops = append(shops, *shop)
	}

	return dedupeRows(shops, l.Hashes), nil
}

func (l *Lookup) completeShops() error {
	for i := range l.json.shops {
		shop := &l.json.shops[i]

		err := l.completeSubShop(shop.PreAirship)
		if err != nil {
			return err
		}

		err = l.completeSubShop(shop.PostAirship)
		if err != nil {
			return err
		}

		l.Shops[Key(*shop)] = *shop
		l.ShopsID[shop.ID] = *shop
	}

	return nil
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

func (l *Lookup) completeSubShop(subShop *SubShop) error {
	if subShop == nil {
		return nil
	}

	err := assignIDs(l, subShop.Items)
	if err != nil {
		return err
	}

	err = assignIDs(l, subShop.Equipment)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lookup) loop3SeedShopItems(qtx *database.Queries, ctx context.Context) error {
	items, err := l.extractShopItems()
	if err != nil {
		return err
	}

	params := database.CreateShopItemBulkParams{
		DataHash: make([]string, len(items)),
		ItemID:   make([]int32, len(items)),
		Price:    make([]int32, len(items)),
	}

	for i, si := range items {
		params.DataHash[i] = generateDataHash(si)
		params.ItemID[i] = si.ItemID
		params.Price[i] = si.Price
	}

	dbRows, err := qtx.CreateShopItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create shop items: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractShopItems() ([]ShopItem, error) {
	items := []ShopItem{}
	var err error

	for i := range l.json.shops {
		shop := &l.json.shops[i]

		if shop.PreAirship != nil {
			for j := range shop.PreAirship.Items {
				item := &shop.PreAirship.Items[j]

				item.ItemID, err = assignFK(item.Name, l.Items)
				if err != nil {
					return nil, err
				}

				items = append(items, *item)
			}
		}

		if shop.PostAirship != nil {
			for j := range shop.PostAirship.Items {
				item := &shop.PostAirship.Items[j]

				item.ItemID, err = assignFK(item.Name, l.Items)
				if err != nil {
					return nil, err
				}

				items = append(items, *item)
			}
		}
	}

	return dedupeRows(items, l.Hashes), nil
}

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