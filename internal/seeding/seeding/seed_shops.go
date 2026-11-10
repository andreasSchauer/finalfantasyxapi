package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop4SeedShops(qtx *database.Queries, ctx context.Context) error {
	shops, err := extractShops()
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
		rawJSON.shops[i].ID = row.ID
		lookup.L.Shops[seedtypes.Key(shops[i])] = shops[i]
		lookup.L.ShopsID[row.ID] = shops[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractShops() ([]seedtypes.Shop, error) {
	shops := []seedtypes.Shop{}
	var err error

	for i := range rawJSON.shops {
		shop := &rawJSON.shops[i]

		if shop.PreAirship != nil {
			shop.PreAirship.Type = database.ShopTypePreAirship
		}

		if shop.PostAirship != nil {
			shop.PostAirship.Type = database.ShopTypePostAirship
		}

		shop.AreaID, err = assignFK(shop.LocationArea, lookup.L.Areas)
		if err != nil {
			return nil, err
		}

		shops = append(shops, *shop)
	}

	return dedupeRows(shops, hashes), nil
}

func completeShops() error {
	for i := range rawJSON.shops {
		shop := &rawJSON.shops[i]

		err := completeSubShop(shop.PreAirship)
		if err != nil {
			return err
		}

		err = completeSubShop(shop.PostAirship)
		if err != nil {
			return err
		}

		lookup.L.Shops[seedtypes.Key(*shop)] = *shop
		lookup.L.ShopsID[shop.ID] = *shop
	}

	return nil
}

func completeSubShop(subShop *seedtypes.SubShop) error {
	if subShop == nil {
		return nil
	}

	err := assignIDs(subShop.Items)
	if err != nil {
		return err
	}

	err = assignIDs(subShop.Equipment)
	if err != nil {
		return err
	}

	return nil
}
