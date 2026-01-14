package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterArenaCreation struct {
	SubquestID                int32
	Name                      string  `json:"name"`
	Category                  string  `json:"category"`
	RequiredArea              *string `json:"required_area"`
	RequiredSpecies           *string `json:"required_species"`
	UnderwaterOnly            bool    `json:"underwater_only"`
	CreationsUnlockedCategory *string `json:"creations_unlocked_category"`
	Amount                    int32   `json:"amount"`
}

func (m MonsterArenaCreation) ToHashFields() []any {
	return []any{
		m.SubquestID,
		m.Category,
		h.DerefOrNil(m.RequiredArea),
		h.DerefOrNil(m.RequiredSpecies),
		m.UnderwaterOnly,
		h.DerefOrNil(m.CreationsUnlockedCategory),
		m.Amount,
	}
}

func (m MonsterArenaCreation) Error() string {
	return fmt.Sprintf("monster arena creation %s", m.Name)
}

func (l *Lookup) seedMonsterArenaCreations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_arena_creations.json"

	var creations []MonsterArenaCreation
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

			err = qtx.CreateMonsterArenaCreation(context.Background(), database.CreateMonsterArenaCreationParams{
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
		}
		return nil
	})
}
