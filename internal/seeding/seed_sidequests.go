package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Sidequest struct {
	Name		string		`json:"name"`
	Subquests 	[]Subquest	`json:"subquests"`
}

func(s Sidequest) ToHashFields() []any {
	return []any{
		s.Name,
	}
}


type Subquest struct {
	//id 				int32
	//dataHash			string
	Name		string		`json:"name"`
}

func(s Subquest) ToHashFields() []any {
	return []any{
		s.Name,
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
			quest, err := qtx.CreateQuest(context.Background(), database.CreateQuestParams{
				DataHash: 	generateDataHash(sidequest),
				Name: 		sidequest.Name,
				Type: 		database.QuestTypeSidequest,
			})
			if err != nil {
				return fmt.Errorf("couldn't create Quest: %s: %v", sidequest.Name, err)
			}

			dbSidequest, err := qtx.CreateSidequest(context.Background(), quest.ID)
			if err != nil {
				return fmt.Errorf("couldn't create Sidequest: %s: %v", sidequest.Name, err)
			}

			for _, subquest := range sidequest.Subquests {
				questSub, err := qtx.CreateQuest(context.Background(), database.CreateQuestParams{
					DataHash: 	generateDataHash(subquest),
					Name: 		subquest.Name,
					Type: 		database.QuestTypeSubquest,
				})
				if err != nil {
					return fmt.Errorf("couldn't create Quest: %s - %s: %v", sidequest.Name, subquest.Name, err)
				}

				err = qtx.CreateSubquest(context.Background(), database.CreateSubquestParams{
					QuestID: 			questSub.ID,
					ParentSidequestID: 	dbSidequest.ID,
				})
				if err != nil {
					return fmt.Errorf("couldn't create Subquest: %s - %s: %v", sidequest.Name, subquest.Name, err)
				}
			}
		}
		return nil
	})
}