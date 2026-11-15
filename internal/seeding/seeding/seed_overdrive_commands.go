package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3SeedOverdriveCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := extractOverdriveCommands()
	if err != nil {
		return err
	}

	params := database.CreateOverdriveCommandBulkParams{
		DataHash:         make([]string, len(commands)),
		Name:             make([]string, len(commands)),
		Description:      make([]string, len(commands)),
		Rank:             make([]int32, len(commands)),
		TopmenuID:        make([]sql.NullInt32, len(commands)),
		SubmenuID:        make([]sql.NullInt32, len(commands)),
		CharacterClassID: make([]sql.NullInt32, len(commands)),
	}

	for i, oc := range commands {
		params.DataHash[i] = sh.GenerateDataHash(oc)
		params.Name[i] = oc.Name
		params.Description[i] = oc.Description
		params.Rank[i] = oc.Rank
		params.TopmenuID[i] = h.GetNullInt32(oc.TopmenuID)
		params.SubmenuID[i] = h.GetNullInt32(oc.SubmenuID)
		params.CharacterClassID[i] = h.GetNullInt32(oc.CharClassID)
	}

	dbRows, err := qtx.CreateOverdriveCommandBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create overdrive commands: %v", err)
	}

	for i, row := range dbRows {
		commands[i].ID = row.ID
		sh.RawJSON.OverdriveCommands[i].ID = row.ID
		lookup.L.OverdriveCommands[seedtypes.Key(commands[i])] = commands[i]
		lookup.L.OverdriveCommandsID[row.ID] = commands[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractOverdriveCommands() ([]seedtypes.OverdriveCommand, error) {
	commands := []seedtypes.OverdriveCommand{}
	var err error

	for i := range sh.RawJSON.OverdriveCommands {
		command := &sh.RawJSON.OverdriveCommands[i]

		command.TopmenuID, err = sh.AssignFKPtr(command.Topmenu, lookup.L.Topmenus)
		if err != nil {
			return nil, err
		}

		command.SubmenuID, err = sh.AssignFKPtr(&command.OpenSubmenu, lookup.L.Submenus)
		if err != nil {
			return nil, err
		}

		command.CharClassID, err = sh.AssignFKPtr(&command.User, lookup.L.CharClasses)
		if err != nil {
			return nil, err
		}

		commands = append(commands, *command)
	}

	return dedupeRows(commands, sh.Hashes), nil
}
