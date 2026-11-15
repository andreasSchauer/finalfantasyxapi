package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getMiscAbilityLearnedBy(ua seedtypes.MiscAbility) ([]seedtypes.CharacterClass, error) {
	return lookup.GetResources(ua.LearnedBy, lookup.L.CharClasses)
}

func MiscAbilitiesLearnedBy(qtx *database.Queries, ctx context.Context) error {
	const desc string = "misc abilities + learned by"
	jParams, err := processJunctions(desc, sh.RawJSON.MiscAbilities, getMiscAbilityLearnedBy)
	if err != nil {
		return err
	}

	return qtx.CreateMiscAbilitiesLearnedByJunctionBulk(ctx, database.CreateMiscAbilitiesLearnedByJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		MiscAbilityID:    jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}
