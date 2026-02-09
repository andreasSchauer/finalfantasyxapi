package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type StatusCondition struct {
	ID                      int32
	Name                    string           `json:"name"`
	Effect                  string           `json:"effect"`
	Visualization           *string          `json:"visualization"`
	RelatedStats            []string         `json:"related_stats"`
	RemovedStatusConditions []string         `json:"removed_status_conditions"`
	AddedElemResist         *ElementalResist `json:"added_elem_resist"`
	NullifyArmored          *string          `json:"h.Nullify_armored"`
	StatChanges             []StatChange     `json:"stat_changes"`
	ModifierChanges         []ModifierChange `json:"modifier_changes"`
}

func (s StatusCondition) ToHashFields() []any {
	return []any{
		s.Name,
		s.Effect,
		h.DerefOrNil(s.Visualization),
		h.ObjPtrToID(s.AddedElemResist),
		h.DerefOrNil(s.NullifyArmored),
	}
}

func (s StatusCondition) GetID() int32 {
	return s.ID
}

func (s StatusCondition) Error() string {
	return fmt.Sprintf("status condition %s", s.Name)
}

func (s StatusCondition) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   s.ID,
		Name: s.Name,
	}
}

func (l *Lookup) seedStatusConditions(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/status_conditions.json"

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
				Visualization:  h.GetNullString(condition.Visualization),
				NullifyArmored: h.NullNullifyArmored(condition.NullifyArmored),
			})
			if err != nil {
				return h.NewErr(condition.Error(), err, "couldn't create status condition")
			}

			condition.ID = dbCondition.ID
			l.StatusConditions[condition.Name] = condition
			l.StatusConditionsID[condition.ID] = condition
		}
		return nil
	})
}

func (l *Lookup) seedStatusConditionsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/status_conditions.json"

	var statusConditions []StatusCondition
	err := loadJSONFile(string(srcPath), &statusConditions)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonCondition := range statusConditions {
			condition, err := GetResource(jsonCondition.Name, l.StatusConditions)
			if err != nil {
				return err
			}

			condition.AddedElemResist, err = seedObjPtrAssignFK(qtx, condition.AddedElemResist, l.seedElementalResist)
			if err != nil {
				return h.NewErr(condition.Error(), err)
			}

			err = qtx.UpdateStatusCondition(context.Background(), database.UpdateStatusConditionParams{
				DataHash:          generateDataHash(condition),
				AddedElemResistID: h.ObjPtrToNullInt32ID(condition.AddedElemResist),
				ID:                condition.ID,
			})
			if err != nil {
				return h.NewErr(condition.Error(), err, "couldn't update status condition")
			}

			relationShipFunctions := []func(*database.Queries, StatusCondition) error{
				l.seedStatusConditionRelatedStats,
				l.seedStatusConditionRemovedConditions,
				l.seedStatusConditionStatChanges,
				l.seedStatusConditionModifierChanges,
			}

			for _, function := range relationShipFunctions {
				err := function(qtx, condition)
				if err != nil {
					return h.NewErr(condition.Error(), err)
				}
			}
		}
		return nil
	})
}

func (l *Lookup) seedStatusConditionRelatedStats(qtx *database.Queries, condition StatusCondition) error {
	for _, jsonStat := range condition.RelatedStats {
		junction, err := createJunction(condition, jsonStat, l.Stats)
		if err != nil {
			return err
		}

		err = qtx.CreateStatusConditionsRelatedStatsJunction(context.Background(), database.CreateStatusConditionsRelatedStatsJunctionParams{
			DataHash:          generateDataHash(junction),
			StatusConditionID: junction.ParentID,
			StatID:            junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonStat, err, "couldn't junction related stat")
		}
	}

	return nil
}

func (l *Lookup) seedStatusConditionRemovedConditions(qtx *database.Queries, condition StatusCondition) error {
	for _, jsonCondition := range condition.RemovedStatusConditions {
		junction, err := createJunction(condition, jsonCondition, l.StatusConditions)
		if err != nil {
			return err
		}

		err = qtx.CreateStatusConditionsRemovedStatusConditionsJunction(context.Background(), database.CreateStatusConditionsRemovedStatusConditionsJunctionParams{
			DataHash:          generateDataHash(junction),
			ParentConditionID: junction.ParentID,
			ChildConditionID:  junction.ChildID,
		})
		if err != nil {
			return h.NewErr(jsonCondition, err, "couldn't junction removed status condition")
		}
	}

	return nil
}

func (l *Lookup) seedStatusConditionStatChanges(qtx *database.Queries, condition StatusCondition) error {
	for _, statChange := range condition.StatChanges {
		junction, err := createJunctionSeed(qtx, condition, statChange, l.seedStatChange)
		if err != nil {
			return err
		}

		err = qtx.CreateStatusConditionsStatChangesJunction(context.Background(), database.CreateStatusConditionsStatChangesJunctionParams{
			DataHash:          generateDataHash(junction),
			StatusConditionID: junction.ParentID,
			StatChangeID:      junction.ChildID,
		})
		if err != nil {
			return h.NewErr(statChange.Error(), err, "couldn't junction stat change")
		}
	}

	return nil
}

func (l *Lookup) seedStatusConditionModifierChanges(qtx *database.Queries, condition StatusCondition) error {
	for _, modifierChange := range condition.ModifierChanges {
		junction, err := createJunctionSeed(qtx, condition, modifierChange, l.seedModifierChange)
		if err != nil {
			return err
		}

		err = qtx.CreateStatusConditionsModifierChangesJunction(context.Background(), database.CreateStatusConditionsModifierChangesJunctionParams{
			DataHash:          generateDataHash(junction),
			StatusConditionID: junction.ParentID,
			ModifierChangeID:  junction.ChildID,
		})
		if err != nil {
			return h.NewErr(modifierChange.Error(), err, "couldn't junction modifier change")
		}
	}

	return nil
}
