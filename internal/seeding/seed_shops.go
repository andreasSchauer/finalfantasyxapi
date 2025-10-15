package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Shop struct {
	//id 		int32
	//dataHash	string
	Version      *int32       `json:"version"`
	LocationArea LocationArea `json:"location_area"`
	AreaID       int32
	Notes        *string `json:"notes"`
	Category     string  `json:"category"`
}

func (s Shop) ToHashFields() []any {
	return []any{
		derefOrNil(s.Version),
		s.AreaID,
		derefOrNil(s.Notes),
		s.Category,
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
			locationArea := shop.LocationArea
			area, err := l.getArea(shop.LocationArea)
			if err != nil {
				return fmt.Errorf("shops: %v", err)
			}

			shop.AreaID = area.ID

			err = qtx.CreateShop(context.Background(), database.CreateShopParams{
				DataHash: generateDataHash(shop),
				Version:  getNullInt32(shop.Version),
				AreaID:   shop.AreaID,
				Notes:    getNullString(shop.Notes),
				Category: database.ShopCategory(shop.Category),
			})
			if err != nil {
				return fmt.Errorf("couldn't create monster formation list: %s - %s - %s - %d - %d: %v", locationArea.Location, locationArea.SubLocation, locationArea.Area, derefOrNil(locationArea.Version), derefOrNil(shop.Version), err)
			}
		}
		return nil
	})
}
