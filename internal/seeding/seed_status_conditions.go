package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type StatusCondition struct {
	ID						int32
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

			condition.ID = dbCondition.ID
			l.statusConditions[condition.Name] = condition
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

			relationShipFunctions := []func(* database.Queries, StatusCondition) error{
				l.createStatusConditionRelatedStats,
				l.createStatusConditionRemovedConditions,
				l.createStatusConditionStatChanges,
				l.createStatusConditionModifierChanges,
			}

			for _, function := range relationShipFunctions {
				err := function(qtx, condition)
				if err != nil {
					return fmt.Errorf("status condition: %s: %v", condition.Name, err)
				}
			}
		}
		return nil
	})
}


func (l *lookup) createStatusConditionRelatedStats(qtx *database.Queries, condition StatusCondition) error {
	for _, jsonStat := range condition.RelatedStats {
		stat, err := l.getStat(jsonStat)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: 	condition.ID,
			ChildID: 	stat.ID,
		}

		err = qtx.CreateStatusConditionStatJunction(context.Background(), database.CreateStatusConditionStatJunctionParams{
			DataHash: 			generateDataHash(junction),
			StatusConditionID: 	junction.ParentID,
			StatID: 			junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create related stats: %v", err)
		}
	}

	return nil
}


func (l *lookup) createStatusConditionRemovedConditions(qtx *database.Queries, condition StatusCondition) error {
	for _, jsonCondition := range condition.RemovedStatusConditions {
		remCondition, err := l.getStatusCondition(jsonCondition)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: 	condition.ID,
			ChildID: 	remCondition.ID,
		}

		err = qtx.CreateStatusConditionSelfJunction(context.Background(), database.CreateStatusConditionSelfJunctionParams{
			DataHash: 			generateDataHash(junction),
			ParentConditionID: 	junction.ParentID,
			ChildConditionID: 	junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create removed conditions: %v", err)
		}
	}

	return nil
}


func (l *lookup) createStatusConditionStatChanges(qtx *database.Queries, condition StatusCondition) error {
	for _, statChange := range condition.StatChanges {
		dbStatChange, err := l.seedStatChange(qtx, statChange)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: 	condition.ID,
			ChildID: 	dbStatChange.ID,
		}

		err = qtx.CreateStatusConditionStatChangeJunction(context.Background(), database.CreateStatusConditionStatChangeJunctionParams{
			DataHash: 			generateDataHash(junction),
			StatusConditionID: 	junction.ParentID,
			StatChangeID: 		junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}


func (l *lookup) createStatusConditionModifierChanges(qtx *database.Queries, condition StatusCondition) error {
	for _, modifierChange := range condition.ModifierChanges {
		dbModifierChange, err := l.seedModifierChange(qtx, modifierChange)
		if err != nil {
			return err
		}

		junction := Junction{
			ParentID: 	condition.ID,
			ChildID: 	dbModifierChange.ID,
		}

		err = qtx.CreateStatusConditionModifierChangeJunction(context.Background(), database.CreateStatusConditionModifierChangeJunctionParams{
			DataHash: 			generateDataHash(junction),
			StatusConditionID: 	junction.ParentID,
			ModifierChangeID: 	junction.ChildID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}