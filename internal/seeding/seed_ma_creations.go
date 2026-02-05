package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type ArenaCreation struct {
	ID                        int32
	SubquestID                int32
	MonsterID				  *int32
	Name                      string  `json:"name"`
	Category                  string  `json:"category"`
	RequiredArea              *string `json:"required_area"`
	RequiredSpecies           *string `json:"required_species"`
	UnderwaterOnly            bool    `json:"underwater_only"`
	CreationsUnlockedCategory *string `json:"creations_unlocked_category"`
	Amount                    int32   `json:"amount"`
}

func (a ArenaCreation) ToHashFields() []any {
	return []any{
		a.SubquestID,
		h.DerefOrNil(a.MonsterID),
		a.Category,
		h.DerefOrNil(a.RequiredArea),
		h.DerefOrNil(a.RequiredSpecies),
		a.UnderwaterOnly,
		h.DerefOrNil(a.CreationsUnlockedCategory),
		a.Amount,
	}
}

func (a ArenaCreation) GetID() int32 {
	return a.ID
}

func (a ArenaCreation) Error() string {
	return fmt.Sprintf("monster arena creation %s", a.Name)
}

func (a ArenaCreation) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID: a.ID,
		Name: a.Name,
	}
}

func (l *Lookup) seedArenaCreations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_arena_creations.json"

	var creations []ArenaCreation
	err := loadJSONFile(string(srcPath), &creations)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, creation := range creations {
			var err error

			creation.SubquestID, err = assignFK(creation.Name, l.Subquests)
			if err != nil {
				return h.NewErr(creation.Error(), err)
			}

			dbCreation, err := qtx.CreateMonsterArenaCreation(context.Background(), database.CreateMonsterArenaCreationParams{
				DataHash:                  generateDataHash(creation),
				SubquestID:                creation.SubquestID,
				Category:                  database.MaCreationCategory(creation.Category),
				RequiredArea:              h.NullMaCreationArea(creation.RequiredArea),
				RequiredSpecies:           h.NullMaCreationSpecies(creation.RequiredSpecies),
				UnderwaterOnly:            creation.UnderwaterOnly,
				CreationsUnlockedCategory: h.NullCreationsUnlockedCategory(creation.CreationsUnlockedCategory),
				Amount:                    creation.Amount,
			})
			if err != nil {
				return h.NewErr(creation.Error(), err, "couldn't create monster arena creation")
			}
			creation.ID = dbCreation.ID
			l.ArenaCreations[creation.Name] = creation
			l.ArenaCreationsID[creation.ID] = creation

		}
		return nil
	})
}


func (l *Lookup) seedArenaCreationsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_arena_creations.json"

	var creations []ArenaCreation
	err := loadJSONFile(string(srcPath), &creations)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCreation := range creations {
			obj := LookupObject{
				Name: jsonCreation.Name,
			}
			creation, err := GetResource(jsonCreation.Name, l.ArenaCreations)
			if err != nil {
				return err
			}

			creation.MonsterID, err = assignFKPtr(&obj, l.Monsters)
			if err != nil {
				return h.NewErr(creation.Error(), err)
			}

			err = qtx.UpdateMonsterArenaCreation(context.Background(), database.UpdateMonsterArenaCreationParams{
				DataHash: generateDataHash(creation),
				MonsterID: h.GetNullInt32(creation.MonsterID),
				ID:       creation.ID,
			})
			if err != nil {
				return h.NewErr(creation.Error(), err, "couldn't update stat")
			}

			l.ArenaCreations[creation.Name] = creation
			l.ArenaCreationsID[creation.ID] = creation
		}
		return nil
	})
}