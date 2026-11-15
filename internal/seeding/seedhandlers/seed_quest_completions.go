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

func Loop3SeedQuestCompletions(qtx *database.Queries, ctx context.Context) error {
	completions, err := extractQuestCompletions()
	if err != nil {
		return err
	}

	params := database.CreateQuestCompletionBulkParams{
		DataHash:     make([]string, len(completions)),
		QuestID:      make([]int32, len(completions)),
		Condition:    make([]sql.NullString, len(completions)),
		ItemAmountID: make([]int32, len(completions)),
	}

	for i, qc := range completions {
		params.DataHash[i] = sh.GenerateDataHash(qc)
		params.QuestID[i] = qc.QuestID
		params.Condition[i] = h.GetNullString(qc.Condition)
		params.ItemAmountID[i] = qc.Reward.ID

	}

	dbRows, err := qtx.CreateQuestCompletionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create quest completions: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractQuestCompletions() ([]seedtypes.QuestCompletion, error) {
	completions := []seedtypes.QuestCompletion{}
	var err error

	for i := range sh.RawJSON.Sidequests {
		sidequest := &sh.RawJSON.Sidequests[i]

		if sidequest.Completion != nil {
			completion := sidequest.Completion
			completion.Reward.ID, err = sh.GetHashID(completion.Reward)
			if err != nil {
				return nil, err
			}

			completion.QuestID, err = sh.AssignFK(sidequest.Quest, lookup.L.Quests)
			if err != nil {
				return nil, err
			}

			completions = append(completions, *completion)
		}

		for j := range sidequest.Subquests {
			subquest := &sidequest.Subquests[j]

			if subquest.Completion != nil {
				completion := subquest.Completion
				completion.Reward.ID, err = sh.GetHashID(completion.Reward)
				if err != nil {
					return nil, err
				}

				completion.QuestID, err = sh.AssignFK(subquest.Quest, lookup.L.Quests)
				if err != nil {
					return nil, err
				}

				completions = append(completions, *completion)
			}
		}
	}

	return sh.DedupeRows(completions, sh.Hashes), nil
}

func Loop4SeedCompletionAreas(qtx *database.Queries, ctx context.Context) error {
	areas, err := extractCompletionAreas()
	if err != nil {
		return err
	}

	params := database.CreateCompletionAreaBulkParams{
		DataHash:     make([]string, len(areas)),
		CompletionID: make([]int32, len(areas)),
		AreaID:       make([]int32, len(areas)),
		Notes:        make([]sql.NullString, len(areas)),
	}

	for i, a := range areas {
		params.DataHash[i] = sh.GenerateDataHash(a)
		params.CompletionID[i] = a.CompletionID
		params.AreaID[i] = a.AreaID
		params.Notes[i] = h.GetNullString(a.Notes)
	}

	dbRows, err := qtx.CreateCompletionAreaBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create completion areas: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractCompletionAreas() ([]seedtypes.CompletionArea, error) {
	areas := []seedtypes.CompletionArea{}
	var err error

	for i := range sh.RawJSON.Sidequests {
		sidequest := &sh.RawJSON.Sidequests[i]

		if sidequest.Completion != nil {
			sidequest.Completion.ID, err = sh.GetHashID(sidequest.Completion)
			if err != nil {
				return nil, err
			}

			areasNew, err := prepareCompletionAreas(sidequest.Completion.Areas, sidequest.Completion.ID)
			if err != nil {
				return nil, err
			}

			areas = append(areas, areasNew...)
		}

		for j := range sidequest.Subquests {
			subquest := &sidequest.Subquests[j]

			if subquest.Completion != nil {
				subquest.Completion.ID, err = sh.GetHashID(subquest.Completion)
				if err != nil {
					return nil, err
				}

				areasNew, err := prepareCompletionAreas(subquest.Completion.Areas, subquest.Completion.ID)
				if err != nil {
					return nil, err
				}

				areas = append(areas, areasNew...)
			}
		}
	}

	return sh.DedupeRows(areas, sh.Hashes), nil
}

func prepareCompletionAreas(areas []seedtypes.CompletionArea, completionID int32) ([]seedtypes.CompletionArea, error) {
	areasNew := []seedtypes.CompletionArea{}
	var err error

	for i := range areas {
		area := &areas[i]
		area.CompletionID = completionID

		area.AreaID, err = sh.AssignFK(area.LocationArea, lookup.L.Areas)
		if err != nil {
			return nil, err
		}

		areasNew = append(areasNew, *area)
	}

	return areasNew, nil
}
