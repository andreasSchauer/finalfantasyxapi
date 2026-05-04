package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop4SeedCharacters(qtx *database.Queries, ctx context.Context) error {
	chars, err := l.extractCharacters()
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
		l.json.characters[i].ID = row.ID
		l.Characters[chars[i].Name] = chars[i]
		l.CharactersID[row.ID] = chars[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractCharacters() ([]Character, error) {
	chars := []Character{}
	var err error

	for i := range l.json.characters {
		char := &l.json.characters[i]

		char.PlayerUnit.ID, err = l.getHashID(char.PlayerUnit)
		if err != nil {
			return nil, err
		}

		char.AreaID, err = assignFKPtr(&char.LocationArea, l.Areas)
		if err != nil {
			return nil, err
		}

		chars = append(chars, *char)
	}

	return dedupeRows(chars, l.Hashes), nil
}

func (l *Lookup) completeCharacters() error {
	for i := range l.json.characters {
		character := &l.json.characters[i]
		err := assignIDs(l, character.BaseStats)
		if err != nil {
			return err
		}

		l.Characters[character.Name] = *character
		l.CharactersID[character.ID] = *character
	}

	return nil
}

func (l *Lookup) getCharacterBaseStats(c Character) ([]BaseStat, error) {
	return c.BaseStats, nil
}

func (l *Lookup) seedJuncCharactersBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "characters + base stats"
	jParams, err := processJunctions(l, desc, l.json.characters, l.getCharacterBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateCharactersBaseStatsJunctionBulk(ctx, database.CreateCharactersBaseStatsJunctionBulkParams{
		DataHash:    jParams.DataHashes,
		CharacterID: jParams.ParentIDs,
		BaseStatID:  jParams.ChildIDs,
	})
}
