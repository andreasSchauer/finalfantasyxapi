package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Character struct {
	ID int32
	PlayerUnit
	LocationArea       LocationArea `json:"location_area"`
	AreaID             *int32
	IsStoryBased       bool       `json:"is_story_based"`
	WeaponType         string     `json:"weapon_type"`
	ArmorType          string     `json:"armor_type"`
	PhysAtkRange       int32      `json:"physical_attack_range"`
	CanFightUnderwater bool       `json:"can_fight_underwater"`
	BaseStats          []BaseStat `json:"base_stats"`
}

func (c Character) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", c),
		c.PlayerUnit.ID,
		h.DerefOrNil(c.AreaID),
		c.IsStoryBased,
		c.WeaponType,
		c.ArmorType,
		c.PhysAtkRange,
		c.CanFightUnderwater,
	}
}

func (c Character) GetID() int32 {
	return c.ID
}

func (c Character) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (c Character) Error() string {
	return fmt.Sprintf("character %s", c.Name)
}

func (l *Lookup) seedCharacters(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/characters.json"

	var characters []Character
	err := loadJSONFile(string(srcPath), &characters)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, character := range characters {
			var err error
			character.Type = database.UnitTypeCharacter

			character.PlayerUnit, err = seedObjAssignID(qtx, character.PlayerUnit, l.seedPlayerUnit)
			if err != nil {
				return h.NewErr(character.Error(), err)
			}

			dbCharacter, err := qtx.CreateCharacter(context.Background(), database.CreateCharacterParams{
				DataHash:            generateDataHash(character),
				UnitID:              character.PlayerUnit.ID,
				IsStoryBased:        character.IsStoryBased,
				WeaponType:          database.WeaponType(character.WeaponType),
				ArmorType:           database.ArmorType(character.ArmorType),
				PhysicalAttackRange: character.PhysAtkRange,
				CanFightUnderwater:  character.CanFightUnderwater,
			})
			if err != nil {
				return h.NewErr(character.Error(), err, "couldn't create character")
			}

			character.ID = dbCharacter.ID
			l.Characters[character.Name] = character
			l.CharactersID[character.ID] = character
		}

		return nil
	})
}

func (l *Lookup) seedCharactersRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/characters.json"

	var characters []Character
	err := loadJSONFile(string(srcPath), &characters)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCharacter := range characters {
			character, err := GetResource(jsonCharacter.Name, l.Characters)
			if err != nil {
				return err
			}

			character.AreaID, err = assignFKPtr(&character.LocationArea, l.Areas)
			if err != nil {
				return h.NewErr(character.Error(), err)
			}

			err = qtx.UpdateCharacter(context.Background(), database.UpdateCharacterParams{
				DataHash: generateDataHash(character),
				AreaID:   h.GetNullInt32(character.AreaID),
				ID:       character.ID,
			})
			if err != nil {
				return h.NewErr(character.Error(), err, "couldn't update character")
			}

			err = l.seedCharacterBaseStats(qtx, character)
			if err != nil {
				return h.NewErr(character.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedCharacterBaseStats(qtx *database.Queries, character Character) error {
	for _, baseStat := range character.BaseStats {
		junction, err := createJunctionSeed(qtx, character, baseStat, l.seedBaseStat)
		if err != nil {
			return h.NewErr(character.Error(), err)
		}

		err = qtx.CreateCharactersBaseStatsJunction(context.Background(), database.CreateCharactersBaseStatsJunctionParams{
			DataHash:    generateDataHash(junction),
			CharacterID: junction.ParentID,
			BaseStatID:  junction.ChildID,
		})
		if err != nil {
			return h.NewErr(baseStat.Error(), err, "couldn't junction base stat")
		}
	}

	return nil
}



func (l *Lookup) loop4SeedCharacters(qtx *database.Queries, ctx context.Context) error {
	chars, err := l.extractCharacters()
	if err != nil {
		return err
	}

	params := database.CreateCharacterBulkParams{
		DataHash:   make([]string, len(chars)),
		UnitID: make([]int32, len(chars)),
		IsStoryBased: make([]bool, len(chars)),
		WeaponType: make([]database.WeaponType, len(chars)),
		ArmorType: make([]database.ArmorType, len(chars)),
		PhysicalAttackRange: make([]int32, len(chars)),
		CanFightUnderwater: make([]bool, len(chars)),
		AreaID: make([]sql.NullInt32, len(chars)),
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