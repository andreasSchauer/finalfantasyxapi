package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

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
