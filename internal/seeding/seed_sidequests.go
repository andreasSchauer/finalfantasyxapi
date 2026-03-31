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
		s.Quest.ID,
		s.SidequestID,
	}
}

func (s Subquest) GetID() int32 {
	return s.ID
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


func (l *Lookup) seedSidequests(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/sidequests.json"

	var sidequests []Sidequest
	err := loadJSONFile(string(srcPath), &sidequests)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, sidequest := range sidequests {
			var err error
			sidequest.Type = database.QuestTypeSidequest

			sidequest.Quest, err = seedObjAssignID(qtx, sidequest.Quest, l.seedQuest)
			if err != nil {
				return h.NewErr(sidequest.Error(), err)
			}

			dbSidequest, err := qtx.CreateSidequest(context.Background(), database.CreateSidequestParams{
				DataHash: generateDataHash(sidequest),
				QuestID:  sidequest.Quest.ID,
			})
			if err != nil {
				return h.NewErr(sidequest.Error(), err, "couldn't create sidequest")
			}

			sidequest.ID = dbSidequest.ID
			l.Sidequests[sidequest.Name] = sidequest
			l.SidequestsID[sidequest.ID] = sidequest

			err = l.seedSubquests(qtx, sidequest)
			if err != nil {
				return h.NewErr(sidequest.Error(), err)
			}
		}
		return nil
	})
}

func (l *Lookup) seedSubquests(qtx *database.Queries, sidequest Sidequest) error {
	for _, subquest := range sidequest.Subquests {
		var err error
		subquest.Type = database.QuestTypeSubquest
		subquest.SidequestID = sidequest.ID

		subquest.Quest, err = seedObjAssignID(qtx, subquest.Quest, l.seedQuest)
		if err != nil {
			return h.NewErr(subquest.Error(), err)
		}

		dbSubquest, err := qtx.CreateSubquest(context.Background(), database.CreateSubquestParams{
			DataHash:     generateDataHash(subquest),
			QuestID:      subquest.Quest.ID,
			SidequestID:  subquest.SidequestID,
		})
		if err != nil {
			return h.NewErr(subquest.Error(), err, "couldn't create subquest")
		}

		subquest.ID = dbSubquest.ID
		l.Subquests[subquest.Name] = subquest
		l.SubquestsID[subquest.ID] = subquest
	}

	return nil
}

func (l *Lookup) seedSidequestsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/sidequests.json"

	var sidequests []Sidequest
	err := loadJSONFile(string(srcPath), &sidequests)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonSidequest := range sidequests {
			sidequest, err := GetResource(jsonSidequest.Name, l.Sidequests)
			if err != nil {
				return err
			}

			if sidequest.Completion != nil {
				sidequest.Completion, err = seedObjPtrAssignFK(qtx, sidequest.Completion, l.seedQuestCompletion)
				if err != nil {
					return err
				}

				err = qtx.UpdateQuest(context.Background(), database.UpdateQuestParams{
					DataHash: generateDataHash(sidequest.Quest),
					CompletionID: h.ObjPtrToNullInt32ID(sidequest.Completion),
					ID: sidequest.Quest.ID,
				})
			}

			for _, jsonSubquest := range sidequest.Subquests {
				subquest, err := GetResource(jsonSubquest.Name, l.Subquests)
				if err != nil {
					return h.NewErr(sidequest.Error(), err)
				}

				subquest.Completion, err = seedObjPtrAssignFK(qtx, subquest.Completion, l.seedQuestCompletion)
				if err != nil {
					return err
				}

				err = qtx.UpdateQuest(context.Background(), database.UpdateQuestParams{
					DataHash: generateDataHash(subquest.Quest),
					CompletionID: h.ObjPtrToNullInt32ID(subquest.Completion),
					ID: subquest.Quest.ID,
				})
			}
		}

		return nil
	})
}