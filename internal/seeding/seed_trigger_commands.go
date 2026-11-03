package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop3SeedTriggerCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := extractTriggerCommands()
	if err != nil {
		return err
	}

	params := database.CreateTriggerCommandBulkParams{
		DataHash:    make([]string, len(commands)),
		AbilityID:   make([]int32, len(commands)),
		Description: make([]string, len(commands)),
		Effect:      make([]string, len(commands)),
		Cursor:      make([]database.TargetType, len(commands)),
		TopmenuID:   make([]sql.NullInt32, len(commands)),
	}

	for i, c := range commands {
		params.DataHash[i] = generateDataHash(c)
		params.AbilityID[i] = c.Ability.ID
		params.Description[i] = c.Description
		params.Effect[i] = c.Effect
		params.Cursor[i] = database.TargetType(c.Cursor)
		params.TopmenuID[i] = h.GetNullInt32(c.TopmenuID)
	}

	dbRows, err := qtx.CreateTriggerCommandBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create trigger command: %v", err)
	}

	for i, row := range dbRows {
		commands[i].ID = row.ID
		rawJSON.triggerCommands[i].ID = row.ID
		L.TriggerCommands[seedtypes.Key(commands[i])] = commands[i]
		L.TriggerCommandsID[row.ID] = commands[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractTriggerCommands() ([]seedtypes.TriggerCommand, error) {
	commands := []seedtypes.TriggerCommand{}
	var err error

	for i := range rawJSON.triggerCommands {
		command := &rawJSON.triggerCommands[i]

		command.Ability.ID, err = GetHashID(command.Ability)
		if err != nil {
			return nil, err
		}

		command.TopmenuID, err = assignFKPtr(command.Topmenu, L.Topmenus)
		if err != nil {
			return nil, err
		}

		commands = append(commands, *command)
	}

	return dedupeRows(commands, hashes), nil
}

func completeTriggerCommands() error {
	for i := range rawJSON.triggerCommands {
		ability := &rawJSON.triggerCommands[i]

		err := completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		L.TriggerCommands[seedtypes.Key(ability)] = *ability
		L.TriggerCommandsID[ability.ID] = *ability
	}

	return nil
}

func getTriggerCommandRelatedStats(tc seedtypes.TriggerCommand) ([]seedtypes.Stat, error) {
	return GetResources(tc.RelatedStats, L.Stats)
}

func seedJuncTriggerCommandsRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "trigger commands + related stats"
	jParams, err := processJunctions(desc, rawJSON.triggerCommands, getTriggerCommandRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateTriggerCommandsRelatedStatsJunctionBulk(ctx, database.CreateTriggerCommandsRelatedStatsJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		TriggerCommandID: jParams.ParentIDs,
		StatID:           jParams.ChildIDs,
	})
}
