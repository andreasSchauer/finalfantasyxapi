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
	Name				string 		`json:"name"`
	WeaponType			string		`json:"weapon_type"`
	ArmorType			string		`json:"armor_type"`
	PhysAtkRange		int32		`json:"physical_attack_range"`
	CanFightUnderwater	bool		`json:"can_fight_underwater"`
}

func(c Character) ToHashFields() []any {
	return []any{
		c.Name,
		c.WeaponType,
		c.ArmorType,
		c.PhysAtkRange,
		c.CanFightUnderwater,
	}
}


func seedCharacters(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/characters.json"

	var characters []Character
	err := loadJSONFile(string(srcPath), &characters)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, character := range characters {
			err = qtx.CreateCharacter(context.Background(), database.CreateCharacterParams{
				DataHash: 				generateDataHash(character),
				Name: 					character.Name,
				WeaponType: 			database.WeaponType(character.WeaponType),
				ArmorType: 				database.ArmorType(character.ArmorType),
				PhysicalAttackRange: 	character.PhysAtkRange,
				CanFightUnderwater: 	character.CanFightUnderwater,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Character: %s: %v", character.Name, err)
			}
		}
		return nil
	})
}