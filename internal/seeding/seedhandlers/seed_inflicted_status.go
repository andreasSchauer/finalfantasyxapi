package seedhandlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop4SeedInflictedStatusses(qtx *database.Queries, ctx context.Context) error {
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
		params.DataHash[i] = sh.GenerateDataHash(s)
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
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractInflictedStatusses() ([]seedtypes.InflictedStatus, error) {
	statusses := []seedtypes.InflictedStatus{}

	for i := range sh.RawJSON.AutoAbilities {
		autoAbility := &sh.RawJSON.AutoAbilities[i]

		if autoAbility.OnHitStatus == nil {
			continue
		}

		status, err := prepareInflictedStatus(autoAbility.OnHitStatus)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, *status)
	}

	for i := range sh.RawJSON.Monsters {
		monster := &sh.RawJSON.Monsters[i]

		for j := range monster.AlteredStates {
			stateStatusses, err := extractAltStateInflictedStatusses(&monster.AlteredStates[j])
			if err != nil {
				return nil, err
			}

			statusses = append(statusses, stateStatusses...)
		}
	}

	for i := range sh.RawJSON.PlayerAbilities {
		ability := &sh.RawJSON.PlayerAbilities[i]

		statussesNew, err := prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range sh.RawJSON.OverdriveAbilities {
		ability := &sh.RawJSON.OverdriveAbilities[i]

		statussesNew, err := prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range sh.RawJSON.Items {
		item := &sh.RawJSON.Items[i]

		statussesNew, err := prepareAbilityInflictedStatusses(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range sh.RawJSON.TriggerCommands {
		command := &sh.RawJSON.TriggerCommands[i]

		statussesNew, err := prepareAbilityInflictedStatusses(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range sh.RawJSON.MiscAbilities {
		ability := &sh.RawJSON.MiscAbilities[i]

		statussesNew, err := prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range sh.RawJSON.EnemyAbilities {
		ability := &sh.RawJSON.EnemyAbilities[i]

		statussesNew, err := prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	return sh.DedupeRows(statusses, sh.Hashes), nil
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

	is.StatusConditionID, err = sh.AssignFK(is.StatusCondition, lookup.L.StatusConditions)
	if err != nil {
		return nil, err
	}

	return is, nil
}
