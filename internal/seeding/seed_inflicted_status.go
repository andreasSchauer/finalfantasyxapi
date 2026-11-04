package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop4SeedInflictedStatusses(qtx *database.Queries, ctx context.Context) error {
	statusses, err := extractInflictedStatusses()
	if err != nil {
		return err
	}

	params := database.CreateInflictedStatusBulkParams{
		DataHash:          make([]string, len(statusses)),
		StatusConditionID: make([]int32, len(statusses)),
		Probability:       make([]int32, len(statusses)),
		DurationType:      make([]database.DurationType, len(statusses)),
		Amount:            make([]sql.NullInt32, len(statusses)),
	}

	for i, s := range statusses {
		params.DataHash[i] = generateDataHash(s)
		params.StatusConditionID[i] = s.StatusConditionID
		params.Probability[i] = s.Probability
		params.DurationType[i] = database.DurationType(s.DurationType)
		params.Amount[i] = h.GetNullInt32(s.Amount)
	}

	dbRows, err := qtx.CreateInflictedStatusBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create inflicted statusses: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractInflictedStatusses() ([]seedtypes.InflictedStatus, error) {
	statusses := []seedtypes.InflictedStatus{}

	for i := range rawJSON.autoAbilities {
		autoAbility := &rawJSON.autoAbilities[i]

		if autoAbility.OnHitStatus == nil {
			continue
		}

		status, err := prepareInflictedStatus(autoAbility.OnHitStatus)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, *status)
	}

	for i := range rawJSON.monsters {
		monster := &rawJSON.monsters[i]

		for j := range monster.AlteredStates {
			stateStatusses, err := extractAltStateInflictedStatusses(&monster.AlteredStates[j])
			if err != nil {
				return nil, err
			}

			statusses = append(statusses, stateStatusses...)
		}
	}

	for i := range rawJSON.playerAbilities {
		ability := &rawJSON.playerAbilities[i]

		statussesNew, err := prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range rawJSON.overdriveAbilities {
		ability := &rawJSON.overdriveAbilities[i]

		statussesNew, err := prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range rawJSON.items {
		item := &rawJSON.items[i]

		statussesNew, err := prepareAbilityInflictedStatusses(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range rawJSON.triggerCommands {
		command := &rawJSON.triggerCommands[i]

		statussesNew, err := prepareAbilityInflictedStatusses(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range rawJSON.miscAbilities {
		ability := &rawJSON.miscAbilities[i]

		statussesNew, err := prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range rawJSON.enemyAbilities {
		ability := &rawJSON.enemyAbilities[i]

		statussesNew, err := prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	return dedupeRows(statusses, hashes), nil
}

func prepareAbilityInflictedStatusses(bis []seedtypes.BattleInteraction) ([]seedtypes.InflictedStatus, error) {
	statusses := []seedtypes.InflictedStatus{}

	for i := range bis {
		bi := &bis[i]

		for j := range bi.InflictedStatusConditions {
			status, err := prepareInflictedStatus(&bi.InflictedStatusConditions[j])
			if err != nil {
				return nil, err
			}
			statusses = append(statusses, *status)
		}

		for j := range bi.CopiedStatusConditions {
			status, err := prepareInflictedStatus(&bi.CopiedStatusConditions[j])
			if err != nil {
				return nil, err
			}
			statusses = append(statusses, *status)
		}
	}

	return statusses, nil
}

func extractAltStateInflictedStatusses(state *seedtypes.AlteredState) ([]seedtypes.InflictedStatus, error) {
	statusses := []seedtypes.InflictedStatus{}

	for i := range state.Changes {
		change := &state.Changes[i]

		if change.AddedStatus == nil {
			continue
		}

		status, err := prepareInflictedStatus(change.AddedStatus)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, *status)
	}

	return statusses, nil
}

func prepareInflictedStatus(is *seedtypes.InflictedStatus) (*seedtypes.InflictedStatus, error) {
	var err error

	is.StatusConditionID, err = assignFK(is.StatusCondition, lookup.L.StatusConditions)
	if err != nil {
		return nil, err
	}

	return is, nil
}
