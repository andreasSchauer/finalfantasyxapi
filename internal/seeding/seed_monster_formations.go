package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterFormation struct {
	ID      int32  `json:"-"`
	Version *int32 `json:"version"`
	MonsterSelection
	FormationData      FormationData             `json:"formation_data"`
	TriggerCommands    []FormationTriggerCommand `json:"trigger_commands"`
	EncounterLocations []EncounterLocation       `json:"encounter_locations"`
}

func (mf MonsterFormation) ToHashFields() []any {
	return []any{
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
	return fmt.Sprintf("monster formation with version: %d, %s, %s", h.DerefOrNil(mf.Version), mf.MonsterSelection, mf.FormationData)
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
	ID             int32              `json:"-"`
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

type EncounterLocation struct {
	ID            int32
	LocationArea  LocationArea `json:"location_area"`
	AreaID        int32
	Specification *string `json:"specification"`
}

func (el EncounterLocation) ToHashFields() []any {
	return []any{
		el.AreaID,
		h.DerefOrNil(el.Specification),
	}
}

func (el EncounterLocation) ToKeyFields() []any {
	return []any{
		CreateLookupKey(el.LocationArea),
		el.Specification,
	}
}

func (el EncounterLocation) GetID() int32 {
	return el.ID
}

func (el EncounterLocation) Error() string {
	return fmt.Sprintf("encounter location with %s, specification: %s", el.LocationArea, h.DerefOrNil(el.Specification))
}

func (el EncounterLocation) GetLocationArea() LocationArea {
	return el.LocationArea
}

type FormationTriggerCommand struct {
	ID int32
	AbilityReference
	TriggerCommandID int32
	Condition        *string  `json:"condition"`
	UseAmount        *int32   `json:"use_amount"`
	Users            []string `json:"users"`
}

func (tc FormationTriggerCommand) ToHashFields() []any {
	return []any{
		tc.TriggerCommandID,
		h.DerefOrNil(tc.Condition),
		h.DerefOrNil(tc.UseAmount),
	}
}

func (tc FormationTriggerCommand) GetID() int32 {
	return tc.ID
}

func (tc FormationTriggerCommand) Error() string {
	return fmt.Sprintf("formation trigger command with %s", tc.AbilityReference)
}

type FormationBossSong struct {
	ID               int32  `json:"-"`
	SongID           int32  `json:"-"`
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

func (l *Lookup) seedMonsterFormations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_formations.json"

	var monsterFormations []MonsterFormation
	err := loadJSONFile(string(srcPath), &monsterFormations)
	if err != nil {
		return err
	}
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, formation := range monsterFormations {
			var err error

			formation.MonsterSelection, err = seedObjAssignID(qtx, formation.MonsterSelection, l.seedMonsterSelection)
			if err != nil {
				return h.NewErr(formation.Error(), err)
			}

			formation.FormationData, err = seedObjAssignID(qtx, formation.FormationData, l.seedFormationData)
			if err != nil {
				return h.NewErr(formation.Error(), err)
			}

			dbFormation, err := qtx.CreateMonsterFormation(context.Background(), database.CreateMonsterFormationParams{
				DataHash:           generateDataHash(formation),
				Version:            h.GetNullInt32(formation.Version),
				MonsterSelectionID: formation.MonsterSelection.ID,
				FormationDataID:    formation.FormationData.ID,
			})
			if err != nil {
				return h.NewErr(formation.Error(), err, "couldn't create monster formation")
			}

			formation.ID = dbFormation.ID
			key := CreateLookupKey(formation)
			l.MonsterFormations[key] = formation
			l.MonsterFormationsID[formation.ID] = formation
		}

		return nil
	})
}

func (l *Lookup) seedEncounterLocation(qtx *database.Queries, encounterLocation EncounterLocation) (EncounterLocation, error) {
	var err error

	locationArea := encounterLocation.LocationArea
	encounterLocation.AreaID, err = assignFK(locationArea, l.Areas)
	if err != nil {
		return EncounterLocation{}, h.NewErr(encounterLocation.Error(), err)
	}

	dbEncounterLocation, err := qtx.CreateEncounterLocation(context.Background(), database.CreateEncounterLocationParams{
		DataHash:      generateDataHash(encounterLocation),
		AreaID:        encounterLocation.AreaID,
		Specification: h.GetNullString(encounterLocation.Specification),
	})
	if err != nil {
		return EncounterLocation{}, h.NewErr(encounterLocation.Error(), err, "couldn't create monster encounter location")
	}

	encounterLocation.ID = dbEncounterLocation.ID
	key := CreateLookupKey(encounterLocation)
	l.EncounterLocations[key] = encounterLocation

	return encounterLocation, nil
}

func (l *Lookup) seedMonsterFormationsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_formations.json"

	var monsterFormations []MonsterFormation
	err := loadJSONFile(string(srcPath), &monsterFormations)
	if err != nil {
		return err
	}
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for i := range monsterFormations {
			id := int32(i + 1)
			formation, err := GetResourceByID(id, l.MonsterFormationsID)
			if err != nil {
				return err
			}

			err = l.seedFormationEncounterLocations(qtx, formation)
			if err != nil {
				return err
			}

			err = l.seedFormationTriggerCommands(qtx, formation)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (l *Lookup) seedMonsterSelection(qtx *database.Queries, selection MonsterSelection) (MonsterSelection, error) {
	dbSelection, err := qtx.CreateMonsterSelection(context.Background(), generateDataHash(selection))
	if err != nil {
		return MonsterSelection{}, h.NewErr(selection.Error(), err, "couldn't create monster selection")
	}

	selection.ID = dbSelection.ID

	err = l.seedSelectionMonsterAmounts(qtx, selection)
	if err != nil {
		return MonsterSelection{}, err
	}

	return selection, nil
}

func (l *Lookup) seedFormationData(qtx *database.Queries, formationData FormationData) (FormationData, error) {
	var err error

	formationData.BossMusic, err = seedObjPtrAssignFK(qtx, formationData.BossMusic, l.seedFormationBossSong)
	if err != nil {
		return FormationData{}, h.NewErr(formationData.Error(), err)
	}

	dbFormation, err := qtx.CreateFormationData(context.Background(), database.CreateFormationDataParams{
		DataHash:       generateDataHash(formationData),
		Category:       database.MonsterFormationCategory(formationData.Category),
		IsForcedAmbush: formationData.IsForcedAmbush,
		CanEscape:      formationData.CanEscape,
		BossSongID:     h.ObjPtrToNullInt32ID(formationData.BossMusic),
		Notes:          h.GetNullString(formationData.Notes),
	})
	if err != nil {
		return FormationData{}, h.NewErr(formationData.Error(), err, "couldn't create monster formation")
	}
	formationData.ID = dbFormation.ID

	return formationData, nil
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

func (l *Lookup) seedSelectionMonsterAmounts(qtx *database.Queries, selection MonsterSelection) error {
	for _, monsterAmount := range selection.Monsters {
		var err error
		key := CreateLookupKey(monsterAmount)
		monsterAmount.MonsterID, err = assignFK(key, l.Monsters)
		if err != nil {
			return err
		}

		junction, err := createJunctionSeed(qtx, selection, monsterAmount, l.seedMonsterAmount)
		if err != nil {
			return err
		}

		err = qtx.CreateMonsterSelectionsMonstersJunction(context.Background(), database.CreateMonsterSelectionsMonstersJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterSelectionID: junction.ParentID,
			MonsterAmountID:    junction.ChildID,
		})
		if err != nil {
			return h.NewErr(monsterAmount.Error(), err, "couldn't junction monster amount")
		}
	}

	return nil
}

func (l *Lookup) seedFormationEncounterLocations(qtx *database.Queries, formation MonsterFormation) error {
	for _, encounterLocation := range formation.EncounterLocations {
		junction, err := createJunctionSeed(qtx, formation, encounterLocation, l.seedEncounterLocation)
		if err != nil {
			return h.NewErr(formation.Error(), err)
		}

		err = qtx.CreateMonsterFormationsEncounterLocationsJunction(context.Background(), database.CreateMonsterFormationsEncounterLocationsJunctionParams{
			DataHash:            generateDataHash(junction),
			MonsterFormationID:  junction.ParentID,
			EncounterLocationID: junction.ChildID,
		})
		if err != nil {
			subjects := h.JoinErrSubjects(formation.Error(), encounterLocation.Error())
			return h.NewErr(subjects, err, "couldn't junction monster formation with encounter location")
		}
	}

	return nil
}

func (l *Lookup) seedFormationTriggerCommands(qtx *database.Queries, formation MonsterFormation) error {
	for _, triggerCommand := range formation.TriggerCommands {
		junction, err := createJunctionSeed(qtx, formation, triggerCommand, l.seedFormationTriggerCommand)
		if err != nil {
			return err
		}

		err = qtx.CreateMonsterFormationsTriggerCommandsJunction(context.Background(), database.CreateMonsterFormationsTriggerCommandsJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterFormationID: junction.ParentID,
			TriggerCommandID:   junction.ChildID,
		})
		if err != nil {
			return h.NewErr(triggerCommand.Error(), err, "couldn't junction with formation trigger command")
		}
	}

	return nil
}

func (l *Lookup) seedFormationTriggerCommand(qtx *database.Queries, triggerCommand FormationTriggerCommand) (FormationTriggerCommand, error) {
	var err error

	triggerCommand.TriggerCommandID, err = assignFK(triggerCommand.AbilityReference.Untyped(), l.TriggerCommands)
	if err != nil {
		return FormationTriggerCommand{}, err
	}

	dbTriggerCommand, err := qtx.CreateFormationTriggerCommand(context.Background(), database.CreateFormationTriggerCommandParams{
		DataHash:         generateDataHash(triggerCommand),
		TriggerCommandID: triggerCommand.TriggerCommandID,
		Condition:        h.GetNullString(triggerCommand.Condition),
		UseAmount:        h.GetNullInt32(triggerCommand.UseAmount),
	})
	if err != nil {
		return FormationTriggerCommand{}, h.NewErr(triggerCommand.Error(), err, "couldn't create formation trigger command")
	}

	triggerCommand.ID = dbTriggerCommand.ID

	return triggerCommand, nil
}
