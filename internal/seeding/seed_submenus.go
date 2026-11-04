package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop2SeedSubmenus(qtx *database.Queries, ctx context.Context) error {
	submenus, err := extractSubmenus()
	if err != nil {
		return err
	}

	params := database.CreateSubmenuBulkParams{
		DataHash:    make([]string, len(submenus)),
		Name:        make([]string, len(submenus)),
		TopmenuID:   make([]sql.NullInt32, len(submenus)),
		Description: make([]sql.NullString, len(submenus)),
		Effect:      make([]string, len(submenus)),
	}

	for i, s := range submenus {
		params.DataHash[i] = generateDataHash(s)
		params.Name[i] = s.Name
		params.TopmenuID[i] = h.GetNullInt32(s.TopmenuID)
		params.Description[i] = h.GetNullString(s.Description)
		params.Effect[i] = s.Effect
	}

	dbRows, err := qtx.CreateSubmenuBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create submenus: %v", err)
	}

	for i, row := range dbRows {
		submenus[i].ID = row.ID
		rawJSON.submenus[i].ID = row.ID
		lookup.L.Submenus[seedtypes.Key(submenus[i])] = submenus[i]
		lookup.L.SubmenusID[row.ID] = submenus[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractSubmenus() ([]seedtypes.Submenu, error) {
	submenus := []seedtypes.Submenu{}
	var err error

	for i := range rawJSON.submenus {
		submenu := &rawJSON.submenus[i]

		if submenu.Topmenu != nil {
			submenu.TopmenuID, err = assignFKPtr(submenu.Topmenu, lookup.L.Topmenus)
			if err != nil {
				return nil, err
			}
		}
		submenus = append(submenus, *submenu)
	}

	return dedupeRows(submenus, hashes), nil
}

func getSubmenuUsers(s seedtypes.Submenu) ([]seedtypes.CharacterClass, error) {
	return lookup.GetResources(s.Users, lookup.L.CharClasses)
}

func seedJuncSubmenusUsers(qtx *database.Queries, ctx context.Context) error {
	const desc string = "submenus + users"
	jParams, err := processJunctions(desc, rawJSON.submenus, getSubmenuUsers)
	if err != nil {
		return err
	}

	return qtx.CreateSubmenusUsersJunctionBulk(ctx, database.CreateSubmenusUsersJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		SubmenuID:        jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}
