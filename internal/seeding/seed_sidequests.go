package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Sidequest struct {
	ID int32
	Quest
	Completion *QuestCompletion `json:"completion"`
	Subquests  []Subquest      `json:"subquests"`
}

func (s Sidequest) ToHashFields() []any {
	return []any{
		s.Quest.ID,
	}
}

func (s Sidequest) GetID() int32 {
	return s.ID
}

type Subquest struct {
	ID int32
	Quest
	SidequestID 	int32
	Completions  	[]QuestCompletion `json:"completion"`
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

type QuestCompletion struct {
	ID        int32
	QuestID   int32
	Condition string               `json:"condition"`
	Locations []CompletionLocation `json:"locations"`
	Reward    ItemAmount           `json:"reward"`
}

func (qc QuestCompletion) ToHashFields() []any {
	return []any{
		qc.QuestID,
		qc.Condition,
		qc.Reward.ID,
	}
}

func (qc QuestCompletion) GetID() int32 {
	return qc.ID
}

type CompletionLocation struct {
	CompletionID int32
	AreaID       int32
	LocationArea LocationArea `json:"location_area"`
	Notes        *string      `json:"notes"`
}

func (cl CompletionLocation) ToHashFields() []any {
	return []any{
		cl.CompletionID,
		cl.AreaID,
		derefOrNil(cl.Notes),
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
			var err error
			sidequest.Type = database.QuestTypeSidequest

			sidequest.Quest, err = seedObjAssignID(qtx, sidequest.Quest, l.seedQuest)
			if err != nil {
				return err
			}

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
		var err error
		subquest.Type = database.QuestTypeSubquest
		subquest.SidequestID = sidequest.ID

		subquest.Quest, err = seedObjAssignID(qtx, subquest.Quest, l.seedQuest)
		if err != nil {
			return err
		}

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

func (l *lookup) seedSidequestsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/sidequests.json"

	var sidequests []Sidequest
	err := loadJSONFile(string(srcPath), &sidequests)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonSidequest := range sidequests {
			completion := jsonSidequest.Completion
			questKey := Quest{
				Name: jsonSidequest.Name,
				Type: database.QuestTypeSidequest,
			}

			if completion != nil {
				err := l.seedQuestCompletionRelationships(qtx, *completion, questKey)
				if err != nil {
					return err
				}
			}


			for _, jsonSubquest := range jsonSidequest.Subquests {
				for _, completion := range jsonSubquest.Completions {
					questKey := Quest{
						Name: jsonSubquest.Name,
						Type: database.QuestTypeSubquest,
					}
					
					err := l.seedQuestCompletionRelationships(qtx, completion, questKey)
					if err != nil {
						return err
					}
				}
			}
		}

		return nil
	})
}


func (l *lookup) seedQuestCompletionRelationships(qtx *database.Queries, completion QuestCompletion, quest Quest) error {
	var err error

	completion.QuestID, err = assignFK(quest, l.getQuest)
	if err != nil {
		return fmt.Errorf("quest %s: %v", quest.Name, err)
	}

	completion, err = seedObjAssignID(qtx, completion, l.seedQuestCompletion)
	if err != nil {
		return fmt.Errorf("quest %s: %v", quest.Name, err)
	}

	err = l.seedCompletionLocations(qtx, completion)
	if err != nil {
		return err
	}

	return nil
}


func (l *lookup) seedQuestCompletion(qtx *database.Queries, completion QuestCompletion) (QuestCompletion, error) {
	var err error

	completion.Reward, err = seedObjAssignID(qtx, completion.Reward, l.seedItemAmount)
	if err != nil {
		return QuestCompletion{}, err
	}

	dbCompletion, err := qtx.CreateQuestCompletion(context.Background(), database.CreateQuestCompletionParams{
		DataHash:     generateDataHash(completion),
		QuestID:      completion.QuestID,
		Condition:    completion.Condition,
		ItemAmountID: completion.Reward.ID,
	})
	if err != nil {
		return QuestCompletion{}, fmt.Errorf("couldn't create quest completion: %v", err)
	}
	completion.ID = dbCompletion.ID

	return completion, nil
}


func (l *lookup) seedCompletionLocations(qtx *database.Queries, completion QuestCompletion) error {
	for _, location := range completion.Locations {
		var err error

		location.AreaID, err = assignFK(location.LocationArea, l.getArea)
		if err != nil {
			return err
		}
		location.CompletionID = completion.ID

		err = qtx.CreateCompletionLocation(context.Background(), database.CreateCompletionLocationParams{
			DataHash:     generateDataHash(location),
			CompletionID: location.CompletionID,
			AreaID:       location.AreaID,
			Notes:        getNullString(location.Notes),
		})
		if err != nil {
			return fmt.Errorf("couldn't create completion location: %s: %v", createLookupKey(location.LocationArea), err)
		}
	}

	return nil
}
