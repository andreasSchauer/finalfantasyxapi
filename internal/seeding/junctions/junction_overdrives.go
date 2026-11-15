package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getOverdriveOverdriveAbilities(o seedtypes.Overdrive) ([]seedtypes.OverdriveAbility, error) {
	return lookup.TypedRefsToResources(o.OverdriveAbilities, lookup.L.OverdriveAbilities)
}

func SeedJuncOverdrivesOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrives + overdrive abilities"
	jParams, err := processJunctions(desc, sh.RawJSON.Overdrives, getOverdriveOverdriveAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateOverdrivesOverdriveAbilitiesJunctionBulk(ctx, database.CreateOverdrivesOverdriveAbilitiesJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		OverdriveID:        jParams.ParentIDs,
		OverdriveAbilityID: jParams.ChildIDs,
	})
}

func SeedJuncDefaultOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "default overdrive + overdrive abilities"
	params := database.CreateDefaultOverdriveAbilityJunctionBulkParams{
		DataHash:  make([]string, 0),
		ClassID:   make([]int32, 0),
		AbilityID: make([]int32, 0),
	}

	for _, overdrive := range sh.RawJSON.Overdrives {
		if overdrive.UnlockCondition != nil {
			continue
		}

		class, err := lookup.GetResource(overdrive.User, lookup.L.CharClasses)
		if err != nil {
			return err
		}

		for _, ref := range overdrive.OverdriveAbilities {
			oa, err := lookup.GetResource(ref.Untyped(), lookup.L.OverdriveAbilities)
			if err != nil {
				return err
			}

			j := StdJunction{
				ParentID: class.ID,
				ChildID:  oa.ID,
			}
			dataHash := GenerateJunctionHash(j, desc)

			params.DataHash = append(params.DataHash, dataHash)
			params.ClassID = append(params.ClassID, class.ID)
			params.AbilityID = append(params.AbilityID, oa.ID)
		}
	}

	return qtx.CreateDefaultOverdriveAbilityJunctionBulk(ctx, params)
}
