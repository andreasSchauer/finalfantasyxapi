package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Character struct {
	ID int32
	PlayerUnit
	StoryOnly          bool       `json:"story_only"`
	WeaponType         string     `json:"weapon_type"`
	ArmorType          string     `json:"armor_type"`
	PhysAtkRange       int32      `json:"physical_attack_range"`
	CanFightUnderwater bool       `json:"can_fight_underwater"`
	BaseStats          []BaseStat `json:"base_stats"`
}

func (c Character) ToHashFields() []any {
	return []any{
		c.PlayerUnit.ID,
		c.StoryOnly,
		c.WeaponType,
		c.ArmorType,
		c.PhysAtkRange,
		c.CanFightUnderwater,
	}
}

func (c Character) GetID() int32 {
	return c.ID
}

func (c Character) Error() string {
	return fmt.Sprintf("character %s", c.Name)
}

func (l *lookup) seedCharacters(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/characters.json"

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
				return getErr(character, err)
			}

			dbCharacter, err := qtx.CreateCharacter(context.Background(), database.CreateCharacterParams{
				DataHash:            generateDataHash(character),
				UnitID:              character.PlayerUnit.ID,
				StoryOnly:           character.StoryOnly,
				WeaponType:          database.WeaponType(character.WeaponType),
				ArmorType:           database.ArmorType(character.ArmorType),
				PhysicalAttackRange: character.PhysAtkRange,
				CanFightUnderwater:  character.CanFightUnderwater,
			})
			if err != nil {
				return getDbErr(character, err, "couldn't create character")
			}

			character.ID = dbCharacter.ID
			key := createLookupKey(character.PlayerUnit)
			l.characters[key] = character

			err = l.seedCharacterClasses(qtx, character.PlayerUnit)
			if err != nil {
				return getErr(character, err)
			}
		}
		return nil
	})
}

func (l *lookup) seedCharactersRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/characters.json"

	var characters []Character
	err := loadJSONFile(string(srcPath), &characters)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCharacter := range characters {
			character, err := l.getCharacter(jsonCharacter.Name)
			if err != nil {
				return err
			}

			err = l.seedCharacterBaseStats(qtx, character)
			if err != nil {
				return getErr(character, err)
			}
		}

		return nil
	})
}


func (l *lookup) seedCharacterBaseStats(qtx *database.Queries, character Character) error {
	for _, baseStat := range character.BaseStats {
		junction, err := createJunctionSeed(qtx, character, baseStat, l.seedBaseStat)
		if err != nil {
			return getErr(character, err)
		}

		err = qtx.CreateCharactersBaseStatsJunction(context.Background(), database.CreateCharactersBaseStatsJunctionParams{
			DataHash:    generateDataHash(junction),
			CharacterID: junction.ParentID,
			BaseStatID:  junction.ChildID,
		})
		if err != nil {
			return getDbErr(baseStat, err, "couldn't junction base stat")
		}
	}

	return nil
}