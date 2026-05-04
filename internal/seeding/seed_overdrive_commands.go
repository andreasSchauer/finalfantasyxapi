package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedOverdriveCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := l.extractOverdriveCommands()
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
		params.DataHash[i] = generateDataHash(oc)
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
		l.json.overdriveCommands[i].ID = row.ID
		l.OverdriveCommands[commands[i].Name] = commands[i]
		l.OverdriveCommandsID[row.ID] = commands[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractOverdriveCommands() ([]OverdriveCommand, error) {
	commands := []OverdriveCommand{}
	var err error

	for i := range l.json.overdriveCommands {
		command := &l.json.overdriveCommands[i]

		command.TopmenuID, err = assignFKPtr(command.Topmenu, l.Topmenus)
		if err != nil {
			return nil, err
		}

		command.SubmenuID, err = assignFKPtr(&command.OpenSubmenu, l.Submenus)
		if err != nil {
			return nil, err
		}

		command.CharClassID, err = assignFKPtr(&command.User, l.CharClasses)
		if err != nil {
			return nil, err
		}

		commands = append(commands, *command)
	}

	return dedupeRows(commands, l.Hashes), nil
}
