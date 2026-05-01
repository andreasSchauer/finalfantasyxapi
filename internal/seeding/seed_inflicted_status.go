package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type InflictedStatus struct {
	ID                int32
	StatusConditionID int32
	StatusCondition   string `json:"name"`
	Probability       int32  `json:"probability"`
	DurationType      string `json:"duration_type"`
	Amount            *int32 `json:"amount"`
}

func (is InflictedStatus) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", is),
		is.StatusConditionID,
		is.Probability,
		is.DurationType,
		h.DerefOrNil(is.Amount),
	}
}

func (is InflictedStatus) GetID() int32 {
	return is.ID
}

func (is *InflictedStatus) SetID(id int32) {
	is.ID = id
}

func (is InflictedStatus) Error() string {
	return fmt.Sprintf("inflicted status with condition: %s, probability: %d, duration type: %s, amount: %v", is.StatusCondition, is.Probability, is.DurationType, is.Amount)
}

func (l *Lookup) seedInflictedStatus(qtx *database.Queries, inflictedStatus InflictedStatus) (InflictedStatus, error) {
	var err error

	inflictedStatus.StatusConditionID, err = assignFK(inflictedStatus.StatusCondition, l.StatusConditions)
	if err != nil {
		return InflictedStatus{}, h.NewErr(inflictedStatus.Error(), err)
	}

	dbInflictedStatus, err := qtx.CreateInflictedStatus(context.Background(), database.CreateInflictedStatusParams{
		DataHash:          generateDataHash(inflictedStatus),
		StatusConditionID: inflictedStatus.StatusConditionID,
		Probability:       inflictedStatus.Probability,
		DurationType:      database.DurationType(inflictedStatus.DurationType),
		Amount:            h.GetNullInt32(inflictedStatus.Amount),
	})
	if err != nil {
		return InflictedStatus{}, h.NewErr(inflictedStatus.Error(), err, "couldn't create inflicted status")
	}
	inflictedStatus.ID = dbInflictedStatus.ID

	return inflictedStatus, nil
}



func (l *Lookup) loop4SeedInflictedStatusses(qtx *database.Queries, ctx context.Context) error {
	statusses, err := l.extractInflictedStatusses()
	if err != nil {
		return err
	}

	params := database.CreateInflictedStatusBulkParams{
		DataHash:   		make([]string, len(statusses)),
		StatusConditionID: 	make([]int32, len(statusses)),
		Probability: 		make([]int32, len(statusses)),
		DurationType: 		make([]database.DurationType, len(statusses)),
		Amount: 			make([]sql.NullInt32, len(statusses)),
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}


func (l *Lookup) extractInflictedStatusses() ([]InflictedStatus, error) {
	statusses := []InflictedStatus{}

	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		if autoAbility.OnHitStatus == nil {
			continue
		}

		status, err := l.prepareInflictedStatus(autoAbility.OnHitStatus)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, *status)
	}

	for i := range l.json.monsters {
		monster := &l.json.monsters[i]

		for j := range monster.AlteredStates {
			stateStatusses, err := l.extractAltStateInflictedStatusses(&monster.AlteredStates[j])
			if err != nil {
				return nil, err
			}

			statusses = append(statusses, stateStatusses...)
		}
	}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]

		statussesNew, err := l.prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]

		statussesNew, err := l.prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range l.json.items {
		item := &l.json.items[i]

		statussesNew, err := l.prepareAbilityInflictedStatusses(item.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]

		statussesNew, err := l.prepareAbilityInflictedStatusses(command.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		statussesNew, err := l.prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]

		statussesNew, err := l.prepareAbilityInflictedStatusses(ability.BattleInteractions)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, statussesNew...)
	}

	return dedupeRows(statusses, l.Hashes), nil
}

func (l *Lookup) prepareAbilityInflictedStatusses(bis []BattleInteraction) ([]InflictedStatus, error) {
	statusses := []InflictedStatus{}

	for i := range bis {
		bi := &bis[i]

		for j := range bi.InflictedStatusConditions {
			status, err := l.prepareInflictedStatus(&bi.InflictedStatusConditions[j])
			if err != nil {
				return nil, err
			}
			statusses = append(statusses, *status)
		}

		for j := range bi.CopiedStatusConditions {
			status, err := l.prepareInflictedStatus(&bi.CopiedStatusConditions[j])
			if err != nil {
				return nil, err
			}
			statusses = append(statusses, *status)
		}
	}

	return statusses, nil
}

func (l *Lookup) extractAltStateInflictedStatusses(state *AlteredState) ([]InflictedStatus, error) {
	statusses := []InflictedStatus{}

	for i := range state.Changes {
		change := &state.Changes[i]

		if change.AddedStatus == nil {
			continue
		}

		status, err := l.prepareInflictedStatus(change.AddedStatus)
		if err != nil {
			return nil, err
		}

		statusses = append(statusses, *status)
	}

	return statusses, nil
}


func (l *Lookup) prepareInflictedStatus(is *InflictedStatus) (*InflictedStatus, error) {
	var err error

	is.StatusConditionID, err = assignFK(is.StatusCondition, l.StatusConditions)
	if err != nil {
		return nil, err
	}

	return is, nil
}