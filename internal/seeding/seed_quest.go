package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop4SeedQuests(qtx *database.Queries, ctx context.Context) error {
	quests, err := l.extractQuests()
	if err != nil {
		return err
	}

	params := database.CreateQuestBulkParams{
		DataHash:     make([]string, len(quests)),
		Name:         make([]string, len(quests)),
		Type:         make([]database.QuestType, len(quests)),
		Availability: make([]database.AvailabilityType, len(quests)),
		IsRepeatable: make([]bool, len(quests)),
		CompletionID: make([]sql.NullInt32, len(quests)),
	}

	for i, q := range quests {
		params.DataHash[i] = generateDataHash(q)
		params.Name[i] = q.Name
		params.Type[i] = q.Type
		params.Availability[i] = database.AvailabilityType(q.Availability)
		params.IsRepeatable[i] = q.IsRepeatable
		params.CompletionID[i] = h.ObjPtrToNullInt32ID(q.Completion)
	}

	dbRows, err := qtx.CreateQuestBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create quests: %v", err)
	}

	for i, row := range dbRows {
		quests[i].ID = row.ID
		l.Quests[Key(quests[i])] = quests[i]
		l.QuestsID[row.ID] = quests[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractQuests() ([]Quest, error) {
	quests := []Quest{}
	var err error

	for i := range l.json.sidequests {
		quest := &l.json.sidequests[i].Quest
		quest.Type = database.QuestTypeSidequest

		if quest.Completion != nil {
			quest.Completion.ID, err = l.getHashID(quest.Completion)
			if err != nil {
				return nil, err
			}
		}

		quests = append(quests, *quest)
	}

	for i := range l.json.sidequests {
		sidequest := &l.json.sidequests[i]

		for j := range sidequest.Subquests {
			quest := &sidequest.Subquests[j].Quest
			quest.Type = database.QuestTypeSubquest

			if quest.Completion != nil {
				quest.Completion.ID, err = l.getHashID(quest.Completion)
				if err != nil {
					return nil, err
				}
			}

			quests = append(quests, *quest)
		}
	}

	return dedupeRows(quests, l.Hashes), nil
}
