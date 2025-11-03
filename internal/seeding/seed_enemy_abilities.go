package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type EnemyAbility struct {
	ID					int32
	Ability
	Effect 				*string 			`json:"effect"`
	BattleInteractions  []BattleInteraction `json:"battle_interactions"`
}

func (a EnemyAbility) ToHashFields() []any {
	return []any{
		a.Ability.ID,
		derefOrNil(a.Effect),
	}
}


func (a EnemyAbility) GetID() int32 {
	return a.ID
}

func (a EnemyAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        a.Name,
		Version:     a.Version,
		AbilityType: string(database.AbilityTypeEnemyAbility),
	}
}

func (a EnemyAbility) Error() string {
	return fmt.Sprintf("enemy ability %s, version %v", a.Name, derefOrNil(a.Version))
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

			enemyAbility.Ability, err = seedObjAssignID(qtx, enemyAbility.Ability, l.seedAbility)
			if err != nil {
				return err
			}

			dbEnemyAbility, err := qtx.CreateEnemyAbility(context.Background(), database.CreateEnemyAbilityParams{
				DataHash:  generateDataHash(enemyAbility),
				AbilityID: enemyAbility.Ability.ID,
				Effect:    getNullString(enemyAbility.Effect),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Enemy Ability: %s: %v", enemyAbility.Name, err)
			}

			enemyAbility.ID = dbEnemyAbility.ID
			key := createLookupKey(enemyAbility.Ability)
			l.enemyAbilities[key] = enemyAbility
		}
		return nil
	})
}


func (l *lookup) seedEnemyAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/enemy_abilities.json"

	var enemyAbilities []EnemyAbility

	err := loadJSONFile(string(srcPath), &enemyAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range enemyAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			ability, err := l.getEnemyAbility(abilityRef)
			if err != nil {
				return err
			}

			l.currentAbility = ability.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, ability.BattleInteractions)
			if err != nil {
				return err
			}
		}

		return nil
	})
}