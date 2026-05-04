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
	Category                string           `json:"category"`
	IsPermanent             bool             `json:"is_permanent"`
	Effect                  string           `json:"effect"`
	Visualization           *string          `json:"visualization"`
	RelatedStats            []string         `json:"related_stats"`
	RemovedStatusConditions []string         `json:"removed_status_conditions"`
	AddedElemResist         *ElementalResist `json:"added_elem_resist"`
	CtbOnInfliction         *InflictedDelay  `json:"ctb_on_infliction"`
	NullifyArmored          *string          `json:"nullify_armored"`
	StatChanges             []StatChange     `json:"stat_changes"`
	ModifierChanges         []ModifierChange `json:"modifier_changes"`
}

func (s StatusCondition) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Name,
		s.Category,
		s.IsPermanent,
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
				Category:       database.StatusConditionCategory(condition.Category),
				IsPermanent:    condition.IsPermanent,
				Effect:         condition.Effect,
				Visualization:  h.GetNullString(condition.Visualization),
				NullifyArmored: database.ToNullNullifyArmored(condition.NullifyArmored),
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

			condition.CtbOnInfliction, err = seedObjPtrAssignFK(qtx, condition.CtbOnInfliction, l.seedInflictedDelay)
			if err != nil {
				return h.NewErr(condition.Error(), err)
			}

			err = qtx.UpdateStatusCondition(context.Background(), database.UpdateStatusConditionParams{
				DataHash:          generateDataHash(condition),
				AddedElemResistID: h.ObjPtrToNullInt32ID(condition.AddedElemResist),
				InflictedDelayID:  h.ObjPtrToNullInt32ID(condition.CtbOnInfliction),
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

func (l *Lookup) loop3SeedStatusConditions(qtx *database.Queries, ctx context.Context) error {
	statusses, err := l.extractStatusConditions()
	if err != nil {
		return err
	}

	params := database.CreateStatusConditionBulkParams{
		DataHash:          make([]string, len(statusses)),
		Name:              make([]string, len(statusses)),
		Category:          make([]database.StatusConditionCategory, len(statusses)),
		IsPermanent:       make([]bool, len(statusses)),
		Effect:            make([]string, len(statusses)),
		Visualization:     make([]sql.NullString, len(statusses)),
		NullifyArmored:    make([]database.NullNullifyArmored, len(statusses)),
		AddedElemResistID: make([]sql.NullInt32, len(statusses)),
		InflictedDelayID:  make([]sql.NullInt32, len(statusses)),
	}

	for i, s := range statusses {
		params.DataHash[i] = generateDataHash(s)
		params.Name[i] = s.Name
		params.Category[i] = database.StatusConditionCategory(s.Category)
		params.IsPermanent[i] = s.IsPermanent
		params.Effect[i] = s.Effect
		params.Visualization[i] = h.GetNullString(s.Visualization)
		params.NullifyArmored[i] = database.ToNullNullifyArmored(s.NullifyArmored)
		params.AddedElemResistID[i] = h.ObjPtrToNullInt32ID(s.AddedElemResist)
		params.InflictedDelayID[i] = h.ObjPtrToNullInt32ID(s.CtbOnInfliction)
	}

	dbRows, err := qtx.CreateStatusConditionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create status conditions: %v", err)
	}

	for i, row := range dbRows {
		statusses[i].ID = row.ID
		l.json.statusConditions[i].ID = row.ID
		l.StatusConditions[statusses[i].Name] = statusses[i]
		l.StatusConditionsID[row.ID] = statusses[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractStatusConditions() ([]StatusCondition, error) {
	statusses := []StatusCondition{}
	var err error

	for i := range l.json.statusConditions {
		status := &l.json.statusConditions[i]

		if status.AddedElemResist != nil {
			status.AddedElemResist.ID, err = l.getHashID(status.AddedElemResist)
			if err != nil {
				return nil, err
			}
		}

		if status.CtbOnInfliction != nil {
			status.CtbOnInfliction.ID, err = l.getHashID(status.CtbOnInfliction)
			if err != nil {
				return nil, err
			}
		}

		statusses = append(statusses, *status)
	}

	return dedupeRows(statusses, l.Hashes), nil
}

func (l *Lookup) completeStatusConditions() error {
	for i := range l.json.statusConditions {
		status := &l.json.statusConditions[i]

		err := assignIDs(l, status.StatChanges)
		if err != nil {
			return err
		}

		err = assignIDs(l, status.ModifierChanges)
		if err != nil {
			return err
		}

		l.StatusConditions[status.Name] = *status
		l.StatusConditionsID[status.ID] = *status
	}

	return nil
}

func (l *Lookup) getStatusConditionModifierChanges(sc StatusCondition) ([]ModifierChange, error) {
	return sc.ModifierChanges, nil
}

func (l *Lookup) getStatusConditionRelatedStats(sc StatusCondition) ([]Stat, error) {
	return getResources(sc.RelatedStats, l.Stats)
}

func (l *Lookup) getStatusConditionRemovedConditions(sc StatusCondition) ([]StatusCondition, error) {
	return getResources(sc.RemovedStatusConditions, l.StatusConditions)
}

func (l *Lookup) getStatusConditionStatChanges(sc StatusCondition) ([]StatChange, error) {
	return sc.StatChanges, nil
}

func (l *Lookup) seedJuncStatusConditionModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + modifier changes"
	jParams, err := processJunctions(l, desc, l.json.statusConditions, l.getStatusConditionModifierChanges)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsModifierChangesJunctionBulk(ctx, database.CreateStatusConditionsModifierChangesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		ModifierChangeID:  jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncStatusConditionRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + related stats"
	jParams, err := processJunctions(l, desc, l.json.statusConditions, l.getStatusConditionRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsRelatedStatsJunctionBulk(ctx, database.CreateStatusConditionsRelatedStatsJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		StatID:            jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncStatusConditionRemovedConditions(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + removed conditions"
	jParams, err := processJunctions(l, desc, l.json.statusConditions, l.getStatusConditionRemovedConditions)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsRemovedStatusConditionsJunctionBulk(ctx, database.CreateStatusConditionsRemovedStatusConditionsJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		ParentConditionID: jParams.ParentIDs,
		ChildConditionID:  jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncStatusConditionStatChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "status conditions + stat changes"
	jParams, err := processJunctions(l, desc, l.json.statusConditions, l.getStatusConditionStatChanges)
	if err != nil {
		return err
	}

	return qtx.CreateStatusConditionsStatChangesJunctionBulk(ctx, database.CreateStatusConditionsStatChangesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		StatusConditionID: jParams.ParentIDs,
		StatChangeID:      jParams.ChildIDs,
	})
}
