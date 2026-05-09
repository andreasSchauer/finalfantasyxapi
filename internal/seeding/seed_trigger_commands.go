package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedTriggerCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := l.extractTriggerCommands()
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
		l.json.triggerCommands[i].ID = row.ID
		l.TriggerCommands[Key(commands[i])] = commands[i]
		l.TriggerCommandsID[row.ID] = commands[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractTriggerCommands() ([]TriggerCommand, error) {
	commands := []TriggerCommand{}
	var err error

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		command.Ability.ID, err = l.GetHashID(command.Ability)
		if err != nil {
			return nil, err
		}

		command.TopmenuID, err = assignFKPtr(command.Topmenu, l.Topmenus)
		if err != nil {
			return nil, err
		}

		commands = append(commands, *command)
	}

	return dedupeRows(commands, l.Hashes), nil
}

func (l *Lookup) completeTriggerCommands() error {
	for i := range l.json.triggerCommands {
		ability := &l.json.triggerCommands[i]

		err := l.completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		l.TriggerCommands[Key(ability)] = *ability
		l.TriggerCommandsID[ability.ID] = *ability
	}

	return nil
}

func (l *Lookup) getTriggerCommandRelatedStats(tc TriggerCommand) ([]Stat, error) {
	return getResources(tc.RelatedStats, l.Stats)
}

func (l *Lookup) seedJuncTriggerCommandsRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "trigger commands + related stats"
	jParams, err := processJunctions(l, desc, l.json.triggerCommands, l.getTriggerCommandRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateTriggerCommandsRelatedStatsJunctionBulk(ctx, database.CreateTriggerCommandsRelatedStatsJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		TriggerCommandID: jParams.ParentIDs,
		StatID:           jParams.ChildIDs,
	})
}
