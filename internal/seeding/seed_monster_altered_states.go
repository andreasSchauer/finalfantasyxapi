package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AlteredState struct {
	ID          int32
	MonsterID   int32
	Condition   string           `json:"condition"`
	IsTemporary bool             `json:"is_temporary"`
	Changes     []AltStateChange `json:"changes"`
}

func (a AlteredState) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.MonsterID,
		a.Condition,
		a.IsTemporary,
	}
}

func (a AlteredState) GetID() int32 {
	return a.ID
}

func (a *AlteredState) SetID(id int32) {
	a.ID = id
}

func (a AlteredState) Error() string {
	return fmt.Sprintf("altered state with monster id: %d, is temporary: %t, condition: %s", a.MonsterID, a.IsTemporary, a.Condition)
}

type AltStateChange struct {
	ID               int32
	AlteredStateID   int32
	AlterationType   string            `json:"alteration_type"`
	Distance         *int32            `json:"distance"`
	Properties       []string          `json:"properties"`
	AutoAbilities    []string          `json:"auto_abilities"`
	BaseStats        []BaseStat        `json:"base_stats"`
	ElemResists      []ElementalResist `json:"elem_resists"`
	StatusImmunities []string          `json:"status_immunities"`
	AddedStatus      *InflictedStatus  `json:"added_status"`
}

func (a AltStateChange) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.AlteredStateID,
		a.AlterationType,
		h.DerefOrNil(a.Distance),
		h.ObjPtrToID(a.AddedStatus),
	}
}

func (a AltStateChange) GetID() int32 {
	return a.ID
}

func (a *AltStateChange) SetID(id int32) {
	a.ID = id
}

func (a AltStateChange) Error() string {
	return fmt.Sprintf("alt stat change with altered state id: %d, alteration type: %s", a.AlteredStateID, a.AlterationType)
}

func (l *Lookup) loop2SeedAlteredStates(qtx *database.Queries, ctx context.Context) error {
	states := l.extractMonsterAlteredStates()

	params := database.CreateAlteredStateBulkParams{
		DataHash:    make([]string, len(states)),
		MonsterID:   make([]int32, len(states)),
		Condition:   make([]string, len(states)),
		IsTemporary: make([]bool, len(states)),
	}

	for i, s := range states {
		params.DataHash[i] = generateDataHash(s)
		params.MonsterID[i] = s.MonsterID
		params.Condition[i] = s.Condition
		params.IsTemporary[i] = s.IsTemporary
	}

	dbRows, err := qtx.CreateAlteredStateBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create altered states: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterAlteredStates() []AlteredState {
	states := []AlteredState{}

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		for j := range mon.AlteredStates {
			state := &mon.AlteredStates[j]
			state.MonsterID = mon.ID
			states = append(states, *state)
		}
	}

	return dedupeRows(states, l.Hashes)
}

func (l *Lookup) loop5SeedAltStateChanges(qtx *database.Queries, ctx context.Context) error {
	changes, err := l.extractAltStateChanges()
	if err != nil {
		return err
	}

	params := database.CreateAltStateChangeBulkParams{
		DataHash:       make([]string, len(changes)),
		AlteredStateID: make([]int32, len(changes)),
		AlterationType: make([]database.AlterationType, len(changes)),
		Distance:       make([]sql.NullInt32, len(changes)),
		AddedStatusID:  make([]sql.NullInt32, len(changes)),
	}

	for i, c := range changes {
		params.DataHash[i] = generateDataHash(c)
		params.AlteredStateID[i] = c.AlteredStateID
		params.AlterationType[i] = database.AlterationType(c.AlterationType)
		params.Distance[i] = h.GetNullInt32(c.Distance)
		params.AddedStatusID[i] = h.ObjPtrToNullInt32ID(c.AddedStatus)
	}

	dbRows, err := qtx.CreateAltStateChangeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create alt state changes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAltStateChanges() ([]AltStateChange, error) {
	changes := []AltStateChange{}

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		changesNew, err := l.prepareAltStateChanges(mon.AlteredStates)
		if err != nil {
			return nil, err
		}

		changes = append(changes, changesNew...)
	}

	return dedupeRows(changes, l.Hashes), nil
}

func (l *Lookup) completeAlteredStates(states []AlteredState) error {
	for i := range states {
		state := &states[i]

		err := l.assignID(state)
		if err != nil {
			return err
		}

		err = l.completeAltStateChanges(state.Changes)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) prepareAltStateChanges(states []AlteredState) ([]AltStateChange, error) {
	changes := []AltStateChange{}
	var err error

	for i := range states {
		state := &states[i]

		for j := range state.Changes {
			change := &state.Changes[j]

			change.AlteredStateID, err = l.getHashID(state)
			if err != nil {
				return nil, err
			}

			if change.AddedStatus != nil {
				change.AddedStatus.ID, err = l.getHashID(change.AddedStatus)
				if err != nil {
					return nil, err
				}
			}

			changes = append(changes, *change)
		}
	}

	return changes, nil
}

func (l *Lookup) completeAltStateChanges(changes []AltStateChange) error {
	for i := range changes {
		change := &changes[i]

		err := l.assignID(change)
		if err != nil {
			return err
		}

		err = assignIDs(l, change.BaseStats)
		if err != nil {
			return err
		}

		err = assignIDs(l, change.ElemResists)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) getAltStateChanges() []AltStateChange {
	changes := []AltStateChange{}

	for _, mon := range l.json.monsters {
		for _, state := range mon.AlteredStates {
			changes = append(changes, state.Changes...)
		}
	}

	return changes
}

func (l *Lookup) getAltStateChangeAutoAbilities(c AltStateChange) ([]AutoAbility, error) {
	return getResources(c.AutoAbilities, l.AutoAbilities)
}

func (l *Lookup) getAltStateChangeBaseStats(c AltStateChange) ([]BaseStat, error) {
	return c.BaseStats, nil
}

func (l *Lookup) getAltStateChangeElementalResists(c AltStateChange) ([]ElementalResist, error) {
	return c.ElemResists, nil
}

func (l *Lookup) getAltStateChangeProperties(c AltStateChange) ([]Property, error) {
	return getResources(c.Properties, l.Properties)
}

func (l *Lookup) getAltStateChangeStatusImmunities(c AltStateChange) ([]StatusCondition, error) {
	return getResources(c.StatusImmunities, l.StatusConditions)
}

func (l *Lookup) seedJuncAltStateChangesAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + auto-abilities"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesAutoAbilitiesJunctionBulk(ctx, database.CreateAltStateChangesAutoAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltStateChangeID: jParams.ParentIDs,
		AutoAbilityID:    jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAltStateChangesBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + base stats"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesBaseStatsJunctionBulk(ctx, database.CreateAltStateChangesBaseStatsJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltStateChangeID: jParams.ParentIDs,
		BaseStatID:       jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAltStateChangesElementalResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + elemental resists"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeElementalResists)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesElemResistsJunctionBulk(ctx, database.CreateAltStateChangesElemResistsJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltStateChangeID: jParams.ParentIDs,
		ElemResistID:     jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAltStateChangesProperties(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + properties"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeProperties)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesPropertiesJunctionBulk(ctx, database.CreateAltStateChangesPropertiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AltStateChangeID: jParams.ParentIDs,
		PropertyID:       jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAltStateChangesStatusImmunities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "alt state changes + status immunities"
	jParams, err := processJunctions(l, desc, l.getAltStateChanges(), l.getAltStateChangeStatusImmunities)
	if err != nil {
		return err
	}

	return qtx.CreateAltStateChangesStatusImmunitiesJunctionBulk(ctx, database.CreateAltStateChangesStatusImmunitiesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		AltStateChangeID:  jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}
