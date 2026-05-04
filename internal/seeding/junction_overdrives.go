package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getOverdriveOverdriveAbilities(o Overdrive) ([]OverdriveAbility, error) {
	return typedRefsToResources(o.OverdriveAbilities, l.OverdriveAbilities)
}

func (l *Lookup) seedJuncOverdrivesOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrives + overdrive abilities"
	jParams, err := processJunctions(l, desc, l.json.overdrives, l.getOverdriveOverdriveAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateOverdrivesOverdriveAbilitiesJunctionBulk(ctx, database.CreateOverdrivesOverdriveAbilitiesJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		OverdriveID:        jParams.ParentIDs,
		OverdriveAbilityID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncDefaultOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "default overdrive + overdrive abilities"
	params := database.CreateDefaultOverdriveAbilityJunctionBulkParams{
		DataHash:  make([]string, 0),
		ClassID:   make([]int32, 0),
		AbilityID: make([]int32, 0),
	}

	for _, overdrive := range l.json.overdrives {
		if overdrive.UnlockCondition != nil {
			continue
		}

		class, err := GetResource(overdrive.User, l.CharClasses)
		if err != nil {
			return err
		}

		for _, ref := range overdrive.OverdriveAbilities {
			oa, err := GetResource(ref.Untyped(), l.OverdriveAbilities)
			if err != nil {
				return err
			}

			j := StdJunction{
				ParentID: class.ID,
				ChildID:  oa.ID,
			}
			dataHash := generateJunctionHash(j, desc)

			params.DataHash = append(params.DataHash, dataHash)
			params.ClassID = append(params.ClassID, class.ID)
			params.AbilityID = append(params.AbilityID, oa.ID)
		}
	}

	return qtx.CreateDefaultOverdriveAbilityJunctionBulk(ctx, params)
}
