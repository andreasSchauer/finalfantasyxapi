package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getMonsterFormationEncounterAreas(mf seedtypes.MonsterFormation) ([]seedtypes.EncounterArea, error) {
	return mf.EncounterAreas, nil
}

func seedJuncMonsterFormationsEncounterAreas(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster-formations + encounter areas"
	jParams, err := processJunctions(desc, rawJSON.monsterFormations, getMonsterFormationEncounterAreas)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterFormationsEncounterAreasJunctionBulk(ctx, database.CreateMonsterFormationsEncounterAreasJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterFormationID: jParams.ParentIDs,
		EncounterAreaID:    jParams.ChildIDs,
	})
}

func getMonsterFormationTriggerCommands(mf seedtypes.MonsterFormation) ([]seedtypes.FormationTriggerCommand, error) {
	return mf.TriggerCommands, nil
}

func seedJuncMonsterFormationsTriggerCommands(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster-formations + trigger commands"
	jParams, err := processJunctions(desc, rawJSON.monsterFormations, getMonsterFormationTriggerCommands)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterFormationsTriggerCommandsJunctionBulk(ctx, database.CreateMonsterFormationsTriggerCommandsJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterFormationID: jParams.ParentIDs,
		TriggerCommandID:   jParams.ChildIDs,
	})
}

func getMonsterSelections() []seedtypes.MonsterSelection {
	selections := []seedtypes.MonsterSelection{}

	for _, formation := range rawJSON.monsterFormations {
		selections = append(selections, formation.MonsterSelection)
	}

	return selections
}

func getMonsterSelectionMonsterAmounts(ms seedtypes.MonsterSelection) ([]seedtypes.MonsterAmount, error) {
	return ms.Monsters, nil
}

func seedJuncMonsterSelectionMonsterAmounts(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster selection + monsters"
	jParams, err := processJunctions(desc, getMonsterSelections(), getMonsterSelectionMonsterAmounts)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterSelectionsMonstersJunctionBulk(ctx, database.CreateMonsterSelectionsMonstersJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterSelectionID: jParams.ParentIDs,
		MonsterAmountID:    jParams.ChildIDs,
	})
}
