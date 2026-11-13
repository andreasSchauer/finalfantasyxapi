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

func loop4SeedCharacters(qtx *database.Queries, ctx context.Context) error {
	chars, err := extractCharacters()
	if err != nil {
		return err
	}

	params := database.CreateCharacterBulkParams{
		DataHash:            make([]string, len(chars)),
		UnitID:              make([]int32, len(chars)),
		IsStoryBased:        make([]bool, len(chars)),
		WeaponType:          make([]database.WeaponType, len(chars)),
		ArmorType:           make([]database.ArmorType, len(chars)),
		PhysicalAttackRange: make([]int32, len(chars)),
		CanFightUnderwater:  make([]bool, len(chars)),
		AreaID:              make([]sql.NullInt32, len(chars)),
	}

	for i, c := range chars {
		params.DataHash[i] = generateDataHash(c)
		params.UnitID[i] = c.PlayerUnit.ID
		params.IsStoryBased[i] = c.IsStoryBased
		params.WeaponType[i] = database.WeaponType(c.WeaponType)
		params.ArmorType[i] = database.ArmorType(c.ArmorType)
		params.PhysicalAttackRange[i] = c.PhysAtkRange
		params.CanFightUnderwater[i] = c.CanFightUnderwater
		params.AreaID[i] = h.GetNullInt32(c.AreaID)
	}

	dbRows, err := qtx.CreateCharacterBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create characters: %v", err)
	}

	for i, row := range dbRows {
		chars[i].ID = row.ID
		rawJSON.characters[i].ID = row.ID
		lookup.L.Characters[seedtypes.Key(chars[i])] = chars[i]
		lookup.L.CharactersID[row.ID] = chars[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractCharacters() ([]seedtypes.Character, error) {
	chars := []seedtypes.Character{}
	var err error

	for i := range rawJSON.characters {
		char := &rawJSON.characters[i]

		char.PlayerUnit.ID, err = getHashID(char.PlayerUnit)
		if err != nil {
			return nil, err
		}

		char.AreaID, err = assignFKPtr(&char.LocationArea, lookup.L.Areas)
		if err != nil {
			return nil, err
		}

		if char.StdSphereGrid != nil {
			char.StdSphereGrid.ID, err = getHashID(*char.StdSphereGrid)
			if err != nil {
				return nil, err
			}
		}

		if char.ExpSphereGrid != nil {
			char.ExpSphereGrid.ID, err = getHashID(*char.ExpSphereGrid)
			if err != nil {
				return nil, err
			}
		}

		chars = append(chars, *char)
	}

	return dedupeRows(chars, hashes), nil
}

func completeCharacters() error {
	for i := range rawJSON.characters {
		character := &rawJSON.characters[i]
		err := assignIDs(character.BaseStats)
		if err != nil {
			return err
		}

		lookup.L.Characters[seedtypes.Key(character)] = *character
		lookup.L.CharactersID[character.ID] = *character
	}

	return nil
}

func getCharacterBaseStats(c seedtypes.Character) ([]seedtypes.BaseStat, error) {
	return c.BaseStats, nil
}

func seedJuncCharactersBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "characters + base stats"
	jParams, err := processJunctions(desc, rawJSON.characters, getCharacterBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateCharactersBaseStatsJunctionBulk(ctx, database.CreateCharactersBaseStatsJunctionBulkParams{
		DataHash:    jParams.DataHashes,
		CharacterID: jParams.ParentIDs,
		BaseStatID:  jParams.ChildIDs,
	})
}
