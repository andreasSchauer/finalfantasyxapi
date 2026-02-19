package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type CharacterClass struct {
	ID   		int32
	Name 		string		`json:"name"`
	Category	string		`json:"category"`
	Units		[]string	`json:"units"`
}

func (cc CharacterClass) ToHashFields() []any {
	return []any{
		cc.Name,
	}
}

func (cc CharacterClass) GetID() int32 {
	return cc.ID
}

func (cc CharacterClass) Error() string {
	return fmt.Sprintf("character class %s", cc.Name)
}

func (cc CharacterClass) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID: 			cc.ID,
		Name: 			cc.Name,
	}
}

func (l *Lookup) seedCharacterClasses(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/character_classes.json"

	var classes []CharacterClass
	err := loadJSONFile(string(srcPath), &classes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, class := range classes {
			dbClass, err := qtx.CreateCharacterClass(context.Background(), database.CreateCharacterClassParams{
				DataHash: generateDataHash(class),
				Name:     class.Name,
				Category: database.CharacterClassCategory(class.Category),
			})
			if err != nil {
				return h.NewErr(class.Error(), err, "couldn't create character class")
			}

			class.ID = dbClass.ID
			l.CharClasses[class.Name] = class
			l.CharClassesID[class.ID] = class
		}
		return nil
	})
}


func (l *Lookup) seedCharacterClassesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/character_classes.json"

	var classes []CharacterClass
	err := loadJSONFile(string(srcPath), &classes)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonClass := range classes {
			class, err := GetResource(jsonClass.Name, l.CharClasses)
			if err != nil {
				return err
			}

			for _, jsonUnit := range class.Units {
				junction, err := createJunction(class, jsonUnit, l.PlayerUnits)
				if err != nil {
					return err
				}
	
				err = qtx.CreateCharacterClassPlayerUnitsJunction(context.Background(), database.CreateCharacterClassPlayerUnitsJunctionParams{
					DataHash: generateDataHash(junction),
					ClassID:  junction.ParentID,
					UnitID:   junction.ChildID,
				})
				if err != nil {
					return h.NewErr(jsonUnit, err, "couldn't junction player unit")
				}
			}
		}
		return nil
	})
}

