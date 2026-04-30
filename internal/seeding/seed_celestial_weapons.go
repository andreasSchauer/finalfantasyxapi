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
		fmt.Sprintf("%T", cw),
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

func (cw CelestialWeapon) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID: 	cw.ID,
		Name: 	cw.Name,
	}
}

func (l *Lookup) seedCelestialWeapons(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/celestial_weapons.json"

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
				return h.NewErr(weapon.Error(), err, "couldn't create celestial weapon")
			}

			weapon.ID = dbWeapon.ID
			l.CelestialWeapons[weapon.Name] = weapon
			l.CelestialWeaponsID[weapon.ID] = weapon
		}
		return nil
	})
}

func (l *Lookup) seedCelestialWeaponsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/celestial_weapons.json"

	var celestialWeapons []CelestialWeapon
	err := loadJSONFile(string(srcPath), &celestialWeapons)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonWeapon := range celestialWeapons {
			weapon, err := GetResource(jsonWeapon.Name, l.CelestialWeapons)
			if err != nil {
				return err
			}

			weapon.CharacterID, err = assignFKPtr(&weapon.Character, l.Characters)
			if err != nil {
				return h.NewErr(weapon.Error(), err)
			}

			weapon.AeonID, err = assignFKPtr(weapon.Aeon, l.Aeons)
			if err != nil {
				return h.NewErr(weapon.Error(), err)
			}

			err = qtx.UpdateCelestialWeapon(context.Background(), database.UpdateCelestialWeaponParams{
				DataHash:    generateDataHash(weapon),
				CharacterID: h.GetNullInt32(weapon.CharacterID),
				AeonID:      h.GetNullInt32(weapon.AeonID),
				ID:          weapon.ID,
			})
			if err != nil {
				return h.NewErr(weapon.Error(), err, "couldn't update celestial weapon")
			}
		}

		return nil
	})
}


func (l *Lookup) loop5SeedCelestialWeapons(qtx *database.Queries, ctx context.Context) error {
	wpns, err := l.extractCelestialWeapons()
	if err != nil {
		return err
	}

	params := database.CreateCelestialWeaponBulkParams{
		DataHash:   	make([]string, len(wpns)),
		Name: 			make([]string, len(wpns)),
		KeyItemBase: 	make([]database.KeyItemBase, len(wpns)),
		Formula: 		make([]database.CelestialFormula, len(wpns)),
		CharacterID: 	make([]sql.NullInt32, len(wpns)),
		AeonID: 		make([]sql.NullInt32, len(wpns)),
	}

	for i, cw := range wpns {
		params.DataHash[i] = generateDataHash(cw)
		params.Name[i] = cw.Name
		params.KeyItemBase[i] = database.KeyItemBase(cw.KeyItemBase)
		params.Formula[i] = database.CelestialFormula(cw.Formula)
		params.CharacterID[i] = h.GetNullInt32(cw.CharacterID)
		params.AeonID[i] = h.GetNullInt32(cw.AeonID)
	}

	dbRows, err := qtx.CreateCelestialWeaponBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create celestial weapons: %v", err)
	}

	for i, row := range dbRows {
		wpns[i].ID = row.ID
		l.json.celestialWeapons[i].ID = row.ID
		l.CelestialWeapons[wpns[i].Name] = wpns[i]
		l.CelestialWeaponsID[row.ID] = wpns[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractCelestialWeapons() ([]CelestialWeapon, error) {
	wpns := []CelestialWeapon{}
	var err error

	for i := range l.json.celestialWeapons {
		wpn := &l.json.celestialWeapons[i]

		wpn.CharacterID, err = assignFKPtr(&wpn.Character, l.Characters)
		if err != nil {
			return nil, err
		}

		wpn.AeonID, err = assignFKPtr(wpn.Aeon, l.Aeons)
		if err != nil {
			return nil, err
		}

		wpns = append(wpns, *wpn)
	}

	return dedupeRows(wpns, l.Hashes), nil
}