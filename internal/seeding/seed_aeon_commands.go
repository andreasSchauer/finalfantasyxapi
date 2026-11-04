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

func loop3SeedAeonCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := extractAeonCommands()
	if err != nil {
		return err
	}

	params := database.CreateAeonCommandBulkParams{
		DataHash:    make([]string, len(commands)),
		Name:        make([]string, len(commands)),
		Description: make([]string, len(commands)),
		Effect:      make([]string, len(commands)),
		Cursor:      make([]database.NullTargetType, len(commands)),
		TopmenuID:   make([]sql.NullInt32, len(commands)),
		SubmenuID:   make([]sql.NullInt32, len(commands)),
	}

	for i, ac := range commands {
		params.DataHash[i] = generateDataHash(ac)
		params.Name[i] = ac.Name
		params.Description[i] = ac.Description
		params.Effect[i] = ac.Effect
		params.Cursor[i] = database.ToNullTargetType(ac.Cursor)
		params.TopmenuID[i] = h.GetNullInt32(ac.TopmenuID)
		params.SubmenuID[i] = h.GetNullInt32(ac.SubmenuID)
	}

	dbRows, err := qtx.CreateAeonCommandBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create aeon commands: %v", err)
	}

	for i, row := range dbRows {
		commands[i].ID = row.ID
		rawJSON.aeonCommands[i].ID = row.ID
		lookup.L.AeonCommands[seedtypes.Key(commands[i])] = commands[i]
		lookup.L.AeonCommandsID[row.ID] = commands[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractAeonCommands() ([]seedtypes.AeonCommand, error) {
	commands := []seedtypes.AeonCommand{}
	var err error

	for i := range rawJSON.aeonCommands {
		command := &rawJSON.aeonCommands[i]

		command.TopmenuID, err = assignFKPtr(command.Topmenu, lookup.L.Topmenus)
		if err != nil {
			return nil, err
		}

		command.SubmenuID, err = assignFKPtr(command.OpenSubmenu, lookup.L.Submenus)
		if err != nil {
			return nil, err
		}

		commands = append(commands, *command)
	}

	return dedupeRows(commands, hashes), nil
}
