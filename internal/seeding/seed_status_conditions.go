package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type StatusCondition struct {
	//id 		int32
	//dataHash	string
	Name           			string  			`json:"name"`
	Effect         			string  			`json:"effect"`
	RelatedStats			[]string			`json:"related_stats"`
	RemovedStatusConditions	[]string			`json:"removed_status_conditions"`
	NullifyArmored 			*string 			`json:"nullify_armored"`
	StatChanges				[]StatChange		`json:"stat_changes"`
	ModifierChanges			[]ModifierChange	`json:"modifier_changes"`
}

func (s StatusCondition) ToHashFields() []any {
	return []any{
		s.Name,
		s.Effect,
		derefOrNil(s.NullifyArmored),
	}
}


type StatusConditionLookup struct {
	StatusCondition
	ID					int32
}


type StatusConditionStatJunction struct {
	StatusConditionID	int32
	StatID				int32
}

func (s StatusConditionStatJunction) ToHashFields() []any {
	return []any{
		s.StatusConditionID,
		s.StatID,
	}
}


type StatusConditionSelfJunction struct {
	ParentConditionID	int32
	RemovedConditionID	int32
}

func (s StatusConditionSelfJunction) ToHashFields() []any {
	return []any{
		s.ParentConditionID,
		s.RemovedConditionID,
	}
}


type StatusConditionStatChangeJunction struct {
	StatusConditionID	int32
	StatChangeID		int32
}

func (s StatusConditionStatChangeJunction) ToHashFields() []any {
	return []any{
		s.StatusConditionID,
		s.StatChangeID,
	}
}


type StatusConditionModifierChangeJunction struct {
	StatusConditionID	int32
	ModifierID			int32
}

func (s StatusConditionModifierChangeJunction) ToHashFields() []any {
	return []any{
		s.StatusConditionID,
		s.ModifierID,
	}
}



func (l *lookup) seedStatusConditions(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/status_conditions.json"

	var statusConditions []StatusCondition
	err := loadJSONFile(string(srcPath), &statusConditions)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, condition := range statusConditions {
			dbCondition, err := qtx.CreateStatusCondition(context.Background(), database.CreateStatusConditionParams{
				DataHash:       generateDataHash(condition),
				Name:           condition.Name,
				Effect:         condition.Effect,
				NullifyArmored: nullNullifyArmored(condition.NullifyArmored),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Status Condition: %s: %v", condition.Name, err)
			}

			l.statusConditions[condition.Name] = StatusConditionLookup{
				StatusCondition: 	condition,
				ID: 				dbCondition.ID,
			}
		}
		return nil
	})
}



func (l *lookup) createStatusConditionsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/status_conditions.json"

	var statusConditions []StatusCondition
	err := loadJSONFile(string(srcPath), &statusConditions)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCondition := range statusConditions {
			condition, err := l.getStatusCondition(jsonCondition.Name)
			if err != nil {
				return err
			}

			relationshipFunctions := []func(*database.Queries, StatusConditionLookup) error{
				l.createStatusConditionRelatedStats,
				l.createStatusConditionRemovedConditions,
				l.createStatusConditionStatChanges,
				l.createStatusConditionModifierChanges,
			}

			for _, function := range relationshipFunctions {
				err := function(db, condition)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}


func (l *lookup) createStatusConditionRelatedStats(qtx *database.Queries, condition StatusConditionLookup) error {
	for _, jsonStat := range condition.RelatedStats {
		stat, err := l.getStat(jsonStat)
		if err != nil {
			return err
		}

		junction := StatusConditionStatJunction{
			StatusConditionID: condition.ID,
			StatID: stat.ID,
		}

		err = qtx.CreateStatusConditionStatJunction(context.Background(), database.CreateStatusConditionStatJunctionParams{
			DataHash: 			generateDataHash(junction),
			StatusConditionID: 	junction.StatusConditionID,
			StatID: 			junction.StatID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}



func (l *lookup) createStatusConditionRemovedConditions(qtx *database.Queries, condition StatusConditionLookup) error {
	for _, jsonCondition := range condition.RemovedStatusConditions {
		remCondition, err := l.getStatusCondition(jsonCondition)
		if err != nil {
			return err
		}

		junction := StatusConditionSelfJunction{
			ParentConditionID: 	condition.ID,
			RemovedConditionID: remCondition.ID,
		}

		err = qtx.CreateStatusConditionSelfJunction(context.Background(), database.CreateStatusConditionSelfJunctionParams{
			DataHash: 			generateDataHash(junction),
			ParentConditionID: 	junction.ParentConditionID,
			ChildConditionID: 	junction.RemovedConditionID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}



func (l *lookup) createStatusConditionStatChanges(qtx *database.Queries, condition StatusConditionLookup) error {
	for _, statChange := range condition.StatChanges {
		dbStatChange, err := l.seedStatChange(qtx, statChange)
		if err != nil {
			return err
		}

		junction := StatusConditionStatChangeJunction{
			StatusConditionID: 	condition.ID,
			StatChangeID: 		dbStatChange.ID,
		}

		err = qtx.CreateStatusConditionStatChangeJunction(context.Background(), database.CreateStatusConditionStatChangeJunctionParams{
			DataHash: 			generateDataHash(junction),
			StatusConditionID: 	junction.StatusConditionID,
			StatChangeID: 		junction.StatChangeID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}



func (l *lookup) createStatusConditionModifierChanges(qtx *database.Queries, condition StatusConditionLookup) error {
	for _, modifierChange := range condition.ModifierChanges {
		dbModifierChange, err := l.seedModifierChange(qtx, modifierChange)
		if err != nil {
			return err
		}

		junction := StatusConditionModifierChangeJunction{
			StatusConditionID: 	condition.ID,
			ModifierID: 		dbModifierChange.ID,
		}

		err = qtx.CreateStatusConditionModifierChangeJunction(context.Background(), database.CreateStatusConditionModifierChangeJunctionParams{
			DataHash: 			generateDataHash(junction),
			StatusConditionID: 	junction.StatusConditionID,
			ModifierChangeID: 	junction.ModifierID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}