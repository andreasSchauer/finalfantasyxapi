package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Shop struct {
	//id 		int32
	//dataHash	string
	Version			*int32			`json:"version"`
	LocationArea	LocationArea 	`json:"location_area"`
	AreaID			int32
	Notes			*string 		`json:"notes"`
	Category		string			`json:"category"`
}


func(s Shop) ToHashFields() []any {
	return []any{
		derefOrNil(s.Version),
		s.AreaID,
		derefOrNil(s.Notes),
		s.Category,
	}
}


func seedShops(qtx *database.Queries, lookup map[string]int32) error {
	const srcPath = "./data/shops.json"

	var shops []Shop
	err := loadJSONFile(string(srcPath), &shops)
	if err != nil {
		return err
	}


	for _, shop := range shops {
		locationArea := shop.LocationArea
		locationAreaID, err := getAreaID(locationArea, lookup)
		if err != nil {
			return fmt.Errorf("shops: %v", err)
		}

		shop.AreaID = locationAreaID

		err = qtx.CreateShop(context.Background(), database.CreateShopParams{
			DataHash: 		generateDataHash(shop),
			Version: 		getNullInt32(shop.Version),
			AreaID: 		shop.AreaID,	
			Notes: 			getNullString(shop.Notes),
			Category: 		database.ShopCategory(shop.Category),
		})
		if err != nil {
			return fmt.Errorf("couldn't create monster formation list: %s - %s - %d - %s - %d - %d: %v", locationArea.Location, locationArea.SubLocation, derefOrNil(locationArea.SVersion), locationArea.Area, derefOrNil(locationArea.AVersion), derefOrNil(shop.Version), err)
		}
	}
	return nil

}