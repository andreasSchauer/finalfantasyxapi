package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type DefaultAbilitiesEntry struct {
	Name             string             `json:"name"`
	DefaultAbilities []AbilityReference `json:"default_abilities"`
}

func (l *Lookup) seedJuncDefaultAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "default abilities"
	params := database.CreateDefaultAbilityJunctionBulkParams{
		DataHash: 	make([]string, 0),
		ClassID: 	make([]int32, 0),
		AbilityID: 	make([]int32, 0),
	}

	for _, entry := range l.json.defaultAbilities {
		class, err := GetResource(entry.Name, l.CharClasses)
		if err != nil {
			return err
		}

		for _, ref := range entry.DefaultAbilities {
			ability, err := GetResource(ref, l.Abilities)
			if err != nil {
				return err
			}

			j := StdJunction{
				ParentID: 	class.ID,
				ChildID: 	ability.ID,
			}
			dataHash := generateJunctionHash(j, desc)

			params.DataHash = append(params.DataHash, dataHash)
			params.ClassID = append(params.ClassID, class.ID)
			params.AbilityID = append(params.AbilityID, ability.ID)
		}
	}

	return qtx.CreateDefaultAbilityJunctionBulk(ctx, params)
}