package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedStatusConditions(qtx *database.Queries, ctx context.Context) error {
	statusses, err := l.extractStatusConditions()
	if err != nil {
		return err
	}

	params := database.CreateStatusConditionBulkParams{
		DataHash:          make([]string, len(statusses)),
		Name:              make([]string, len(statusses)),
		Category:          make([]database.StatusConditionCategory, len(statusses)),
		IsPermanent:       make([]bool, len(statusses)),
		Effect:            make([]string, len(statusses)),
		Visualization:     make([]sql.NullString, len(statusses)),
		NullifyArmored:    make([]database.NullNullifyArmored, len(statusses)),
		AddedElemResistID: make([]sql.NullInt32, len(statusses)),
		InflictedDelayID:  make([]sql.NullInt32, len(statusses)),
	}

	for i, s := range statusses {
		params.DataHash[i] = generateDataHash(s)
		params.Name[i] = s.Name
		params.Category[i] = database.StatusConditionCategory(s.Category)
		params.IsPermanent[i] = s.IsPermanent
		params.Effect[i] = s.Effect
		params.Visualization[i] = h.GetNullString(s.Visualization)
		params.NullifyArmored[i] = database.ToNullNullifyArmored(s.NullifyArmored)
		params.AddedElemResistID[i] = h.ObjPtrToNullInt32ID(s.AddedElemResist)
		params.InflictedDelayID[i] = h.ObjPtrToNullInt32ID(s.CtbOnInfliction)
	}

	dbRows, err := qtx.CreateStatusConditionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create status conditions: %v", err)
	}

	for i, row := range dbRows {
		statusses[i].ID = row.ID
		l.json.statusConditions[i].ID = row.ID
		l.StatusConditions[statusses[i].Name] = statusses[i]
		l.StatusConditionsID[row.ID] = statusses[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractStatusConditions() ([]StatusCondition, error) {
	statusses := []StatusCondition{}
	var err error

	for i := range l.json.statusConditions {
		status := &l.json.statusConditions[i]

		if status.AddedElemResist != nil {
			status.AddedElemResist.ID, err = l.getHashID(status.AddedElemResist)
			if err != nil {
				return nil, err
			}
		}

		if status.CtbOnInfliction != nil {
			status.CtbOnInfliction.ID, err = l.getHashID(status.CtbOnInfliction)
			if err != nil {
				return nil, err
			}
		}

		statusses = append(statusses, *status)
	}

	return dedupeRows(statusses, l.Hashes), nil
}

func (l *Lookup) completeStatusConditions() error {
	for i := range l.json.statusConditions {
		status := &l.json.statusConditions[i]

		err := assignIDs(l, status.StatChanges)
		if err != nil {
			return err
		}

		err = assignIDs(l, status.ModifierChanges)
		if err != nil {
			return err
		}

		l.StatusConditions[status.Name] = *status
		l.StatusConditionsID[status.ID] = *status
	}

	return nil
}
