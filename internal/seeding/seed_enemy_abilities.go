package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type EnemyAbility struct {
	Ability
	Effect *string `json:"effect"`
}

func (a EnemyAbility) ToHashFields() []any {
	return []any{
		a.Ability.ID,
		derefOrNil(a.Effect),
	}
}

func (l *lookup) seedEnemyAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/enemy_abilities.json"

	var enemyAbilities []EnemyAbility

	err := loadJSONFile(string(srcPath), &enemyAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, enemyAbility := range enemyAbilities {
			var err error
			enemyAbility.Type = database.AbilityTypeEnemyAbility

			enemyAbility.Ability, err = seedObjAssignFK(qtx, enemyAbility.Ability, l.seedAbility)
			if err != nil {
				return err
			}

			err = qtx.CreateEnemyAbility(context.Background(), database.CreateEnemyAbilityParams{
				DataHash:  generateDataHash(enemyAbility),
				AbilityID: enemyAbility.Ability.ID,
				Effect:    getNullString(enemyAbility.Effect),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Enemy Ability: %s: %v", enemyAbility.Name, err)
			}
		}
		return nil
	})
}
