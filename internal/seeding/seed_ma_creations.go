package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type MonsterArenaCreation struct {
	//id 		int32
	//dataHash	string
	Name						string 		`json:"name"`
	Category					string		`json:"category"`
	RequiredArea				*string		`json:"required_area"`
	RequiredSpecies				*string		`json:"required_species"`
	UnderwaterOnly				bool		`json:"underwater_only"`
	CreationsUnlockedCategory	*string		`json:"creations_unlocked_category"`
	Amount						int32		`json:"amount"`
}

func(m MonsterArenaCreation) ToHashFields() []any {
	return []any{
		m.Name,
		m.Category,
		derefOrNil(m.RequiredArea),
		derefOrNil(m.RequiredSpecies),
		m.UnderwaterOnly,
		derefOrNil(m.CreationsUnlockedCategory),
		m.Amount,
	}
}


func seedMonsterArenaCreations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_arena_creations.json"

	var creations []MonsterArenaCreation
	err := loadJSONFile(string(srcPath), &creations)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, creation := range creations {
			err = qtx.CreateMonsterArenaCreation(context.Background(), database.CreateMonsterArenaCreationParams{
				DataHash: 					generateDataHash(creation),
				Name: 						creation.Name,
				Category: 					database.MaCreationCategory(creation.Category),
				RequiredArea: 				nullMaCreationArea(creation.RequiredArea),
				RequiredSpecies: 			nullMaCreationSpecies(creation.RequiredSpecies),
				UnderwaterOnly: 			creation.UnderwaterOnly,
				CreationsUnlockedCategory: 	nullCreationsUnlockedCategory(creation.CreationsUnlockedCategory),
				Amount: 					creation.Amount,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Monster Arena Creation: %s: %v", creation.Name, err)
			}
		}
		return nil
	})
}