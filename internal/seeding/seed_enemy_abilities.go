package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type EnemyAbility struct {
	Ability
	AbilityAttributes
	AbilityID			int32
	Effect				*string		`json:"effect"`
}


func(a EnemyAbility) ToHashFields() []any {
	return []any{
		a.AbilityID,
		derefOrNil(a.Effect),
	}
}


func seedEnemyAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/enemy_abilities.json"

	var enemyAbilities []EnemyAbility

	err := loadJSONFile(string(srcPath), &enemyAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, enemyAbility := range enemyAbilities {
			ability := enemyAbility.Ability
			attributes := enemyAbility.AbilityAttributes
			ability.Type = database.AbilityTypeEnemyAbility

			dbAbility, err := seedAbility(qtx, attributes, ability)
			if err != nil {
				return err
			}
			
			enemyAbility.AbilityID = dbAbility.ID

			err = qtx.CreateEnemyAbility(context.Background(), database.CreateEnemyAbilityParams{
				DataHash: 		generateDataHash(enemyAbility),
				AbilityID: 		enemyAbility.AbilityID,
				Effect: 		getNullString(enemyAbility.Effect),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Enemy Ability: %s: %v", ability.Name, err)
			}
		}
		return nil
	})
}