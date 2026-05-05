package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop5SeedSidequests(qtx *database.Queries, ctx context.Context) error {
	sidequests, err := l.extractSidequests()
	if err != nil {
		return err
	}

	params := database.CreateSidequestBulkParams{
		DataHash: make([]string, len(sidequests)),
		QuestID:  make([]int32, len(sidequests)),
	}

	for i, s := range sidequests {
		params.DataHash[i] = generateDataHash(s)
		params.QuestID[i] = s.Quest.ID
	}

	dbRows, err := qtx.CreateSidequestBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create sidequests: %v", err)
	}

	for i, row := range dbRows {
		sidequests[i].ID = row.ID
		l.json.sidequests[i].ID = row.ID
		l.Sidequests[Key(sidequests[i])] = sidequests[i]
		l.SidequestsID[row.ID] = sidequests[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractSidequests() ([]Sidequest, error) {
	sidequests := []Sidequest{}
	var err error

	for i := range l.json.sidequests {
		sidequest := &l.json.sidequests[i]

		sidequest.Quest.ID, err = l.getHashID(sidequest.Quest)
		if err != nil {
			return nil, err
		}

		sidequests = append(sidequests, *sidequest)
	}

	return dedupeRows(sidequests, l.Hashes), nil
}

func (l *Lookup) completeSidequests() error {
	for i := range l.json.sidequests {
		sidequest := &l.json.sidequests[i]

		err := assignIDs(l, sidequest.Subquests)
		if err != nil {
			return err
		}

		l.Sidequests[Key(sidequest)] = *sidequest
		l.SidequestsID[sidequest.ID] = *sidequest
	}

	return nil
}

func (l *Lookup) loop6SeedSubquests(qtx *database.Queries, ctx context.Context) error {
	subquests, err := l.extractSubquests()
	if err != nil {
		return err
	}

	params := database.CreateSubquestBulkParams{
		DataHash:    make([]string, len(subquests)),
		QuestID:     make([]int32, len(subquests)),
		SidequestID: make([]int32, len(subquests)),
	}

	for i, q := range subquests {
		params.DataHash[i] = generateDataHash(q)
		params.QuestID[i] = q.Quest.ID
		params.SidequestID[i] = q.SidequestID
	}

	dbRows, err := qtx.CreateSubquestBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create subquests: %v", err)
	}

	for i, row := range dbRows {
		subquests[i].ID = row.ID
		l.Subquests[Key(subquests[i])] = subquests[i]
		l.SubquestsID[row.ID] = subquests[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractSubquests() ([]Subquest, error) {
	quests := []Subquest{}
	var err error

	for i := range l.json.sidequests {
		sidequest := &l.json.sidequests[i]

		for j := range sidequest.Subquests {
			subquest := &sidequest.Subquests[j]
			subquest.SidequestID = sidequest.ID

			subquest.Quest.ID, err = l.getHashID(subquest.Quest)
			if err != nil {
				return nil, err
			}

			quests = append(quests, *subquest)
		}
	}

	return dedupeRows(quests, l.Hashes), nil
}
