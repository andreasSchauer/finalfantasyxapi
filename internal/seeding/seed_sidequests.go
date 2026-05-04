package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Sidequest struct {
	ID 			int32
	Quest
	Subquests  	[]Subquest	`json:"subquests"`
}

func (s Sidequest) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Quest.ID,
	}
}

func (s Sidequest) GetID() int32 {
	return s.ID
}

func (s Sidequest) Error() string {
	return fmt.Sprintf("sidequest %s", s.Name)
}

func (s Sidequest) GetResParamsQuest() h.ResParamsQuest {
	return h.ResParamsQuest{
		ID:        		s.ID,
		Sidequest:		&s.Name,
		Subquest:  		nil,
		Type:			string(s.Quest.Type),
	}
}

type Subquest struct {
	ID 				int32
	Quest
	SidequestID 	int32
}

func (s Subquest) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Quest.ID,
		s.SidequestID,
	}
}

func (s Subquest) GetID() int32 {
	return s.ID
}

func (s *Subquest) SetID(id int32) {
	s.ID = id
}

func (s Subquest) Error() string {
	return fmt.Sprintf("subquest %s", s.Name)
}

func (s Subquest) GetResParamsQuest() h.ResParamsQuest {
	return h.ResParamsQuest{
		ID:        		s.ID,
		Sidequest: 		nil,
		Subquest:  		&s.Name,
		Type:			string(s.Quest.Type),
	}
}


func (l *Lookup) loop5SeedSidequests(qtx *database.Queries, ctx context.Context) error {
	sidequests, err := l.extractSidequests()
	if err != nil {
		return err
	}

	params := database.CreateSidequestBulkParams{
		DataHash:   make([]string, len(sidequests)),
		QuestID: 	make([]int32, len(sidequests)),
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
		l.Sidequests[sidequests[i].Name] = sidequests[i]
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

		l.Sidequests[sidequest.Name] = *sidequest
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
		DataHash:   	make([]string, len(subquests)),
		QuestID: 		make([]int32, len(subquests)),
		SidequestID: 	make([]int32, len(subquests)),
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
		l.Subquests[subquests[i].Name] = subquests[i]
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


func (l *Lookup) loop4SeedCompletionAreas(qtx *database.Queries, ctx context.Context) error {
	areas, err := l.extractCompletionAreas()
	if err != nil {
		return err
	}

	params := database.CreateCompletionAreaBulkParams{
		DataHash:   	make([]string, len(areas)),
		CompletionID: 	make([]int32, len(areas)),
		AreaID: 		make([]int32, len(areas)),
		Notes: 			make([]sql.NullString, len(areas)),
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

	for i := range l.json.sidequests {
		sidequest := &l.json.sidequests[i]

		if sidequest.Completion != nil {
			areasNew, err := l.prepareCompletionAreas(sidequest.Completion.Areas, sidequest.Completion.ID)
			if err != nil {
				return nil, err
			}
			areas = append(areas, areasNew...)
		}

		for j := range sidequest.Subquests {
			subquest := &sidequest.Subquests[j]

			if subquest.Completion != nil {
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