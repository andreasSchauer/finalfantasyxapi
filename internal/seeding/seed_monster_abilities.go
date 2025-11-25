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
		m.AbilityID,
		m.IsForced,
		m.IsUnused,
	}
}

func (m MonsterAbility) GetID() int32 {
	return m.ID
}

func (m MonsterAbility) Error() string {
	return fmt.Sprintf("monster ability %s-%v, type: %s, is forced: %t, is unused: %t", m.Name, h.DerefOrNil(m.Version), m.AbilityType, m.IsForced, m.IsUnused)
}

func (l *Lookup) seedMonsterAbility(qtx *database.Queries, monsterAbility MonsterAbility) (MonsterAbility, error) {
	var err error

	monsterAbility.AbilityID, err = assignFK(monsterAbility.AbilityReference, l.getAbility)
	if err != nil {
		return MonsterAbility{}, h.GetErr(monsterAbility.Error(), err)
	}

	dbMonsterAbility, err := qtx.CreateMonsterAbility(context.Background(), database.CreateMonsterAbilityParams{
		DataHash:  generateDataHash(monsterAbility),
		AbilityID: monsterAbility.AbilityID,
		IsForced:  monsterAbility.IsForced,
		IsUnused:  monsterAbility.IsUnused,
	})
	if err != nil {
		return MonsterAbility{}, h.GetErr(monsterAbility.Error(), err, "couldn't create monster ability")
	}

	monsterAbility.ID = dbMonsterAbility.ID

	return monsterAbility, nil
}
