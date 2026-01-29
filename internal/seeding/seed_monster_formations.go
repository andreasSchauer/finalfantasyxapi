package seeding

import (
	"cmp"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterFormationNew struct {
	MonsterSelection
	TriggerCommands 	[]AbilityReference 	`json:"trigger_commands"`
	Encounters			[]Encounter			`json:"encounters"`
}

type MonsterFormationMap struct {
	MonsterSelections	map[string]MonsterSelectionMap
}

type MonsterSelectionMap struct {
	MonsterSelection
	Encounters		map[string]EncounterMap
	TriggerCommands	[]AbilityReference
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
	Version      		*int32       		`json:"version"`
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

type EncounterLocation struct {
	ID           		int32
	Version      		*int32       		`json:"version"`
	LocationArea 		LocationArea 		`json:"location_area"`
	AreaID       		int32
	Specification   	 *string            `json:"specification"`
	Formations   	 	[]MonsterFormation 	`json:"formations"`
}

func (el EncounterLocation) ToHashFields() []any {
	return []any{
		h.DerefOrNil(el.Version),
		el.AreaID,
		h.DerefOrNil(el.Specification),
	}
}

func (el EncounterLocation) ToKeyFields() []any {
	return []any{
		h.DerefOrNil(el.Version),
		CreateLookupKey(el.LocationArea),
	}
}

func (el EncounterLocation) GetID() int32 {
	return el.ID
}

func (el EncounterLocation) Error() string {
	return fmt.Sprintf("encounter location with version: %v, %s", h.DerefOrNil(el.Version), el.LocationArea)
}

type MonsterFormation struct {
	ID 					int32
	Version				*int32				`json:"version"`
	MonsterSelection
	FormationData
	TriggerCommands 	[]AbilityReference 	`json:"trigger_commands"`
}

func (mf MonsterFormation) ToHashFields() []any {
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
	return fmt.Sprintf("monster formation with version: %d, %s, %s", h.DerefOrNil(mf.Version), mf.MonsterSelection, mf.FormationData)
}

func (mf MonsterFormation) GetResParamsUnnamed() h.ResParamsUnnamed {
	return h.ResParamsUnnamed{
		ID: mf.ID,
	}
}

type MonsterSelection struct {
	ID       int32			`json:"-"`
	SortID	int	`json:"sort_id"`
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

func (fd FormationData) Error() string {
	return fmt.Sprintf("formation data with category: %s, forced ambush: %t, can escape: %t, boss music id: %v, notes: %v", fd.Category, fd.IsForcedAmbush, fd.CanEscape, h.ObjPtrToID(fd.BossMusic), h.DerefOrNil(fd.Notes))
}

type FormationBossSong struct {
	ID               int32	`json:"-"`
	SongID           int32	`json:"-"`
	Song             string `json:"music"`
	CelebrateVictory bool   `json:"celebrate_victory"`
}

func (s FormationBossSong) ToHashFields() []any {
	return []any{
		s.SongID,
		s.CelebrateVictory,
	}
}

func (s FormationBossSong) GetID() int32 {
	return s.ID
}

func (s FormationBossSong) Error() string {
	return fmt.Sprintf("formation boss song %s, celebrate victory: %t", s.Song, s.CelebrateVictory)
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


func (l *Lookup) seedEncounterLocations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_formations.json"

	var encounterLocations []EncounterLocation
	err := loadJSONFile(string(srcPath), &encounterLocations)
	if err != nil {
		return err
	}
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, encounterLocation := range encounterLocations {
			var err error

			locationArea := encounterLocation.LocationArea
			encounterLocation.AreaID, err = assignFK(locationArea, l.Areas)
			if err != nil {
				return h.NewErr(encounterLocation.Error(), err)
			}

			dbEncounterLocation, err := qtx.CreateEncounterLocation(context.Background(), database.CreateEncounterLocationParams{
				DataHash: generateDataHash(encounterLocation),
				Version:  h.GetNullInt32(encounterLocation.Version),
				AreaID:   encounterLocation.AreaID,
				Notes:    h.GetNullString(encounterLocation.Specification),
			})
			if err != nil {
				return h.NewErr(encounterLocation.Error(), err, "couldn't create monster encounter location")
			}

			encounterLocation.ID = dbEncounterLocation.ID
			key := CreateLookupKey(encounterLocation)
			l.EncounterLocations[key] = encounterLocation
		}
		return nil
	})
}

func (l *Lookup) seedMonsterFormationsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_formations.json"

	var encounterLocations []EncounterLocation
	err := loadJSONFile(string(srcPath), &encounterLocations)
	if err != nil {
		return err
	}
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonEncounterLocation := range encounterLocations {
			key := CreateLookupKey(jsonEncounterLocation)
			encounterLocation, err := GetResource(key, l.EncounterLocations)
			if err != nil {
				return err
			}

			for _, monsterFormation := range encounterLocation.Formations {
				var err error
				junction, err := createJunctionSeed(qtx, encounterLocation, monsterFormation, l.seedMonsterFormation)
				if err != nil {
					return h.NewErr(encounterLocation.Error(), err)
				}

				err = qtx.CreateEncounterLocationFormationsJunction(context.Background(), database.CreateEncounterLocationFormationsJunctionParams{
					DataHash:            generateDataHash(junction),
					EncounterLocationID: junction.ParentID,
					MonsterFormationID:  junction.ChildID,
				})
				if err != nil {
					subjects := h.JoinErrSubjects(encounterLocation.Error(), monsterFormation.Error())
					return h.NewErr(subjects, err, "couldn't junction encounter location with monster formation")
				}
			}

		}

		return nil
	})
}

func (l *Lookup) seedMonsterFormation(qtx *database.Queries, formation MonsterFormation) (MonsterFormation, error) {
	var err error

	formation.BossMusic, err = seedObjPtrAssignFK(qtx, formation.BossMusic, l.seedFormationBossSong)
	if err != nil {
		return MonsterFormation{}, h.NewErr(formation.Error(), err)
	}

	dbFormation, err := qtx.CreateMonsterFormation(context.Background(), database.CreateMonsterFormationParams{
		DataHash:       generateDataHash(formation),
		Category:       database.MonsterFormationCategory(formation.Category),
		IsForcedAmbush: formation.IsForcedAmbush,
		CanEscape:      formation.CanEscape,
		BossSongID:     h.ObjPtrToNullInt32ID(formation.BossMusic),
		Notes:          h.GetNullString(formation.Notes),
	})
	if err != nil {
		return MonsterFormation{}, h.NewErr(formation.Error(), err, "couldn't create monster formation")
	}
	formation.ID = dbFormation.ID
	l.MonsterFormationsID[formation.ID] = formation

	err = l.seedFormationMonsterAmounts(qtx, formation)
	if err != nil {
		return MonsterFormation{}, h.NewErr(formation.Error(), err)
	}

	err = l.seedFormationTriggerCommands(qtx, formation)
	if err != nil {
		return MonsterFormation{}, h.NewErr(formation.Error(), err)
	}

	return formation, nil
}

func (l *Lookup) seedFormationBossSong(qtx *database.Queries, bossSong FormationBossSong) (FormationBossSong, error) {
	var err error

	bossSong.SongID, err = assignFK(bossSong.Song, l.Songs)
	if err != nil {
		return FormationBossSong{}, h.NewErr(bossSong.Error(), err)
	}

	dbBossSong, err := qtx.CreateFormationBossSong(context.Background(), database.CreateFormationBossSongParams{
		DataHash:         generateDataHash(bossSong),
		SongID:           bossSong.SongID,
		CelebrateVictory: bossSong.CelebrateVictory,
	})
	if err != nil {
		return FormationBossSong{}, h.NewErr(bossSong.Error(), err, "couldn't create formation boss song")
	}
	bossSong.ID = dbBossSong.ID

	return bossSong, nil
}

func (l *Lookup) seedFormationMonsterAmounts(qtx *database.Queries, formation MonsterFormation) error {
	for _, monsterAmount := range formation.Monsters {
		var err error
		key := CreateLookupKey(monsterAmount)
		monsterAmount.MonsterID, err = assignFK(key, l.Monsters)
		if err != nil {
			return err
		}

		junction, err := createJunctionSeed(qtx, formation, monsterAmount, l.seedMonsterAmount)
		if err != nil {
			return err
		}

		err = qtx.CreateMonsterFormationsMonstersJunction(context.Background(), database.CreateMonsterFormationsMonstersJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterFormationID: junction.ParentID,
			MonsterAmountID:    junction.ChildID,
		})
		if err != nil {
			return h.NewErr(monsterAmount.Error(), err, "couldn't junction monster amount")
		}
	}

	return nil
}

func (l *Lookup) seedFormationTriggerCommands(qtx *database.Queries, formation MonsterFormation) error {
	for _, abilityRef := range formation.TriggerCommands {
		junction, err := createJunction(formation, abilityRef.Untyped(), l.TriggerCommands)
		if err != nil {
			return err
		}

		err = qtx.CreateMonsterFormationsTriggerCommandsJunction(context.Background(), database.CreateMonsterFormationsTriggerCommandsJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterFormationID: junction.ParentID,
			TriggerCommandID:   junction.ChildID,
		})
		if err != nil {
			return h.NewErr(abilityRef.Error(), err, "couldn't junction with trigger command")
		}
	}

	return nil
}

