package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedAeonCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := l.extractAeonCommands()
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
		l.json.aeonCommands[i].ID = row.ID
		l.AeonCommands[commands[i].Name] = commands[i]
		l.AeonCommandsID[row.ID] = commands[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAeonCommands() ([]AeonCommand, error) {
	commands := []AeonCommand{}
	var err error

	for i := range l.json.aeonCommands {
		command := &l.json.aeonCommands[i]

		command.TopmenuID, err = assignFKPtr(command.Topmenu, l.Topmenus)
		if err != nil {
			return nil, err
		}

		command.SubmenuID, err = assignFKPtr(command.OpenSubmenu, l.Submenus)
		if err != nil {
			return nil, err
		}

		commands = append(commands, *command)
	}

	return dedupeRows(commands, l.Hashes), nil
}
