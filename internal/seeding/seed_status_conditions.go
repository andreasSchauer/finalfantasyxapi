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
