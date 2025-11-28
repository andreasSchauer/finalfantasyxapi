package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type CelestialWeapon struct {
	ID          int32
	Name        string `json:"name"`
	Character   string `json:"character"`
	CharacterID *int32
	KeyItemBase string  `json:"key_item_base"`
	Aeon        *string `json:"aeon"`
	AeonID      *int32
	Formula     string `json:"formula"`
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

func (cw CelestialWeapon) Error() string {
	return fmt.Sprintf("celestial weapon %s", cw.Name)
}

func (l *Lookup) seedCelestialWeapons(db *database.Queries, dbConn *sql.DB) error {
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
				return h.GetErr(weapon.Error(), err, "couldn't create celestial weapon")
			}

			weapon.ID = dbWeapon.ID
			l.celestialWeapons[weapon.Name] = weapon
		}
		return nil
	})
}

func (l *Lookup) seedCelestialWeaponsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/celestial_weapons.json"

	var celestialWeapons []CelestialWeapon
	err := loadJSONFile(string(srcPath), &celestialWeapons)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonWeapon := range celestialWeapons {
			weapon, err := getResource(jsonWeapon.Name, l.celestialWeapons)
			if err != nil {
				return err
			}

			weapon.CharacterID, err = assignFKPtr(&weapon.Character, l.characters)
			if err != nil {
				return h.GetErr(weapon.Error(), err)
			}

			weapon.AeonID, err = assignFKPtr(weapon.Aeon, l.aeons)
			if err != nil {
				return h.GetErr(weapon.Error(), err)
			}

			err = qtx.UpdateCelestialWeapon(context.Background(), database.UpdateCelestialWeaponParams{
				DataHash:    generateDataHash(weapon),
				CharacterID: h.GetNullInt32(weapon.CharacterID),
				AeonID:      h.GetNullInt32(weapon.AeonID),
				ID:          weapon.ID,
			})
			if err != nil {
				return h.GetErr(weapon.Error(), err, "couldn't update celestial weapon")
			}
		}

		return nil
	})
}
