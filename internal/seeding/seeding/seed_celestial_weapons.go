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

func loop5SeedCelestialWeapons(qtx *database.Queries, ctx context.Context) error {
	wpns, err := extractCelestialWeapons()
	if err != nil {
		return err
	}

	params := database.CreateCelestialWeaponBulkParams{
		DataHash:    make([]string, len(wpns)),
		Name:        make([]string, len(wpns)),
		KeyItemBase: make([]database.KeyItemBase, len(wpns)),
		Formula:     make([]database.CelestialFormula, len(wpns)),
		CharacterID: make([]sql.NullInt32, len(wpns)),
		AeonID:      make([]sql.NullInt32, len(wpns)),
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
		rawJSON.celestialWeapons[i].ID = row.ID
		lookup.L.CelestialWeapons[seedtypes.Key(wpns[i])] = wpns[i]
		lookup.L.CelestialWeaponsID[row.ID] = wpns[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractCelestialWeapons() ([]seedtypes.CelestialWeapon, error) {
	wpns := []seedtypes.CelestialWeapon{}
	var err error

	for i := range rawJSON.celestialWeapons {
		wpn := &rawJSON.celestialWeapons[i]

		wpn.CharacterID, err = assignFKPtr(&wpn.Character, lookup.L.Characters)
		if err != nil {
			return nil, err
		}

		wpn.AeonID, err = assignFKPtr(wpn.Aeon, lookup.L.Aeons)
		if err != nil {
			return nil, err
		}

		wpns = append(wpns, *wpn)
	}

	return dedupeRows(wpns, hashes), nil
}
