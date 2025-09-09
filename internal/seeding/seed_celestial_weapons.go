package seeding

import (
	"context"
	"database/sql"
	"fmt"
	
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type CelestialWeapon struct {
	//id 			int32
	//dataHash		string
	Name        	string		`json:"name"`
	KeyItemBase		string		`json:"key_item_base"`
	Formula			string		`json:"formula"`
}

func (cw CelestialWeapon) ToHashFields() []any {
	return []any{
		cw.Name,
		cw.KeyItemBase,
		cw.Formula,
	}
}


func seedCelestialWeapons(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/celestial_weapons.json"

	var celestialWeapons []CelestialWeapon
	err := loadJSONFile(string(srcPath), &celestialWeapons)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, weapon := range celestialWeapons {
			err = qtx.CreateCelestialWeapon(context.Background(), database.CreateCelestialWeaponParams{
				DataHash:     	generateDataHash(weapon),
				Name:         	weapon.Name,
				KeyItemBase: 	database.KeyItemBase(weapon.KeyItemBase),
				Formula: 		database.CelestialFormula(weapon.Formula),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Celestial Weapon: %s: %v", weapon.Name, err)
			}
		}
		return nil
	})
}
