package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type AeonCommand struct {
	ID					int32
	SubmenuID			*int32
	Name        		string  			`json:"name"`
	Description 		string  			`json:"description"`
	Effect      		string  			`json:"effect"`
	Topmenu				string				`json:"topmenu"`
	OpenSubmenu			*string				`json:"open_submenu"`
	Cursor      		*string 			`json:"cursor"`
	PossibleAbilities 	[]PossibleAbility 	`json:"possible_abilities"`
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

type PossibleAbility struct {
	User		string				`json:"user"`
	Abilities	[]AbilityReference	`json:"abilities"`
}

type PossibleAbilityJunction struct {
	Junction
	ClassID		int32
}

func (j PossibleAbilityJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		j.ClassID,
	}
}


func (l *lookup) seedAeonCommands(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/aeon_commands.json"

	var aeonCommands []AeonCommand
	err := loadJSONFile(string(srcPath), &aeonCommands)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, command := range aeonCommands {
			dbAeonCommand, err := qtx.CreateAeonCommand(context.Background(), database.CreateAeonCommandParams{
				DataHash:    	generateDataHash(command),
				Name:        	command.Name,
				Description: 	command.Description,
				Effect:      	command.Effect,
				Topmenu: 		database.TopmenuType(command.Topmenu),
				Cursor:      	nullTargetType(command.Cursor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Aeon Command: %s: %v", command.Name, err)
			}

			command.ID = dbAeonCommand.ID
			l.aeonCommands[command.Name] = command
		}
		return nil
	})
}


func (l *lookup) createAeonCommandsRelationships(db *database.Queries, dbConn *sql.DB) error {
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
				return err
			}

			qtx.UpdateAeonCommand(context.Background(), database.UpdateAeonCommandParams{
				DataHash:    	generateDataHash(command),
				SubmenuID: 		getNullInt32(command.SubmenuID),
				ID: 			command.ID,
			})

			err = l.seedAeonCommandPossibleAbilities(qtx, command)
			if err != nil {
				return err
			}
		}

		return nil
	})
}


func (l *lookup) seedAeonCommandPossibleAbilities(qtx *database.Queries, command AeonCommand) error {
	for _, possibleAbility := range command.PossibleAbilities {
		for _, abilityRef := range possibleAbility.Abilities {
			var err error
			paJunction := PossibleAbilityJunction{}
			
			paJunction.ClassID, err = assignFK(possibleAbility.User, l.getCharacterClass)
			if err != nil {
				return err
			}

			paJunction.Junction, err = createJunction(command, abilityRef, l.getAbility)
			if err != nil {
				return err
			}

			err = qtx.CreateAeonCommandAbilityJunction(context.Background(), database.CreateAeonCommandAbilityJunctionParams{
				DataHash: 			generateDataHash(paJunction),
				AeonCommandID: 		paJunction.ParentID,
				AbilityID: 			paJunction.ChildID,
				CharacterClassID: 	paJunction.ClassID,
			})
			if err != nil {
				return fmt.Errorf("couldn't create junction between aeon command %s, ability %s, and character class %s: %v", command.Name, abilityRef.Name, possibleAbility.User, err)
			}
		}
	}

	return nil
}