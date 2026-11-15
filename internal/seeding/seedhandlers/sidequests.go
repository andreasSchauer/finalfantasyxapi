package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop5Sidequests(qtx *database.Queries, ctx context.Context) error {
	sidequests, err := extractSidequests()
	if err != nil {
		return err
	}

	params := database.CreateSidequestBulkParams{
		DataHash: make([]string, len(sidequests)),
		QuestID:  make([]int32, len(sidequests)),
	}

	for i, s := range sidequests {
		params.DataHash[i] = sh.GenerateDataHash(s)
		params.QuestID[i] = s.Quest.ID
	}

	dbRows, err := qtx.CreateSidequestBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create sidequests: %v", err)
	}

	for i, row := range dbRows {
		sidequests[i].ID = row.ID
		sh.RawJSON.Sidequests[i].ID = row.ID
		lookup.L.Sidequests[seedtypes.Key(sidequests[i])] = sidequests[i]
		lookup.L.SidequestsID[row.ID] = sidequests[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractSidequests() ([]seedtypes.Sidequest, error) {
	sidequests := []seedtypes.Sidequest{}
	var err error

	for i := range sh.RawJSON.Sidequests {
		sidequest := &sh.RawJSON.Sidequests[i]

		sidequest.Quest.ID, err = sh.GetHashID(sidequest.Quest)
		if err != nil {
			return nil, err
		}

		sidequests = append(sidequests, *sidequest)
	}

	return sh.DedupeRows(sidequests, sh.Hashes), nil
}

func CompleteSidequests() error {
	for i := range sh.RawJSON.Sidequests {
		sidequest := &sh.RawJSON.Sidequests[i]

		err := sh.AssignIDs(sidequest.Subquests)
		if err != nil {
			return err
		}

		lookup.L.Sidequests[seedtypes.Key(sidequest)] = *sidequest
		lookup.L.SidequestsID[sidequest.ID] = *sidequest
	}

	return nil
}

func Loop6Subquests(qtx *database.Queries, ctx context.Context) error {
	subquests, err := extractSubquests()
	if err != nil {
		return err
	}

	params := database.CreateSubquestBulkParams{
		DataHash:    make([]string, len(subquests)),
		QuestID:     make([]int32, len(subquests)),
		SidequestID: make([]int32, len(subquests)),
	}

	for i, q := range subquests {
		params.DataHash[i] = sh.GenerateDataHash(q)
		params.QuestID[i] = q.Quest.ID
		params.SidequestID[i] = q.SidequestID
	}

	dbRows, err := qtx.CreateSubquestBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create subquests: %v", err)
	}

	for i, row := range dbRows {
		subquests[i].ID = row.ID
		lookup.L.Subquests[seedtypes.Key(subquests[i])] = subquests[i]
		lookup.L.SubquestsID[row.ID] = subquests[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractSubquests() ([]seedtypes.Subquest, error) {
	quests := []seedtypes.Subquest{}
	var err error

	for i := range sh.RawJSON.Sidequests {
		sidequest := &sh.RawJSON.Sidequests[i]

		for j := range sidequest.Subquests {
			subquest := &sidequest.Subquests[j]
			subquest.SidequestID = sidequest.ID

			subquest.Quest.ID, err = sh.GetHashID(subquest.Quest)
			if err != nil {
				return nil, err
			}

			quests = append(quests, *subquest)
		}
	}

	return sh.DedupeRows(quests, sh.Hashes), nil
}
