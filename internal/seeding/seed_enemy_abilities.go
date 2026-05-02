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

func (e EnemyAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
		e.Ability.ID,
		h.DerefOrNil(e.Effect),
	}
}

func (e EnemyAbility) ToKeyFields() []any {
	return []any{
		e.Ability.Name,
		h.DerefOrNil(e.Ability.Version),
	}
}

func (e EnemyAbility) GetID() int32 {
	return e.ID
}

func (e EnemyAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        e.Name,
		Version:     e.Version,
		AbilityType: string(database.AbilityTypeEnemyAbility),
	}
}

func (e EnemyAbility) Error() string {
	return fmt.Sprintf("enemy ability '%s'", h.NameToString(e.Name, e.Version, e.Specification))
}

func (e EnemyAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            e.ID,
		Name:          e.Name,
		Version:       e.Version,
		Specification: e.Specification,
	}
}

func (l *Lookup) seedEnemyAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/enemy_abilities.json"

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
				return h.NewErr(enemyAbility.Error(), err)
			}

			dbEnemyAbility, err := qtx.CreateEnemyAbility(context.Background(), database.CreateEnemyAbilityParams{
				DataHash:  generateDataHash(enemyAbility),
				AbilityID: enemyAbility.Ability.ID,
				Effect:    h.GetNullString(enemyAbility.Effect),
			})
			if err != nil {
				return h.NewErr(enemyAbility.Error(), err, "couldn't create enemy ability")
			}

			enemyAbility.ID = dbEnemyAbility.ID
			key := Key(enemyAbility)
			l.EnemyAbilities[key] = enemyAbility
			l.EnemyAbilitiesID[enemyAbility.ID] = enemyAbility
		}
		return nil
	})
}

func (l *Lookup) seedEnemyAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/enemy_abilities.json"

	var enemyAbilities []EnemyAbility

	err := loadJSONFile(string(srcPath), &enemyAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range enemyAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			enemyAbility, err := GetResource(abilityRef.Untyped(), l.EnemyAbilities)
			if err != nil {
				return err
			}

			l.currentAbility = enemyAbility.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, enemyAbility.BattleInteractions)
			if err != nil {
				return h.NewErr(enemyAbility.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) loop3SeedEnemyAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractEnemyAbilities()
	if err != nil {
		return err
	}

	params := database.CreateEnemyAbilityBulkParams{
		DataHash:  make([]string, len(abilities)),
		AbilityID: make([]int32, len(abilities)),
		Effect:    make([]sql.NullString, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
		params.Effect[i] = h.GetNullString(a.Effect)
	}

	dbRows, err := qtx.CreateEnemyAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create enemy abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		l.json.enemyAbilities[i].ID = row.ID
		key := Key(abilities[i])
		l.EnemyAbilities[key] = abilities[i]
		l.EnemyAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractEnemyAbilities() ([]EnemyAbility, error) {
	abilities := []EnemyAbility{}
	var err error

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		ability.Ability.ID, err = l.getHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) completeEnemyAbilities() error {
	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		err := l.completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		l.EnemyAbilities[Key(ability)] = *ability
		l.EnemyAbilitiesID[ability.ID] = *ability
	}

	return nil
}