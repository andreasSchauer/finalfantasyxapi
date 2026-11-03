package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop4SeedFormationTriggerCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := extractFormationTriggerCommands()
	if err != nil {
		return err
	}

	params := database.CreateFormationTriggerCommandBulkParams{
		DataHash:         make([]string, len(commands)),
		TriggerCommandID: make([]int32, len(commands)),
		Condition:        make([]sql.NullString, len(commands)),
		UseAmount:        make([]sql.NullInt32, len(commands)),
	}

	for i, c := range commands {
		params.DataHash[i] = generateDataHash(c)
		params.TriggerCommandID[i] = c.TriggerCommandID
		params.Condition[i] = h.GetNullString(c.Condition)
		params.UseAmount[i] = h.GetNullInt32(c.UseAmount)
	}

	dbRows, err := qtx.CreateFormationTriggerCommandBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create formation trigger commands: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractFormationTriggerCommands() ([]seedtypes.FormationTriggerCommand, error) {
	commands := []seedtypes.FormationTriggerCommand{}
	var err error

	for i := range rawJSON.monsterFormations {
		mf := &rawJSON.monsterFormations[i]

		for j := range mf.TriggerCommands {
			command := &mf.TriggerCommands[j]

			command.TriggerCommandID, err = assignFK(command.AbilityReference.Untyped(), L.TriggerCommands)
			if err != nil {
				return nil, err
			}

			commands = append(commands, *command)
		}
	}

	return dedupeRows(commands, hashes), nil
}

func getFormationTriggerCommands() []seedtypes.FormationTriggerCommand {
	commands := []seedtypes.FormationTriggerCommand{}

	for _, formation := range rawJSON.monsterFormations {
		commands = append(commands, formation.TriggerCommands...)
	}

	return commands
}

func getFormationTriggerCommandUsers(tc seedtypes.FormationTriggerCommand) ([]seedtypes.CharacterClass, error) {
	return GetResources(tc.Users, L.CharClasses)
}

func seedJuncFormationTriggerCommandsUsers(qtx *database.Queries, ctx context.Context) error {
	const desc string = "formation trigger commands + users"
	jParams, err := processJunctions(desc, getFormationTriggerCommands(), getFormationTriggerCommandUsers)
	if err != nil {
		return err
	}

	return qtx.CreateFormationTriggerCommandsUsersJunctionBulk(ctx, database.CreateFormationTriggerCommandsUsersJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		TriggerCommandID: jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}
