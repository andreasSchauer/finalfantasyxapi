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
	Version		*int32		`json:"version"`
	Notes		*string 	`json:"notes"`
	Category	string		`json:"category"`
}

func(s Shop) ToHashFields() []any {
	return []any{
		derefOrNil(s.Version),
		derefOrNil(s.Notes),
		s.Category,
	}
}


func seedShops(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/shops.json"

	var shops []Shop
	err := loadJSONFile(string(srcPath), &shops)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for i, shop := range shops {
			err = qtx.CreateShop(context.Background(), database.CreateShopParams{
				DataHash: 				generateDataHash(shop),
				Version: 				getNullInt32(shop.Version),
				Notes: 					getNullString(shop.Notes),
				Category: 				database.ShopCategory(shop.Category),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Shop: %d: %v", i, err)
			}
		}
		return nil
	})
}