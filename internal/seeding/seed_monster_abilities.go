package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterAbility struct {
	ID        int32
	AbilityID int32
	AbilityReference
	IsForced bool `json:"is_forced"`
	IsUnused bool `json:"is_unused"`
}

func (m MonsterAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.AbilityID,
		m.IsForced,
		m.IsUnused,
	}
}

func (m MonsterAbility) GetID() int32 {
	return m.ID
}

func (m MonsterAbility) Error() string {
	return fmt.Sprintf("monster ability '%s', type: %s, is forced: %t, is unused: %t", h.NameToString(m.Name, m.Version, nil), m.AbilityType, m.IsForced, m.IsUnused)
}

func (l *Lookup) seedMonsterAbility(qtx *database.Queries, monsterAbility MonsterAbility) (MonsterAbility, error) {
	var err error

	monsterAbility.AbilityID, err = assignFK(monsterAbility.AbilityReference, l.Abilities)
	if err != nil {
		return MonsterAbility{}, h.NewErr(monsterAbility.Error(), err)
	}

	dbMonsterAbility, err := qtx.CreateMonsterAbility(context.Background(), database.CreateMonsterAbilityParams{
		DataHash:  generateDataHash(monsterAbility),
		AbilityID: monsterAbility.AbilityID,
		IsForced:  monsterAbility.IsForced,
		IsUnused:  monsterAbility.IsUnused,
	})
	if err != nil {
		return MonsterAbility{}, h.NewErr(monsterAbility.Error(), err, "couldn't create monster ability")
	}

	monsterAbility.ID = dbMonsterAbility.ID

	return monsterAbility, nil
}



func (l *Lookup) loop3SeedMonsterAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractMonsterAbilities()
	if err != nil {
		return err
	}

	params := database.CreateMonsterAbilityBulkParams{
		DataHash: 	make([]string, len(abilities)),
		AbilityID: 	make([]int32, len(abilities)),
		IsForced: 	make([]bool, len(abilities)),
		IsUnused: 	make([]bool, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.AbilityID[i] = a.AbilityID
		params.IsForced[i] = a.IsForced
		params.IsUnused[i] = a.IsUnused
	}

	dbRows, err := qtx.CreateMonsterAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster abilities: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterAbilities() ([]MonsterAbility, error) {
	abilities := []MonsterAbility{}
	var err error

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		for j := range mon.Abilities {
			ability := &mon.Abilities[j]

			ability.AbilityID, err = assignFK(ability.AbilityReference, l.Abilities)
			if err != nil {
				return nil, err
			}

			abilities = append(abilities, *ability)
		}
	}

	return dedupeRows(abilities, l.Hashes), nil
}