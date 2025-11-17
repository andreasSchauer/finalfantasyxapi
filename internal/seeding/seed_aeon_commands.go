package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type AeonCommand struct {
	ID                int32
	SubmenuID         *int32
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Effect            string            `json:"effect"`
	Topmenu           string            `json:"topmenu"`
	OpenSubmenu       *string           `json:"open_submenu"`
	Cursor            *string           `json:"cursor"`
	PossibleAbilities []PossibleAbility `json:"possible_abilities"`
}

func (c AeonCommand) ToHashFields() []any {
	return []any{
		c.Name,
		c.Description,
		c.Effect,
		c.Topmenu,
		derefOrNil(c.Cursor),
		derefOrNil(c.SubmenuID),
	}
}

func (c AeonCommand) GetID() int32 {
	return c.ID
}

func (c AeonCommand) Error() string {
	return fmt.Sprintf("aeon command %s", c.Name)
}

type PossibleAbility struct {
	User      string             `json:"user"`
	Abilities []AbilityReference `json:"abilities"`
}

func (pa PossibleAbility) Error() string {
	return fmt.Sprintf("possible abilities for %s", pa.User)
}

func (l *Lookup) seedAeonCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/aeon_commands.json"

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
				Topmenu:     database.TopmenuType(command.Topmenu),
				Cursor:      nullTargetType(command.Cursor),
			})
			if err != nil {
				return getErr(command.Error(), err, "couldn't create aeon command")
			}

			command.ID = dbAeonCommand.ID
			l.aeonCommands[command.Name] = command
		}
		return nil
	})
}

func (l *Lookup) seedAeonCommandsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/aeon_commands.json"

	var aeonCommands []AeonCommand
	err := loadJSONFile(string(srcPath), &aeonCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCommand := range aeonCommands {
			command, err := l.getAeonCommand(jsonCommand.Name)
			if err != nil {
				return err
			}

			command.SubmenuID, err = assignFKPtr(command.OpenSubmenu, l.getSubmenu)
			if err != nil {
				return getErr(command.Error(), err)
			}

			err = qtx.UpdateAeonCommand(context.Background(), database.UpdateAeonCommandParams{
				DataHash:  generateDataHash(command),
				SubmenuID: getNullInt32(command.SubmenuID),
				ID:        command.ID,
			})
			if err != nil {
				return getErr(command.Error(), err, "couldn't update aeon command")
			}

			err = l.seedAeonCommandPossibleAbilities(qtx, command)
			if err != nil {
				return getErr(command.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedAeonCommandPossibleAbilities(qtx *database.Queries, command AeonCommand) error {
	for _, possibleAbility := range command.PossibleAbilities {
		for _, abilityRef := range possibleAbility.Abilities {
			var err error
			charClass, err := l.getCharacterClass(possibleAbility.User)
			if err != nil {
				return err
			}

			threeWay, err := createThreeWayJunction(command, charClass, abilityRef, l.getAbility)
			if err != nil {
				return getErr(charClass.Error(), err)
			}

			err = qtx.CreateAeonCommandsPossibleAbilitiesJunction(context.Background(), database.CreateAeonCommandsPossibleAbilitiesJunctionParams{
				DataHash:         generateDataHash(threeWay),
				AeonCommandID:    threeWay.GrandparentID,
				CharacterClassID: threeWay.ParentID,
				AbilityID:        threeWay.ChildID,
			})
			if err != nil {
				return getErr(abilityRef.Error(), err, "couldn't junction possible ability")
			}
		}
	}

	return nil
}
