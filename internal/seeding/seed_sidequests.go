package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Sidequest struct {
	Name		string				`json:"name"`
	Type		database.QuestType
	Subquests 	[]Subquest			`json:"subquests"`
}

func(s Sidequest) ToHashFields() []any {
	return []any{
		s.Name,
		s.Type,
	}
}


type Subquest struct {
	//id 				int32
	//dataHash			string
	Name		string				`json:"name"`
	Type		database.QuestType
}

func(s Subquest) ToHashFields() []any {
	return []any{
		s.Name,
		s.Type,
	}
}


func seedSidequests(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/sidequests.json"

	var sidequests []Sidequest
	err := loadJSONFile(string(srcPath), &sidequests)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, sidequest := range sidequests {
			sidequest.Type = database.QuestTypeSidequest
			
			quest, err := qtx.CreateQuest(context.Background(), database.CreateQuestParams{
				DataHash: 	generateDataHash(sidequest),
				Name: 		sidequest.Name,
				Type: 		sidequest.Type,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Quest: %s: %v", sidequest.Name, err)
			}

			dbSidequest, err := qtx.CreateSidequest(context.Background(), database.CreateSidequestParams{
				DataHash: 	manualDataHash([]any{ quest.ID }),
				QuestID: 	quest.ID,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Sidequest: %s: %v", sidequest.Name, err)
			}

			err = seedSubquests(qtx, sidequest, dbSidequest.ID)
			if err != nil {
				return err
			}
		}
		return nil
	})
}


func seedSubquests(qtx *database.Queries, sidequest Sidequest, sidequestID int32) error {
	for _, subquest := range sidequest.Subquests {
		subquest.Type = database.QuestTypeSubquest
		
		quest, err := qtx.CreateQuest(context.Background(), database.CreateQuestParams{
			DataHash: 	generateDataHash(subquest),
			Name: 		subquest.Name,
			Type: 		subquest.Type,
		})
		if err != nil {
			return fmt.Errorf("couldn't create Quest: %s - %s: %v", sidequest.Name, subquest.Name, err)
		}

		err = qtx.CreateSubquest(context.Background(), database.CreateSubquestParams{
			DataHash: 			manualDataHash([]any{ quest.ID, sidequestID }),
			QuestID: 			quest.ID,
			ParentSidequestID: 	sidequestID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create Subquest: %s - %s: %v", sidequest.Name, subquest.Name, err)
		}
	}

	return nil
}