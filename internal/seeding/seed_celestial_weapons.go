package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type CelestialWeapon struct {
	ID				int32
	Name        	string	`json:"name"`
	Character		string	`json:"character"`
	CharacterID		*int32
	KeyItemBase 	string	`json:"key_item_base"`
	Aeon			*string	`json:"aeon"`
	AeonID			*int32
	Formula     	string	`json:"formula"`
}

func (cw CelestialWeapon) ToHashFields() []any {
	return []any{
		cw.Name,
		cw.KeyItemBase,
		cw.Formula,
	}
}

func (cw CelestialWeapon) GetID() int32 {
	return cw.ID
}


func (l *lookup) seedCelestialWeapons(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/celestial_weapons.json"

	var celestialWeapons []CelestialWeapon
	err := loadJSONFile(string(srcPath), &celestialWeapons)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, weapon := range celestialWeapons {
			dbWeapon, err := qtx.CreateCelestialWeapon(context.Background(), database.CreateCelestialWeaponParams{
				DataHash:    generateDataHash(weapon),
				Name:        weapon.Name,
				KeyItemBase: database.KeyItemBase(weapon.KeyItemBase),
				Formula:     database.CelestialFormula(weapon.Formula),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Celestial Weapon: %s: %v", weapon.Name, err)
			}

			weapon.ID = dbWeapon.ID
			l.celestialWeapons[weapon.Name] = weapon
		}
		return nil
	})
}



func (l *lookup) seedCelestialWeaponsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/celestial_weapons.json"

	var celestialWeapons []CelestialWeapon
	err := loadJSONFile(string(srcPath), &celestialWeapons)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonWeapon := range celestialWeapons {
			weapon, err := l.getCelestialWeapon(jsonWeapon.Name)
			if err != nil {
				return err
			}

			weapon.CharacterID, err = assignFKPtr(&weapon.Character, l.getCharacter)
			if err != nil {
				return err
			}

			weapon.AeonID, err = assignFKPtr(weapon.Aeon, l.getAeon)
			if err != nil {
				return err
			}

			err = qtx.UpdateCelestialWeapon(context.Background(), database.UpdateCelestialWeaponParams{
				DataHash:    	generateDataHash(weapon),
				CharacterID: 	getNullInt32(weapon.CharacterID),
				AeonID: 		getNullInt32(weapon.AeonID),
				ID:				weapon.ID,
			})
			if err != nil {
				return fmt.Errorf("couldn't update Celestial Weapon: %s: %v", weapon.Name, err)
			}
		}

		return nil
	})
}