package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedQuestCompletions(qtx *database.Queries, ctx context.Context) error {
	completions, err := l.extractQuestCompletions()
	if err != nil {
		return err
	}

	params := database.CreateQuestCompletionBulkParams{
		DataHash:     make([]string, len(completions)),
		QuestID: 	  make([]int32, len(completions)),
		Condition:    make([]sql.NullString, len(completions)),
		ItemAmountID: make([]int32, len(completions)),
	}

	for i, qc := range completions {
		params.DataHash[i] = generateDataHash(qc)
		params.QuestID[i] = qc.QuestID
		params.Condition[i] = h.GetNullString(qc.Condition)
		params.ItemAmountID[i] = qc.Reward.ID

	}

	dbRows, err := qtx.CreateQuestCompletionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create quest completions: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractQuestCompletions() ([]QuestCompletion, error) {
	completions := []QuestCompletion{}
	var err error

	for i := range l.json.sidequests {
		sidequest := &l.json.sidequests[i]

		if sidequest.Completion != nil {
			completion := sidequest.Completion
			completion.Reward.ID, err = l.GetHashID(completion.Reward)
			if err != nil {
				return nil, err
			}

			completion.QuestID, err = assignFK(sidequest.Quest, l.Quests)
			if err != nil {
				return nil, err
			}

			completions = append(completions, *completion)
		}

		for j := range sidequest.Subquests {
			subquest := &sidequest.Subquests[j]

			if subquest.Completion != nil {
				completion := subquest.Completion
				completion.Reward.ID, err = l.GetHashID(completion.Reward)
				if err != nil {
					return nil, err
				}

				completion.QuestID, err = assignFK(subquest.Quest, l.Quests)
				if err != nil {
					return nil, err
				}

				completions = append(completions, *completion)
			}
		}
	}

	return dedupeRows(completions, l.Hashes), nil
}

func (l *Lookup) loop4SeedCompletionAreas(qtx *database.Queries, ctx context.Context) error {
	areas, err := l.extractCompletionAreas()
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
		params.DataHash[i] = generateDataHash(a)
		params.CompletionID[i] = a.CompletionID
		params.AreaID[i] = a.AreaID
		params.Notes[i] = h.GetNullString(a.Notes)
	}

	dbRows, err := qtx.CreateCompletionAreaBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create completion areas: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractCompletionAreas() ([]CompletionArea, error) {
	areas := []CompletionArea{}
	var err error

	for i := range l.json.sidequests {
		sidequest := &l.json.sidequests[i]

		if sidequest.Completion != nil {
			sidequest.Completion.ID, err = l.GetHashID(sidequest.Completion)
			if err != nil {
				return nil, err
			}

			areasNew, err := l.prepareCompletionAreas(sidequest.Completion.Areas, sidequest.Completion.ID)
			if err != nil {
				return nil, err
			}

			areas = append(areas, areasNew...)
		}

		for j := range sidequest.Subquests {
			subquest := &sidequest.Subquests[j]

			if subquest.Completion != nil {
				subquest.Completion.ID, err = l.GetHashID(subquest.Completion)
				if err != nil {
					return nil, err
				}

				areasNew, err := l.prepareCompletionAreas(subquest.Completion.Areas, subquest.Completion.ID)
				if err != nil {
					return nil, err
				}

				areas = append(areas, areasNew...)
			}
		}
	}

	return dedupeRows(areas, l.Hashes), nil
}

func (l *Lookup) prepareCompletionAreas(areas []CompletionArea, completionID int32) ([]CompletionArea, error) {
	areasNew := []CompletionArea{}
	var err error

	for i := range areas {
		area := &areas[i]
		area.CompletionID = completionID

		area.AreaID, err = assignFK(area.LocationArea, l.Areas)
		if err != nil {
			return nil, err
		}

		areasNew = append(areasNew, *area)
	}

	return areasNew, nil
}
