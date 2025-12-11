package seeding

import (
	"context"
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
		a.MonsterID,
		a.Condition,
		a.IsTemporary,
	}
}

func (a AlteredState) GetID() int32 {
	return a.ID
}

func (a AlteredState) Error() string {
	return fmt.Sprintf("altered state with monster id: %d, is temporary: %t, condition: %s", a.MonsterID, a.IsTemporary, a.Condition)
}

type AltStateChange struct {
	ID               int32
	AlteredStateID   int32
	AlterationType   string             `json:"alteration_type"`
	Distance         *int32             `json:"distance"`
	Properties       *[]string          `json:"properties"`
	AutoAbilities    *[]string          `json:"auto_abilities"`
	BaseStats        *[]BaseStat        `json:"base_stats"`
	ElemResists      *[]ElementalResist `json:"elem_resists"`
	StatusImmunities *[]string          `json:"status_immunities"`
	AddedStatus   	 *InflictedStatus 	`json:"added_status"`
}

func (a AltStateChange) ToHashFields() []any {
	return []any{
		a.AlteredStateID,
		a.AlterationType,
		h.DerefOrNil(a.Distance),
		h.ObjPtrToID(a.AddedStatus),
	}
}

func (a AltStateChange) GetID() int32 {
	return a.ID
}

func (a AltStateChange) Error() string {
	return fmt.Sprintf("alt stat change with altered state id: %d, alteration type: %s", a.AlteredStateID, a.AlterationType)
}

func (l *Lookup) seedAlteredStates(qtx *database.Queries, monster Monster) error {
	for _, state := range monster.AlteredStates {
		var err error
		state.MonsterID = monster.ID

		state, err = seedObjAssignID(qtx, state, l.seedAlteredState)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedAlteredState(qtx *database.Queries, state AlteredState) (AlteredState, error) {
	dbAlteredState, err := qtx.CreateAlteredState(context.Background(), database.CreateAlteredStateParams{
		DataHash:    generateDataHash(state),
		MonsterID:   state.MonsterID,
		Condition:   state.Condition,
		IsTemporary: state.IsTemporary,
	})
	if err != nil {
		return AlteredState{}, h.GetErr(state.Error(), err, "couldn't create altered state")
	}

	state.ID = dbAlteredState.ID

	err = l.seedAltStateChanges(qtx, state)
	if err != nil {
		return AlteredState{}, h.GetErr(state.Error(), err)
	}

	return state, nil
}

func (l *Lookup) seedAltStateChanges(qtx *database.Queries, state AlteredState) error {
	for _, change := range state.Changes {
		var err error
		change.AlteredStateID = state.ID

		change, err = seedObjAssignID(qtx, change, l.seedAltStateChange)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedAltStateChange(qtx *database.Queries, change AltStateChange) (AltStateChange, error) {
	var err error
	change.AddedStatus, err = seedObjPtrAssignFK(qtx, change.AddedStatus, l.seedInflictedStatus)
	if err != nil {
		return AltStateChange{}, h.GetErr(change.Error(), err)
	}

	dbAltStateChange, err := qtx.CreateAltStateChange(context.Background(), database.CreateAltStateChangeParams{
		DataHash:       generateDataHash(change),
		AlteredStateID: change.AlteredStateID,
		AlterationType: database.AlterationType(change.AlterationType),
		Distance:       h.GetNullInt32(change.Distance),
		AddedStatusID: 	h.ObjPtrToNullInt32ID(change.AddedStatus),
	})
	if err != nil {
		return AltStateChange{}, h.GetErr(change.Error(), err, "couldn't create alt state change")
	}

	change.ID = dbAltStateChange.ID

	err = l.seedAltStateChangeJunctions(qtx, change)
	if err != nil {
		return AltStateChange{}, h.GetErr(change.Error(), err)
	}

	return change, nil
}

func (l *Lookup) seedAltStateChangeJunctions(qtx *database.Queries, change AltStateChange) error {
	functions := []func(*database.Queries, AltStateChange) error{
		l.seedAltStateChangeProperties,
		l.seedAltStateChangeAutoAbilities,
		l.seedAltStateBaseStats,
		l.seedAltStateElemResists,
		l.seedAltStateChangeStatusImmunities,
	}

	for _, function := range functions {
		err := function(qtx, change)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedAltStateChangeProperties(qtx *database.Queries, change AltStateChange) error {
	if change.Properties == nil {
		return nil
	}

	for _, propertyStr := range *change.Properties {
		junction, err := createJunction(change, propertyStr, l.Properties)
		if err != nil {
			return err
		}

		err = qtx.CreateAltStateChangesPropertiesJunction(context.Background(), database.CreateAltStateChangesPropertiesJunctionParams{
			DataHash:         generateDataHash(junction),
			AltStateChangeID: junction.ParentID,
			PropertyID:       junction.ChildID,
		})
		if err != nil {
			return h.GetErr(propertyStr, err, "couldn't junction property")
		}
	}

	return nil
}

func (l *Lookup) seedAltStateChangeAutoAbilities(qtx *database.Queries, change AltStateChange) error {
	if change.AutoAbilities == nil {
		return nil
	}

	for _, autoAbilityStr := range *change.AutoAbilities {
		junction, err := createJunction(change, autoAbilityStr, l.AutoAbilities)
		if err != nil {
			return err
		}

		err = qtx.CreateAltStateChangesAutoAbilitiesJunction(context.Background(), database.CreateAltStateChangesAutoAbilitiesJunctionParams{
			DataHash:         generateDataHash(junction),
			AltStateChangeID: junction.ParentID,
			AutoAbilityID:    junction.ChildID,
		})
		if err != nil {
			return h.GetErr(autoAbilityStr, err, "couldn't junction auto-ability")
		}
	}

	return nil
}

func (l *Lookup) seedAltStateBaseStats(qtx *database.Queries, change AltStateChange) error {
	if change.BaseStats == nil {
		return nil
	}

	for _, baseStat := range *change.BaseStats {
		junction, err := createJunctionSeed(qtx, change, baseStat, l.seedBaseStat)
		if err != nil {
			return err
		}

		err = qtx.CreateAltStateChangesBaseStatsJunction(context.Background(), database.CreateAltStateChangesBaseStatsJunctionParams{
			DataHash:         generateDataHash(junction),
			AltStateChangeID: junction.ParentID,
			BaseStatID:       junction.ChildID,
		})
		if err != nil {
			return h.GetErr(baseStat.Error(), err, "couldn't junction base stat")
		}
	}

	return nil
}

func (l *Lookup) seedAltStateElemResists(qtx *database.Queries, change AltStateChange) error {
	if change.ElemResists == nil {
		return nil
	}

	for _, elemResist := range *change.ElemResists {
		junction, err := createJunctionSeed(qtx, change, elemResist, l.seedElementalResist)
		if err != nil {
			return err
		}

		err = qtx.CreateAltStateChangesElemResistsJunction(context.Background(), database.CreateAltStateChangesElemResistsJunctionParams{
			DataHash:         generateDataHash(junction),
			AltStateChangeID: junction.ParentID,
			ElemResistID:     junction.ChildID,
		})
		if err != nil {
			return h.GetErr(elemResist.Error(), err, "couldn't junction elemental resist")
		}
	}

	return nil
}

func (l *Lookup) seedAltStateChangeStatusImmunities(qtx *database.Queries, change AltStateChange) error {
	if change.StatusImmunities == nil {
		return nil
	}

	for _, conditionStr := range *change.StatusImmunities {
		junction, err := createJunction(change, conditionStr, l.StatusConditions)
		if err != nil {
			return err
		}

		err = qtx.CreateAltStateChangesStatusImmunitiesJunction(context.Background(), database.CreateAltStateChangesStatusImmunitiesJunctionParams{
			DataHash:          generateDataHash(junction),
			AltStateChangeID:  junction.ParentID,
			StatusConditionID: junction.ChildID,
		})
		if err != nil {
			return h.GetErr(conditionStr, err, "couldn't junction status immunity")
		}
	}

	return nil
}
