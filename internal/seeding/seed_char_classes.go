package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type CharacterClass struct {
	ID       int32
	Name     string   `json:"name"`
	Category string   `json:"category"`
	Members  []string `json:"members"`
}

func (cc CharacterClass) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", cc),
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
		ID:   cc.ID,
		Name: cc.Name,
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

			for _, jsonUnit := range class.Members {
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

func (l *Lookup) loop1SeedCharacterClasses(qtx *database.Queries, ctx context.Context) error {
	classes := dedupeRows(l.json.characterClasses, l.Hashes)

	params := database.CreateCharacterClassBulkParams{
		DataHash: make([]string, len(classes)),
		Name:     make([]string, len(classes)),
		Category: make([]database.CharacterClassCategory, len(classes)),
	}

	for i, c := range classes {
		params.DataHash[i] = generateDataHash(c)
		params.Name[i] = c.Name
		params.Category[i] = database.CharacterClassCategory(c.Category)
	}

	dbRows, err := qtx.CreateCharacterClassBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create character classes: %v", err)
	}

	for i, row := range dbRows {
		classes[i].ID = row.ID
		l.json.characterClasses[i].ID = row.ID
		l.CharClasses[classes[i].Name] = classes[i]
		l.CharClassesID[row.ID] = classes[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getCharacterClassPlayerUnits(c CharacterClass) ([]PlayerUnit, error) {
	return toObjects(c.Members, l.PlayerUnits)
}

func (l *Lookup) getDefaultAbilitiesCharacterClasses(entries []DefaultAbilitiesEntry) ([]CharacterClass, error) {
	classes := make([]CharacterClass, len(entries))

	for i, entry := range entries {
		class, err := GetResource(entry.Name, l.CharClasses)
		if err != nil {
			return nil, err
		}
		classes[i] = class
	}

	return classes, nil
}

func (l *Lookup) getDefaultAbilities(entry DefaultAbilitiesEntry) ([]Ability, error) {
	return toObjects(entry.DefaultAbilities, l.Abilities)
}

func (l *Lookup) seedJuncCharacterClassesPlayerUnits(qtx *database.Queries, ctx context.Context) error {
	const desc string = "character classes + player units"
	jParams, err := processJunctions(l, desc, l.json.characterClasses, l.getCharacterClassPlayerUnits)
	if err != nil {
		return err
	}

	return qtx.CreateCharacterClassPlayerUnitsJunctionBulk(ctx, database.CreateCharacterClassPlayerUnitsJunctionBulkParams{
		DataHash: jParams.DataHashes,
		ClassID:  jParams.ParentIDs,
		UnitID:   jParams.ChildIDs,
	})
}

