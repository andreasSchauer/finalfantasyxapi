package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)



type Sidequest struct {
	ID			int32
	Quest
	Subquests 	[]Subquest `json:"subquests"`
}

func (s Sidequest) ToHashFields() []any {
	return []any{
		s.Quest.ID,
	}
}

type Subquest struct {
	ID			int32
	Quest
	SidequestID int32
}

func (s Subquest) ToHashFields() []any {
	return []any{
		s.Quest.ID,
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

			dbQuest, err := l.seedQuest(qtx, sidequest.Quest)
			if err != nil {
				return err
			}

			sidequest.Quest.ID = dbQuest.ID

			dbSidequest, err := qtx.CreateSidequest(context.Background(), database.CreateSidequestParams{
				DataHash: generateDataHash(sidequest),
				QuestID:  sidequest.Quest.ID,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Sidequest: %s: %v", sidequest.Name, err)
			}

			sidequest.ID = dbSidequest.ID

			err = l.seedSubquests(qtx, sidequest)
			if err != nil {
				return err
			}
		}
		return nil
	})
}


func (l *lookup) seedSubquests(qtx *database.Queries, sidequest Sidequest) error {
	for _, subquest := range sidequest.Subquests {
		subquest.Type = database.QuestTypeSubquest

		dbQuest, err := l.seedQuest(qtx, subquest.Quest)
		if err != nil {
			return err
		}

		subquest.Quest.ID = dbQuest.ID
		subquest.SidequestID = sidequest.ID

		dbSubquest, err := qtx.CreateSubquest(context.Background(), database.CreateSubquestParams{
			DataHash:          generateDataHash(subquest),
			QuestID:           subquest.Quest.ID,
			ParentSidequestID: subquest.SidequestID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create Subquest: %s - %s: %v", sidequest.Name, subquest.Name, err)
		}

		subquest.ID = dbSubquest.ID
		key := createLookupKey(subquest.Quest)
		l.subquests[key] = subquest
	}

	return nil
}