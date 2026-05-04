package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

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
