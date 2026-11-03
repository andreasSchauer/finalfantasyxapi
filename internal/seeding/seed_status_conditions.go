package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop3SeedStatusConditions(qtx *database.Queries, ctx context.Context) error {
	statusses, err := extractStatusConditions()
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
	}

	dbRows, err := qtx.CreateStatusConditionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create status conditions: %v", err)
	}

	for i, row := range dbRows {
		statusses[i].ID = row.ID
		rawJSON.statusConditions[i].ID = row.ID
		L.StatusConditions[seedtypes.Key(statusses[i])] = statusses[i]
		L.StatusConditionsID[row.ID] = statusses[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractStatusConditions() ([]seedtypes.StatusCondition, error) {
	statusses := []seedtypes.StatusCondition{}
	var err error

	for i := range rawJSON.statusConditions {
		status := &rawJSON.statusConditions[i]

		if status.AddedElemResist != nil {
			status.AddedElemResist.ID, err = GetHashID(status.AddedElemResist)
			if err != nil {
				return nil, err
			}
		}

		statusses = append(statusses, *status)
	}

	return dedupeRows(statusses, hashes), nil
}

func completeStatusConditions() error {
	for i := range rawJSON.statusConditions {
		status := &rawJSON.statusConditions[i]

		err := assignIDs(status.StatChanges)
		if err != nil {
			return err
		}

		err = assignIDs(status.ModifierChanges)
		if err != nil {
			return err
		}

		L.StatusConditions[seedtypes.Key(status)] = *status
		L.StatusConditionsID[status.ID] = *status
	}

	return nil
}
