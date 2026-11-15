package junctions

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getSubmenuUsers(s seedtypes.Submenu) ([]seedtypes.CharacterClass, error) {
	return lookup.GetResources(s.Users, lookup.L.CharClasses)
}

func SubmenusUsers(qtx *database.Queries, ctx context.Context) error {
	const desc string = "submenus + users"
	jParams, err := processJunctions(desc, sh.RawJSON.Submenus, getSubmenuUsers)
	if err != nil {
		return err
	}

	return qtx.CreateSubmenusUsersJunctionBulk(ctx, database.CreateSubmenusUsersJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		SubmenuID:        jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}
