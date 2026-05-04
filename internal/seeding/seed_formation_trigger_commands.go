package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop4SeedFormationTriggerCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := l.extractFormationTriggerCommands()
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractFormationTriggerCommands() ([]FormationTriggerCommand, error) {
	commands := []FormationTriggerCommand{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		for j := range mf.TriggerCommands {
			command := &mf.TriggerCommands[j]

			command.TriggerCommandID, err = assignFK(command.AbilityReference.Untyped(), l.TriggerCommands)
			if err != nil {
				return nil, err
			}

			commands = append(commands, *command)
		}
	}

	return dedupeRows(commands, l.Hashes), nil
}

func (l *Lookup) getFormationTriggerCommands() []FormationTriggerCommand {
	commands := []FormationTriggerCommand{}

	for _, formation := range l.json.monsterFormations {
		commands = append(commands, formation.TriggerCommands...)
	}

	return commands
}

func (l *Lookup) getFormationTriggerCommandUsers(tc FormationTriggerCommand) ([]CharacterClass, error) {
	return getResources(tc.Users, l.CharClasses)
}

func (l *Lookup) seedJuncFormationTriggerCommandsUsers(qtx *database.Queries, ctx context.Context) error {
	const desc string = "formation trigger commands + users"
	jParams, err := processJunctions(l, desc, l.getFormationTriggerCommands(), l.getFormationTriggerCommandUsers)
	if err != nil {
		return err
	}

	return qtx.CreateFormationTriggerCommandsUsersJunctionBulk(ctx, database.CreateFormationTriggerCommandsUsersJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		TriggerCommandID: jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}
