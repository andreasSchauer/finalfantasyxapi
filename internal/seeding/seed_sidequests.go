package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Sidequest struct {
	ID int32
	Quest
	Completion *QuestCompletion `json:"completion"`
	Subquests  []Subquest       `json:"subquests"`
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

type Subquest struct {
	ID int32
	Quest
	SidequestID int32
	Completions []QuestCompletion `json:"completion"`
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

func (qc QuestCompletion) Error() string {
	return fmt.Sprintf("quest completion with quest id: %d, reward item: %s, amount: %d, condition: %s", qc.QuestID, qc.Reward.ItemName, qc.Reward.Amount, qc.Condition)
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
		h.DerefOrNil(cl.Notes),
	}
}

func (cl CompletionLocation) Error() string {
	return fmt.Sprintf("completion location %s, with completion id: %d, notes: %v", cl.LocationArea, cl.CompletionID, h.DerefOrNil(cl.Notes))
}

func (l *Lookup) seedSidequests(db *database.Queries, dbConn *sql.DB) error {
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
				return h.GetErr(sidequest.Error(), err)
			}

			dbSidequest, err := qtx.CreateSidequest(context.Background(), database.CreateSidequestParams{
				DataHash: generateDataHash(sidequest),
				QuestID:  sidequest.Quest.ID,
			})
			if err != nil {
				return h.GetErr(sidequest.Error(), err, "couldn't create sidequest")
			}

			sidequest.ID = dbSidequest.ID
			key := createLookupKey(sidequest.Quest)
			l.sidequests[key] = sidequest

			err = l.seedSubquests(qtx, sidequest)
			if err != nil {
				return h.GetErr(sidequest.Error(), err)
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
			return h.GetErr(subquest.Error(), err)
		}

		dbSubquest, err := qtx.CreateSubquest(context.Background(), database.CreateSubquestParams{
			DataHash:          generateDataHash(subquest),
			QuestID:           subquest.Quest.ID,
			ParentSidequestID: subquest.SidequestID,
		})
		if err != nil {
			return h.GetErr(subquest.Error(), err, "couldn't create subquest")
		}

		subquest.ID = dbSubquest.ID
		key := createLookupKey(subquest.Quest)
		l.subquests[key] = subquest
	}

	return nil
}

func (l *Lookup) seedSidequestsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/sidequests.json"

	var sidequests []Sidequest
	err := loadJSONFile(string(srcPath), &sidequests)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonSidequest := range sidequests {
			sidequest, err := l.getSidequest(jsonSidequest.Name)
			if err != nil {
				return err
			}

			if sidequest.Completion != nil {
				err := l.seedQuestCompletionRelationships(qtx, *sidequest.Completion, sidequest.Quest)
				if err != nil {
					return h.GetErr(sidequest.Error(), err)
				}
			}

			for _, jsonSubquest := range sidequest.Subquests {
				subquest, err := l.getSubquest(jsonSubquest.Name)
				if err != nil {
					return h.GetErr(sidequest.Error(), err)
				}

				for _, completion := range subquest.Completions {
					err := l.seedQuestCompletionRelationships(qtx, completion, subquest.Quest)
					if err != nil {
						subjects := h.JoinSubjects(sidequest.Error(), subquest.Error())
						return h.GetErr(subjects, err)
					}
				}
			}
		}

		return nil
	})
}

func (l *Lookup) seedQuestCompletionRelationships(qtx *database.Queries, completion QuestCompletion, quest Quest) error {
	var err error

	completion.QuestID, err = assignFK(quest, l.getQuest)
	if err != nil {
		return h.GetErr(completion.Error(), err)
	}

	completion, err = seedObjAssignID(qtx, completion, l.seedQuestCompletion)
	if err != nil {
		return err
	}

	err = l.seedCompletionLocations(qtx, completion)
	if err != nil {
		return h.GetErr(completion.Error(), err)
	}

	return nil
}

func (l *Lookup) seedQuestCompletion(qtx *database.Queries, completion QuestCompletion) (QuestCompletion, error) {
	var err error

	completion.Reward, err = seedObjAssignID(qtx, completion.Reward, l.seedItemAmount)
	if err != nil {
		return QuestCompletion{}, h.GetErr(completion.Error(), err)
	}

	dbCompletion, err := qtx.CreateQuestCompletion(context.Background(), database.CreateQuestCompletionParams{
		DataHash:     generateDataHash(completion),
		QuestID:      completion.QuestID,
		Condition:    completion.Condition,
		ItemAmountID: completion.Reward.ID,
	})
	if err != nil {
		return QuestCompletion{}, h.GetErr(completion.Error(), err, "couldn't create quest completion")
	}
	completion.ID = dbCompletion.ID

	return completion, nil
}

func (l *Lookup) seedCompletionLocations(qtx *database.Queries, completion QuestCompletion) error {
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
			Notes:        h.GetNullString(location.Notes),
		})
		if err != nil {
			return h.GetErr(location.Error(), err, "couldn't create completion location")
		}
	}

	return nil
}
