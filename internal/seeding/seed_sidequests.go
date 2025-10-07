package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Quest struct {
	Name string `json:"name"`
	Type database.QuestType
}

func (q Quest) ToHashFields() []any {
	return []any{
		q.Name,
		q.Type,
	}
}

type Sidequest struct {
	Quest
	QuestID   int32
	Subquests []Subquest `json:"subquests"`
}

func (s Sidequest) ToHashFields() []any {
	return []any{
		s.QuestID,
	}
}

type Subquest struct {
	//id 				int32
	//dataHash			string
	Quest
	QuestID     int32
	SidequestID int32
}

func (s Subquest) ToHashFields() []any {
	return []any{
		s.QuestID,
		s.SidequestID,
	}
}

func (l *lookup) seedSidequests(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/sidequests.json"

	var sidequests []Sidequest
	err := loadJSONFile(string(srcPath), &sidequests)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, sidequest := range sidequests {
			sidequest.Type = database.QuestTypeSidequest

			dbQuest, err := qtx.CreateQuest(context.Background(), database.CreateQuestParams{
				DataHash: generateDataHash(sidequest.Quest),
				Name:     sidequest.Name,
				Type:     sidequest.Type,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Quest: %s: %v", sidequest.Name, err)
			}

			sidequest.QuestID = dbQuest.ID

			dbSidequest, err := qtx.CreateSidequest(context.Background(), database.CreateSidequestParams{
				DataHash: generateDataHash(sidequest),
				QuestID:  sidequest.QuestID,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Sidequest: %s: %v", sidequest.Name, err)
			}

			err = l.seedSubquests(qtx, sidequest, dbSidequest.ID)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (l *lookup) seedSubquests(qtx *database.Queries, sidequest Sidequest, sidequestID int32) error {
	for _, subquest := range sidequest.Subquests {
		subquest.Type = database.QuestTypeSubquest

		dbQuest, err := qtx.CreateQuest(context.Background(), database.CreateQuestParams{
			DataHash: generateDataHash(subquest.Quest),
			Name:     subquest.Name,
			Type:     subquest.Type,
		})
		if err != nil {
			return fmt.Errorf("couldn't create Quest: %s - %s: %v", sidequest.Name, subquest.Name, err)
		}

		subquest.QuestID = dbQuest.ID
		subquest.SidequestID = sidequestID

		err = qtx.CreateSubquest(context.Background(), database.CreateSubquestParams{
			DataHash:          generateDataHash(subquest),
			QuestID:           subquest.QuestID,
			ParentSidequestID: subquest.SidequestID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create Subquest: %s - %s: %v", sidequest.Name, subquest.Name, err)
		}
	}

	return nil
}
