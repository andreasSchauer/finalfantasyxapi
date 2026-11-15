package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getFormationTriggerCommands() []seedtypes.FormationTriggerCommand {
	commands := []seedtypes.FormationTriggerCommand{}

	for _, formation := range sh.RawJSON.MonsterFormations {
		commands = append(commands, formation.TriggerCommands...)
	}

	return commands
}

func getFormationTriggerCommandUsers(tc seedtypes.FormationTriggerCommand) ([]seedtypes.CharacterClass, error) {
	return lookup.GetResources(tc.Users, lookup.L.CharClasses)
}

func SeedJuncFormationTriggerCommandsUsers(qtx *database.Queries, ctx context.Context) error {
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
