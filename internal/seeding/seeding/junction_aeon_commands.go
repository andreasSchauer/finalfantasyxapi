package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func processAeonCommandsPossibleAbilities(desc string) (JunctionParams, error) {
	params := JunctionParams{
		DataHashes:     make([]string, 0),
		GrandParentIDs: make([]int32, 0),
		ParentIDs:      make([]int32, 0),
		ChildIDs:       make([]int32, 0),
	}

	for _, command := range rawJSON.aeonCommands {
		for _, list := range command.PossibleAbilities {
			class, err := lookup.GetResource(list.User, lookup.L.CharClasses)
			if err != nil {
				return JunctionParams{}, err
			}

			for _, ref := range list.Abilities {
				ability, err := lookup.GetResource(ref, lookup.L.Abilities)
				if err != nil {
					return JunctionParams{}, err
				}

				j := ThreeWayJunction{
					GrandparentID: command.ID,
					ParentID: class.ID,
					ChildID: ability.ID,
				}
				dataHash := generateJunctionHash(j, desc)

				params.DataHashes = append(params.DataHashes, dataHash)
				params.GrandParentIDs = append(params.GrandParentIDs, command.ID)
				params.ParentIDs = append(params.ParentIDs, class.ID)
				params.ChildIDs = append(params.ChildIDs, ability.ID)
			}
		}
	}

	return params, nil
}

func seedJuncAeonCommandsPossibleAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeon commands + possible abilities"
	jParams, err := processAeonCommandsPossibleAbilities(desc)
	if err != nil {
		return err
	}

	return qtx.CreateAeonCommandsPossibleAbilitiesJunctionBulk(ctx, database.CreateAeonCommandsPossibleAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AeonCommandID:    jParams.GrandParentIDs,
		CharacterClassID: jParams.ParentIDs,
		AbilityID:        jParams.ChildIDs,
	})
}
