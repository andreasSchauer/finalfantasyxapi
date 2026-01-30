package seeding

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)



type MonsterFormationNewest struct {
	ID					int32								`json:"-"`
	Version				*int32								`json:"version"`
	MonsterSelection
	FormationData			FormationData					`json:"formation_data"`
	TriggerCommands 		[]FormationTriggerCommand 		`json:"trigger_commands"`
	EncounterLocations		[]EncounterLocationNew			`json:"encounter_locations"`
}

type MonsterFormationNew struct {
	MonsterSelection
	TriggerCommands 	[]FormationTriggerCommand 	`json:"trigger_commands"`
	Encounters			[]Encounter			`json:"encounters"`
}

type MonsterFormationMap struct {
	MonsterSelections	map[string]MonsterSelectionMap
}

type MonsterSelectionMap struct {
	MonsterSelection
	Encounters		map[string]EncounterMap
	TriggerCommands	[]FormationTriggerCommand
}

type EncounterMap struct {
	FormationData
	EncounterLocations map[string]EncounterLocationNew
}


type Encounter struct {
	FormationData		FormationData			`json:"formation_data"`
	EncounterLocations	[]EncounterLocationNew	`json:"encounter_locations"`
}

type EncounterLocationNew struct {
	LocationArea 		LocationArea 		`json:"location_area"`
	Version      		*int32       		`json:"-"`
	Specification   	 *string            `json:"specification"`
}

func (el EncounterLocationNew) ToKeyFields() []any {
	return []any{
		h.DerefOrNil(el.Version),
		CreateLookupKey(el.LocationArea),
	}
}

func (l *Lookup) GetAreaID(el EncounterLocationNew) int32 {
	area, _ := GetResource(el.LocationArea, l.Areas)
	return area.ID
}


type MonsterSelection struct {
	ID       int32			`json:"-"`
	SortID	int	`json:"-"`
	Monsters []MonsterAmount `json:"monsters"`
}

func (ms MonsterSelection) ToHashFields() []any {
	monsters := ms.Monsters

	sort.SliceStable(monsters, func(i, j int) bool { return monsters[i].MonsterID < monsters[j].MonsterID })
	monsterKeys := []any{}

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

type FormationData struct {
	ID             int32				`json:"-"`
	Category       string             `json:"category"`
	IsForcedAmbush bool               `json:"is_forced_ambush"`
	CanEscape      bool               `json:"can_escape"`
	BossMusic      *FormationBossSong `json:"boss_music"`
	Notes          *string            `json:"notes"`
}

func (fd FormationData) GetID() int32 {
	return fd.ID
}

func (fd FormationData) ToHashFields() []any {
	return []any{
		fd.Category,
		fd.IsForcedAmbush,
		fd.CanEscape,
		h.ObjPtrToID(fd.BossMusic),
		h.DerefOrNil(fd.Notes),
	}
}



func (l *Lookup) reformatMonsterFormations2() error {
	const srcPath = "./data/monster_formations_format.json"
	const destPath = "./data/monster_formations_new.json"

	var monsterFormations []MonsterFormationNew
	err := loadJSONFile(string(srcPath), &monsterFormations)
	if err != nil {
		return err
	}

	formationsNew := []MonsterFormationNewest{}

	for _, mfOld := range monsterFormations {
		for i, encounter := range mfOld.Encounters {
			var versionPtr *int32
			if len(mfOld.Encounters) > 1 {
				versionInt := int32(i + 1)
				versionPtr = &versionInt
			}
			newFormation := MonsterFormationNewest{
				Version: versionPtr,
				MonsterSelection: mfOld.MonsterSelection,
				FormationData: encounter.FormationData,
				TriggerCommands: mfOld.TriggerCommands,
				EncounterLocations: encounter.EncounterLocations,
			}
			formationsNew = append(formationsNew, newFormation)
		}
	}

	json, err := json.MarshalIndent(formationsNew, "", "    ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	os.WriteFile(destPath, json, 0666)

	return nil
}


func (l *Lookup) reformatMonsterFormations() error {
	const srcPath = "./data/monster_formations.json"
	const destPath = "./data/monster_formations_format.json"

	var encounterLocations []EncounterLocation
	err := loadJSONFile(string(srcPath), &encounterLocations)
	if err != nil {
		return err
	}

	formationMap := make(map[string]MonsterSelectionMap)
	formationsNew := []MonsterFormationNew{}
	msSortID := 0

	for _, encounterLocationOld := range encounterLocations {
		for _, mfOld := range encounterLocationOld.Formations {
			msKey := generateDataHash(mfOld.MonsterSelection)

			_, ok := formationMap[msKey]
			if !ok {
				msSortID++
				mfOld.MonsterSelection.SortID = msSortID

				formationMap[msKey] = MonsterSelectionMap{
					MonsterSelection: mfOld.MonsterSelection,
					TriggerCommands: mfOld.TriggerCommands,
					Encounters: make(map[string]EncounterMap),
				}
			}

			formationMapEntry := formationMap[msKey]
			ecMap := formationMapEntry.Encounters
			fdKey := generateDataHash(mfOld.FormationData)

			_, ok = ecMap[fdKey]
			if !ok {
				ecMap[fdKey] = EncounterMap{
					FormationData: mfOld.FormationData,
					EncounterLocations: make(map[string]EncounterLocationNew),
				}
			}

			ecMapEntry := ecMap[fdKey]
			eclMap := ecMapEntry.EncounterLocations
			eclKey := CreateLookupKey(encounterLocationOld)

			_, ok = eclMap[eclKey]
			if !ok {
				eclMap[eclKey] = EncounterLocationNew{
					LocationArea: encounterLocationOld.LocationArea,
					Version: encounterLocationOld.Version,
					Specification: encounterLocationOld.Specification,
				}
			}

			ecMapEntry.EncounterLocations = eclMap
			formationMapEntry.Encounters = ecMap
			formationMap[msKey] = formationMapEntry
		}
	}

	for key := range formationMap {
		msMap := formationMap[key]
		encounters := []Encounter{}
		ecMap := msMap.Encounters

		for ecKey := range ecMap {
			entry := ecMap[ecKey]
			encounterLocations := []EncounterLocationNew{}
			eclMap := entry.EncounterLocations
			
			for eclKey := range eclMap {
				ecl := eclMap[eclKey]
				encounterLocations = append(encounterLocations, ecl)
			}
			slices.SortStableFunc(encounterLocations, func (a, b EncounterLocationNew) int {
				aID := l.GetAreaID(a)
				bID := l.GetAreaID(b)

				if aID < bID {
					return -1
				}

				if aID > bID {
					return 1
				}

				return cmp.Compare(*a.Version, *b.Version)
			})

			encounter := Encounter{
				FormationData: 		entry.FormationData,
				EncounterLocations: encounterLocations,
			}

			encounters = append(encounters, encounter)
		}

		mfNew := MonsterFormationNew{
			MonsterSelection: 	msMap.MonsterSelection,
			TriggerCommands: 	msMap.TriggerCommands,
			Encounters: 		encounters,
		}

		formationsNew = append(formationsNew, mfNew)
	}

	slices.SortStableFunc(formationsNew, func (a, b MonsterFormationNew) int {
		return cmp.Compare(a.SortID, b.SortID)
	})

	json, err := json.MarshalIndent(formationsNew, "", "    ")
	if err != nil {
		return fmt.Errorf("couldn't encode JSON: %v", err)
	}

	os.WriteFile(destPath, json, 0666)

	return nil
}
