package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Character struct {
	//id 		int32
	//dataHash	string
	PlayerUnit
	UnitID             int32
	StoryOnly          bool       `json:"story_only"`
	WeaponType         string     `json:"weapon_type"`
	ArmorType          string     `json:"armor_type"`
	PhysAtkRange       int32      `json:"physical_attack_range"`
	CanFightUnderwater bool       `json:"can_fight_underwater"`
	BaseStats          []BaseStat `json:"base_stats"`
}

func (c Character) ToHashFields() []any {
	return []any{
		c.UnitID,
		c.StoryOnly,
		c.WeaponType,
		c.ArmorType,
		c.PhysAtkRange,
		c.CanFightUnderwater,
	}
}

type CharacterLookup struct {
	Character
	ID int32
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
			character.Type = database.UnitTypeCharacter

			dbPlayerUnit, err := l.seedPlayerUnit(qtx, character.PlayerUnit)
			if err != nil {
				return err
			}

			character.UnitID = dbPlayerUnit.ID

			dbCharacter, err := qtx.CreateCharacter(context.Background(), database.CreateCharacterParams{
				DataHash:            generateDataHash(character),
				UnitID:              character.UnitID,
				StoryOnly:           character.StoryOnly,
				WeaponType:          database.WeaponType(character.WeaponType),
				ArmorType:           database.ArmorType(character.ArmorType),
				PhysicalAttackRange: character.PhysAtkRange,
				CanFightUnderwater:  character.CanFightUnderwater,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Character: %s: %v", character.Name, err)
			}

			key := createLookupKey(character.PlayerUnit)
			l.characters[key] = CharacterLookup{
				Character: character,
				ID:        dbCharacter.ID,
			}

			err = l.seedCharacterClasses(qtx, character.PlayerUnit)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (l *lookup) createCharactersRelationships(db *database.Queries, dbConn *sql.DB) error {
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

			for i, baseStat := range character.BaseStats {
				dbBaseStat, err := l.seedBaseStat(qtx, baseStat)
				if err != nil {
					return err
				}
				baseStat.StatID = dbBaseStat.StatID
				character.BaseStats[i] = baseStat

				junction := Junction{
					ParentID: 	character.ID,
					ChildID: 	dbBaseStat.ID,
				}

				err = qtx.CreateCharacterBaseStatJunction(context.Background(), database.CreateCharacterBaseStatJunctionParams{
					DataHash:    generateDataHash(junction),
					CharacterID: junction.ParentID,
					BaseStatID:  junction.ChildID,
				})
				if err != nil {
					return err
				}
			}

			l.characters[character.Name] = character
		}
		return nil
	})
}
