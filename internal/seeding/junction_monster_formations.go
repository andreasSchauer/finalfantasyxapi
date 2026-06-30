package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getMonsterFormationEncounterAreas(mf MonsterFormation) ([]EncounterArea, error) {
	return mf.EncounterAreas, nil
}

func (l *Lookup) seedJuncMonsterFormationsEncounterAreas(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster-formations + encounter areas"
	jParams, err := processJunctions(l, desc, l.json.monsterFormations, l.getMonsterFormationEncounterAreas)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterFormationsEncounterAreasJunctionBulk(ctx, database.CreateMonsterFormationsEncounterAreasJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterFormationID: jParams.ParentIDs,
		EncounterAreaID:    jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterFormationTriggerCommands(mf MonsterFormation) ([]FormationTriggerCommand, error) {
	return mf.TriggerCommands, nil
}

func (l *Lookup) seedJuncMonsterFormationsTriggerCommands(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster-formations + trigger commands"
	jParams, err := processJunctions(l, desc, l.json.monsterFormations, l.getMonsterFormationTriggerCommands)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterFormationsTriggerCommandsJunctionBulk(ctx, database.CreateMonsterFormationsTriggerCommandsJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterFormationID: jParams.ParentIDs,
		TriggerCommandID:   jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterSelections() []MonsterSelection {
	selections := []MonsterSelection{}

	for _, formation := range l.json.monsterFormations {
		selections = append(selections, formation.MonsterSelection)
	}

	return selections
}

func (l *Lookup) getMonsterSelectionMonsterAmounts(ms MonsterSelection) ([]MonsterAmount, error) {
	return ms.Monsters, nil
}

func (l *Lookup) seedJuncMonsterSelectionMonsterAmounts(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster selection + monsters"
	jParams, err := processJunctions(l, desc, l.getMonsterSelections(), l.getMonsterSelectionMonsterAmounts)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterSelectionsMonstersJunctionBulk(ctx, database.CreateMonsterSelectionsMonstersJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterSelectionID: jParams.ParentIDs,
		MonsterAmountID:    jParams.ChildIDs,
	})
}
