package seeding

import (
	"fmt"
	"sort"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterFormation struct {
	ID      int32  `json:"-"`
	Version *int32 `json:"version"`
	MonsterSelection
	FormationData   FormationData             `json:"formation_data"`
	TriggerCommands []FormationTriggerCommand `json:"trigger_commands"`
	EncounterAreas  []EncounterArea           `json:"encounter_areas"`
}

func (mf MonsterFormation) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", mf),
		h.DerefOrNil(mf.Version),
		mf.MonsterSelection.ID,
		mf.FormationData.ID,
	}
}

func (mf MonsterFormation) ToKeyFields() []any {
	return []any{
		h.DerefOrNil(mf.Version),
		mf.MonsterSelection.ID,
		mf.FormationData.ID,
	}
}

func (mf MonsterFormation) GetID() int32 {
	return mf.ID
}

func (mf MonsterFormation) Error() string {
	return fmt.Sprintf("monster formation with version: %s, %s, %s", h.PtrToString(mf.Version), mf.MonsterSelection, mf.FormationData)
}

func (mf MonsterFormation) GetResParamsUnnamed() h.ResParamsUnnamed {
	return h.ResParamsUnnamed{
		ID: mf.ID,
	}
}

type MonsterSelection struct {
	ID       int32           `json:"-"`
	Monsters []MonsterAmount `json:"monsters"`
}

func (ms MonsterSelection) ToHashFields() []any {
	monsters := ms.Monsters

	sort.SliceStable(monsters, func(i, j int) bool { return monsters[i].MonsterID < monsters[j].MonsterID })
	monsterKeys := []any{
		fmt.Sprintf("%T", ms),
	}

	for _, mon := range ms.Monsters {
		key := combineFields(mon.ToFormationHashFields())
		monsterKeys = append(monsterKeys, key)
	}

	return monsterKeys
}

func (ms MonsterSelection) GetID() int32 {
	return ms.ID
}

func (ms MonsterSelection) Error() string {
	errs := []string{}

	for _, ma := range ms.Monsters {
		errs = append(errs, ma.Error())
	}

	return strings.Join(errs, " | ")
}
