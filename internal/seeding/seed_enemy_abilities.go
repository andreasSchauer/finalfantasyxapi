package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EnemyAbility struct {
	ID int32
	Ability
	Effect             *string             `json:"effect"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (a EnemyAbility) ToHashFields() []any {
	return []any{
		a.Ability.ID,
		h.DerefOrNil(a.Effect),
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
	return fmt.Sprintf("enemy ability %s, version %v", a.Name, h.DerefOrNil(a.Version))
}

func (l *Lookup) seedEnemyAbilities(db *database.Queries, dbConn *sql.DB) error {
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
				return h.GetErr(enemyAbility.Error(), err)
			}

			dbEnemyAbility, err := qtx.CreateEnemyAbility(context.Background(), database.CreateEnemyAbilityParams{
				DataHash:  generateDataHash(enemyAbility),
				AbilityID: enemyAbility.Ability.ID,
				Effect:    h.GetNullString(enemyAbility.Effect),
			})
			if err != nil {
				return h.GetErr(enemyAbility.Error(), err, "couldn't create enemy ability")
			}

			enemyAbility.ID = dbEnemyAbility.ID
			key := CreateLookupKey(enemyAbility.Ability)
			l.EnemyAbilities[key] = enemyAbility
		}
		return nil
	})
}

func (l *Lookup) seedEnemyAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/enemy_abilities.json"

	var enemyAbilities []EnemyAbility

	err := loadJSONFile(string(srcPath), &enemyAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range enemyAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			enemyAbility, err := GetResource(abilityRef, l.EnemyAbilities)
			if err != nil {
				return err
			}

			l.currentAbility = enemyAbility.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, enemyAbility.BattleInteractions)
			if err != nil {
				return h.GetErr(enemyAbility.Error(), err)
			}
		}

		return nil
	})
}
