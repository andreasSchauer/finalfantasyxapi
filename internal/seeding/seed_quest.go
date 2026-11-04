package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop1SeedQuests(qtx *database.Queries, ctx context.Context) error {
	quests, err := extractQuests()
	if err != nil {
		return err
	}

	params := database.CreateQuestBulkParams{
		DataHash:     make([]string, len(quests)),
		Name:         make([]string, len(quests)),
		Type:         make([]database.QuestType, len(quests)),
		Availability: make([]database.AvailabilityType, len(quests)),
		IsRepeatable: make([]bool, len(quests)),
	}

	for i, q := range quests {
		params.DataHash[i] = generateDataHash(q)
		params.Name[i] = q.Name
		params.Type[i] = q.Type
		params.Availability[i] = database.AvailabilityType(q.Availability)
		params.IsRepeatable[i] = q.IsRepeatable
	}

	dbRows, err := qtx.CreateQuestBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create quests: %v", err)
	}

	for i, row := range dbRows {
		quests[i].ID = row.ID
		lookup.L.Quests[seedtypes.Key(quests[i])] = quests[i]
		lookup.L.QuestsID[row.ID] = quests[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractQuests() ([]seedtypes.Quest, error) {
	quests := []seedtypes.Quest{}

	for i := range rawJSON.sidequests {
		quest := &rawJSON.sidequests[i].Quest
		quest.Type = database.QuestTypeSidequest

		quests = append(quests, *quest)
	}

	for i := range rawJSON.sidequests {
		sidequest := &rawJSON.sidequests[i]

		for j := range sidequest.Subquests {
			quest := &sidequest.Subquests[j].Quest
			quest.Type = database.QuestTypeSubquest

			quests = append(quests, *quest)
		}
	}

	return dedupeRows(quests, hashes), nil
}
