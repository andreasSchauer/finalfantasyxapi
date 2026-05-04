package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AeonCommand struct {
	ID                int32
	TopmenuID         *int32
	SubmenuID         *int32
	Name              string                `json:"name"`
	Description       string                `json:"description"`
	Effect            string                `json:"effect"`
	Topmenu           *string               `json:"topmenu"`
	OpenSubmenu       *string               `json:"open_submenu"`
	Cursor            *string               `json:"cursor"`
	PossibleAbilities []PossibleAbilityList `json:"possible_abilities"`
}

func (c AeonCommand) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", c),
		c.Name,
		c.Description,
		c.Effect,
		h.DerefOrNil(c.TopmenuID),
		h.DerefOrNil(c.Cursor),
		h.DerefOrNil(c.SubmenuID),
	}
}

func (c AeonCommand) GetID() int32 {
	return c.ID
}

func (c AeonCommand) Error() string {
	return fmt.Sprintf("aeon command %s", c.Name)
}

func (c AeonCommand) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   c.ID,
		Name: c.Name,
	}
}

type PossibleAbilityList struct {
	User      string             `json:"user"`
	Abilities []AbilityReference `json:"abilities"`
}

func (pa PossibleAbilityList) Error() string {
	return fmt.Sprintf("possible abilities for %s", pa.User)
}

func (l *Lookup) seedAeonCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/aeon_commands.json"

	var aeonCommands []AeonCommand
	err := loadJSONFile(string(srcPath), &aeonCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range aeonCommands {
			dbAeonCommand, err := qtx.CreateAeonCommand(context.Background(), database.CreateAeonCommandParams{
				DataHash:    generateDataHash(command),
				Name:        command.Name,
				Description: command.Description,
				Effect:      command.Effect,
				Cursor:      database.ToNullTargetType(command.Cursor),
			})
			if err != nil {
				return h.NewErr(command.Error(), err, "couldn't create aeon command")
			}

			command.ID = dbAeonCommand.ID
			l.AeonCommands[command.Name] = command
			l.AeonCommandsID[command.ID] = command
		}
		return nil
	})
}

func (l *Lookup) seedAeonCommandsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/aeon_commands.json"

	var aeonCommands []AeonCommand
	err := loadJSONFile(string(srcPath), &aeonCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCommand := range aeonCommands {
			command, err := GetResource(jsonCommand.Name, l.AeonCommands)
			if err != nil {
				return err
			}

			command.TopmenuID, err = assignFKPtr(command.Topmenu, l.Topmenus)
			if err != nil {
				return h.NewErr(command.Error(), err)
			}

			command.SubmenuID, err = assignFKPtr(command.OpenSubmenu, l.Submenus)
			if err != nil {
				return h.NewErr(command.Error(), err)
			}

			err = qtx.UpdateAeonCommand(context.Background(), database.UpdateAeonCommandParams{
				DataHash:  generateDataHash(command),
				TopmenuID: h.GetNullInt32(command.TopmenuID),
				SubmenuID: h.GetNullInt32(command.SubmenuID),
				ID:        command.ID,
			})
			if err != nil {
				return h.NewErr(command.Error(), err, "couldn't update aeon command")
			}

			err = l.seedAeonCommandPossibleAbilities(qtx, command)
			if err != nil {
				return h.NewErr(command.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedAeonCommandPossibleAbilities(qtx *database.Queries, command AeonCommand) error {
	for _, possibleAbility := range command.PossibleAbilities {
		for _, abilityRef := range possibleAbility.Abilities {
			var err error
			charClass, err := GetResource(possibleAbility.User, l.CharClasses)
			if err != nil {
				return err
			}

			threeWay, err := createThreeWayJunction(command, charClass, abilityRef, l.Abilities)
			if err != nil {
				return h.NewErr(charClass.Error(), err)
			}

			err = qtx.CreateAeonCommandsPossibleAbilitiesJunction(context.Background(), database.CreateAeonCommandsPossibleAbilitiesJunctionParams{
				DataHash:         generateDataHash(threeWay),
				AeonCommandID:    threeWay.GrandparentID,
				CharacterClassID: threeWay.ParentID,
				AbilityID:        threeWay.ChildID,
			})
			if err != nil {
				return h.NewErr(abilityRef.Error(), err, "couldn't junction possible ability")
			}
		}
	}

	return nil
}



func (l *Lookup) loop3SeedAeonCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := l.extractAeonCommands()
	if err != nil {
		return err
	}

	params := database.CreateAeonCommandBulkParams{
		DataHash:   	make([]string, len(commands)),
		Name:      		make([]string, len(commands)),
		Description:    make([]string, len(commands)),
		Effect: 		make([]string, len(commands)),
		Cursor: 		make([]database.NullTargetType, len(commands)),
		TopmenuID: 		make([]sql.NullInt32, len(commands)),
		SubmenuID: 		make([]sql.NullInt32, len(commands)),
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

func (l *Lookup) processAeonCommandsPossibleAbilities(desc string) (JunctionParams, error) {
	params := JunctionParams{
		DataHashes: 	make([]string, 0),
		GrandParentIDs: make([]int32, 0),
		ParentIDs:  	make([]int32, 0),
		ChildIDs:   	make([]int32, 0),
	}

	for _, command := range l.json.aeonCommands {
		for _, list := range command.PossibleAbilities {
			class, err := GetResource(list.User, l.CharClasses)
			if err != nil {
				return JunctionParams{}, err
			}

			for _, ref := range list.Abilities {
				ability, err := GetResource(ref, l.Abilities)
				if err != nil {
					return JunctionParams{}, err
				}

				j := ThreeWayJunction{}
				j.GrandparentID = command.ID
				j.ParentID = class.ID
				j.ChildID = ability.ID
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

func (l *Lookup) seedJuncAeonCommandsPossibleAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeon commands + possible abilities"
	jParams, err := l.processAeonCommandsPossibleAbilities(desc)
	if err != nil {
		return err
	}

	return qtx.CreateAeonCommandsPossibleAbilitiesJunctionBulk(ctx, database.CreateAeonCommandsPossibleAbilitiesJunctionBulkParams{
		DataHash:   		jParams.DataHashes,
		AeonCommandID: 		jParams.GrandParentIDs,
		CharacterClassID: 	jParams.ParentIDs,
		AbilityID:  		jParams.ChildIDs,
	})
}