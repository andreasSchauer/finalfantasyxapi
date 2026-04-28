package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type OverdriveCommand struct {
	ID          int32
	CharClassID *int32
	TopmenuID	*int32
	SubmenuID   *int32
	Name        string `json:"name"`
	Description string `json:"description"`
	User        string `json:"user"`
	Rank        int32  `json:"rank"`
	Topmenu     *string `json:"topmenu"`
	OpenSubmenu string `json:"open_submenu"`
}

func (oc OverdriveCommand) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", oc),
		oc.Name,
		oc.Description,
		oc.Rank,
		h.DerefOrNil(oc.TopmenuID),
		oc.OpenSubmenu,
		h.DerefOrNil(oc.CharClassID),
		h.DerefOrNil(oc.SubmenuID),
	}
}

func (oc OverdriveCommand) GetID() int32 {
	return oc.ID
}

func (oc OverdriveCommand) Error() string {
	return fmt.Sprintf("overdrive command %s", oc.Name)
}

func (o OverdriveCommand) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   o.ID,
		Name: o.Name,
	}
}

func (l *Lookup) seedOverdriveCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrive_commands.json"

	var overdriveCommands []OverdriveCommand
	err := loadJSONFile(string(srcPath), &overdriveCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range overdriveCommands {
			dbODCommand, err := qtx.CreateOverdriveCommand(context.Background(), database.CreateOverdriveCommandParams{
				DataHash:    generateDataHash(command),
				Name:        command.Name,
				Description: command.Description,
				Rank:        command.Rank,
			})
			if err != nil {
				return h.NewErr(command.Error(), err, "couldn't create overdrive command")
			}

			command.ID = dbODCommand.ID
			l.OverdriveCommands[command.Name] = command
			l.OverdriveCommandsID[command.ID] = command
		}

		return nil
	})
}

func (l *Lookup) seedOverdriveCommandsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrive_commands.json"

	var overdriveCommands []OverdriveCommand
	err := loadJSONFile(string(srcPath), &overdriveCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCommand := range overdriveCommands {
			command, err := GetResource(jsonCommand.Name, l.OverdriveCommands)
			if err != nil {
				return err
			}

			command.TopmenuID, err = assignFKPtr(command.Topmenu, l.Topmenus)
			if err != nil {
				return h.NewErr(command.Error(), err)
			}

			command.CharClassID, err = assignFKPtr(&command.User, l.CharClasses)
			if err != nil {
				return h.NewErr(command.Error(), err)
			}

			command.SubmenuID, err = assignFKPtr(&command.OpenSubmenu, l.Submenus)
			if err != nil {
				return h.NewErr(command.Error(), err)
			}

			err = qtx.UpdateOverdriveCommand(context.Background(), database.UpdateOverdriveCommandParams{
				DataHash:         generateDataHash(command),
				CharacterClassID: h.GetNullInt32(command.CharClassID),
				TopmenuID: 		  h.GetNullInt32(command.TopmenuID),
				SubmenuID:        h.GetNullInt32(command.SubmenuID),
				ID:               command.ID,
			})
			if err != nil {
				return h.NewErr(command.Error(), err, "couldn't update overdrive command")
			}
		}

		return nil
	})
}



func (l *Lookup) loop3SeedOverdriveCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := l.extractOverdriveCommands()
	if err != nil {
		return err
	}

	params := database.CreateOverdriveCommandBulkParams{
		DataHash:   		make([]string, len(commands)),
		Name:      			make([]string, len(commands)),
		Description:    	make([]string, len(commands)),
		Rank: 				make([]int32, len(commands)),
		TopmenuID: 			make([]sql.NullInt32, len(commands)),
		SubmenuID: 			make([]sql.NullInt32, len(commands)),
		CharacterClassID: 	make([]sql.NullInt32, len(commands)),
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