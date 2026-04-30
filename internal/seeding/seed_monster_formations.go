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
	ID      			int32  						`json:"-"`
	Version 			*int32 						`json:"version"`
	MonsterSelection
	FormationData   	FormationData             	`json:"formation_data"`
	TriggerCommands 	[]FormationTriggerCommand 	`json:"trigger_commands"`
	EncounterAreas  	[]EncounterArea           	`json:"encounter_areas"`
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

type FormationData struct {
	ID             int32              `json:"-"`
	Category       string             `json:"category"`
	Availability   string			  `json:"availability"`
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
		fmt.Sprintf("%T", fd),
		fd.Category,
		fd.Availability,
		fd.IsForcedAmbush,
		fd.CanEscape,
		h.ObjPtrToID(fd.BossMusic),
		h.DerefOrNil(fd.Notes),
	}
}

func (fd FormationData) Error() string {
	return fmt.Sprintf("formation data with category: %s, forced ambush: %t, can escape: %t, boss music id: %v, notes: %v", fd.Category, fd.IsForcedAmbush, fd.CanEscape, h.ObjPtrToID(fd.BossMusic), h.PtrToString(fd.Notes))
}

type EncounterArea struct {
	ID            int32
	LocationArea  LocationArea `json:"location_area"`
	AreaID        int32
	Specification *string `json:"specification"`
}

func (el EncounterArea) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", el),
		el.AreaID,
		h.DerefOrNil(el.Specification),
	}
}

func (el EncounterArea) ToKeyFields() []any {
	return []any{
		CreateLookupKey(el.LocationArea),
		el.Specification,
	}
}

func (el EncounterArea) GetID() int32 {
	return el.ID
}

func (el EncounterArea) Error() string {
	return fmt.Sprintf("encounter location with %s, specification: %s", el.LocationArea, h.PtrToString(el.Specification))
}

func (el EncounterArea) GetLocationArea() LocationArea {
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
		fmt.Sprintf("%T", tc),
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
		fmt.Sprintf("%T", s),
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
	const srcPath = "data/monster_formations.json"

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

func (l *Lookup) seedEncounterArea(qtx *database.Queries, encounterArea EncounterArea) (EncounterArea, error) {
	var err error

	locationArea := encounterArea.LocationArea
	encounterArea.AreaID, err = assignFK(locationArea, l.Areas)
	if err != nil {
		return EncounterArea{}, h.NewErr(encounterArea.Error(), err)
	}

	dbEncounterArea, err := qtx.CreateEncounterArea(context.Background(), database.CreateEncounterAreaParams{
		DataHash:      generateDataHash(encounterArea),
		AreaID:        encounterArea.AreaID,
		Specification: h.GetNullString(encounterArea.Specification),
	})
	if err != nil {
		return EncounterArea{}, h.NewErr(encounterArea.Error(), err, "couldn't create monster encounter location")
	}

	encounterArea.ID = dbEncounterArea.ID
	key := CreateLookupKey(encounterArea)
	l.EncounterAreas[key] = encounterArea

	return encounterArea, nil
}

func (l *Lookup) seedMonsterFormationsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/monster_formations.json"

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

			err = l.seedFormationEncounterAreas(qtx, formation)
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
		Availability:   database.AvailabilityType(formationData.Availability),
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

func (l *Lookup) seedFormationEncounterAreas(qtx *database.Queries, formation MonsterFormation) error {
	for _, encounterArea := range formation.EncounterAreas {
		junction, err := createJunctionSeed(qtx, formation, encounterArea, l.seedEncounterArea)
		if err != nil {
			return h.NewErr(formation.Error(), err)
		}

		err = qtx.CreateMonsterFormationsEncounterAreasJunction(context.Background(), database.CreateMonsterFormationsEncounterAreasJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterFormationID: junction.ParentID,
			EncounterAreaID:    junction.ChildID,
		})
		if err != nil {
			subjects := h.JoinErrSubjects(formation.Error(), encounterArea.Error())
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

	err = l.seedFormationTriggerCommandUsers(qtx, triggerCommand)
	if err != nil {
		return FormationTriggerCommand{}, h.NewErr(triggerCommand.Error(), err, "couldn't junction users to formation trigger command")
	}

	return triggerCommand, nil
}

func (l *Lookup) seedFormationTriggerCommandUsers(qtx *database.Queries, triggerCommand FormationTriggerCommand) error {
	for _, user := range triggerCommand.Users {
		junction, err := createJunction(triggerCommand, user, l.CharClasses)
		if err != nil {
			return err
		}

		err = qtx.CreateFormationTriggerCommandsUsersJunction(context.Background(), database.CreateFormationTriggerCommandsUsersJunctionParams{
			DataHash:         generateDataHash(junction),
			TriggerCommandID: junction.ParentID,
			CharacterClassID: junction.ChildID,
		})
	}

	return nil
}

func (l *Lookup) loop5SeedMonsterFormations(qtx *database.Queries, ctx context.Context) error {
	formations, err := l.extractMonsterFormations()
	if err != nil {
		return err
	}

	params := database.CreateMonsterFormationBulkParams{
		DataHash:   		make([]string, len(formations)),
		Version: 			make([]sql.NullInt32, len(formations)),
		MonsterSelectionID: make([]int32, len(formations)),
		FormationDataID: 	make([]int32, len(formations)),
	}

	for i, mf := range formations {
		params.DataHash[i] = generateDataHash(mf)
		params.Version[i] = h.GetNullInt32(mf.Version)
		params.MonsterSelectionID[i] = mf.MonsterSelection.ID
		params.FormationDataID[i] = mf.FormationData.ID
	}

	dbRows, err := qtx.CreateMonsterFormationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster formations: %v", err)
	}

	for i, row := range dbRows {
		formations[i].ID = row.ID
		l.json.monsterFormations[i].ID = row.ID
		key := CreateLookupKey(formations[i])
		l.MonsterFormations[key] = formations[i]
		l.MonsterFormationsID[row.ID] = formations[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterFormations() ([]MonsterFormation, error) {
	formations := []MonsterFormation{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		mf.FormationData.ID, err = l.getHashID(mf.FormationData)
		if err != nil {
			return nil, err
		}

		mf.MonsterSelection.ID, err = l.getHashID(mf.MonsterSelection)
		if err != nil {
			return nil, err
		}

		formations = append(formations, *mf)
	}

	return dedupeRows(formations, l.Hashes), nil
}


func (l *Lookup) loop1SeedMonsterSelections(qtx *database.Queries, ctx context.Context) error {
	selections, err := l.extractMonsterSelections()
	if err != nil {
		return err
	}

	dataHashes := make([]string, len(selections))

	for i, s := range selections {
		dataHashes[i] = generateDataHash(s)
	}

	dbRows, err := qtx.CreateMonsterSelectionBulk(ctx, dataHashes)
	if err != nil {
		return fmt.Errorf("couldn't create monster selections: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterSelections() ([]MonsterSelection, error) {
	selections := []MonsterSelection{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		for j := range mf.MonsterSelection.Monsters {
			ma := &mf.MonsterSelection.Monsters[j]

			key := CreateLookupKey(*ma)
			ma.MonsterID, err = assignFK(key, l.Monsters)
			if err != nil {
				return nil, err
			}
		}

		selections = append(selections, mf.MonsterSelection)
	}

	return dedupeRows(selections, l.Hashes), nil
}



func (l *Lookup) loop4SeedFormationData(qtx *database.Queries, ctx context.Context) error {
	data, err := l.extractFormationData()
	if err != nil {
		return err
	}

	params := database.CreateFormationDataBulkParams{
		DataHash:   	make([]string, len(data)),
		Category: 		make([]database.MonsterFormationCategory, len(data)),
		Availability: 	make([]database.AvailabilityType, len(data)),
		IsForcedAmbush: make([]bool, len(data)),
		CanEscape: 		make([]bool, len(data)),
		BossSongID: 	make([]sql.NullInt32, len(data)),
		Notes: 			make([]sql.NullString, len(data)),
	}

	for i, d := range data {
		params.DataHash[i] = generateDataHash(d)
		params.Category[i] = database.MonsterFormationCategory(d.Category)
		params.Availability[i] = database.AvailabilityType(d.Availability)
		params.IsForcedAmbush[i] = d.IsForcedAmbush
		params.CanEscape[i] = d.CanEscape
		params.BossSongID[i] = h.ObjPtrToNullInt32ID(d.BossMusic)
		params.Notes[i] = h.GetNullString(d.Notes)
	}

	dbRows, err := qtx.CreateFormationDataBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create formation data: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractFormationData() ([]FormationData, error) {
	data := []FormationData{}
	var err error

	for i := range l.json.monsterFormations {
		mfData := &l.json.monsterFormations[i].FormationData

		if mfData.BossMusic != nil {
			mfData.BossMusic.ID, err = l.getHashID(mfData.BossMusic)
			if err != nil {
				return nil, err
			}
		}

		data = append(data, *mfData)
	}

	return dedupeRows(data, l.Hashes), nil
}


func (l *Lookup) loop2SeedMonsterAmounts(qtx *database.Queries, ctx context.Context) error {
	mas, err := l.extractMonsterAmounts()
	if err != nil {
		return err
	}

	params := database.CreateMonsterAmountBulkParams{
		DataHash:  make([]string, len(mas)),
		MonsterID: make([]int32, len(mas)),
		Amount:    make([]int32, len(mas)),
	}

	for i, c := range mas {
		params.DataHash[i] = generateDataHash(c)
		params.MonsterID[i] = c.MonsterID
		params.Amount[i] = c.Amount
	}

	dbRows, err := qtx.CreateMonsterAmountBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster amounts: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterAmounts() ([]MonsterAmount, error) {
	mas := []MonsterAmount{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		for j := range mf.Monsters {
			ma := &mf.Monsters[j]

			key := CreateLookupKey(*ma)
			ma.MonsterID, err = assignFK(key, l.Monsters)
			if err != nil {
				return nil, err
			}

			mas = append(mas, *ma)
		}
	}

	return dedupeRows(mas, l.Hashes), nil
}


func (l *Lookup) loop3SeedFormationBossSongs(qtx *database.Queries, ctx context.Context) error {
	songs, err := l.extractFormationBossSongs()
	if err != nil {
		return err
	}

	params := database.CreateFormationBossSongBulkParams{
		DataHash: 			make([]string, len(songs)),
		SongID: 			make([]int32, len(songs)),
		CelebrateVictory: 	make([]bool, len(songs)),
	}

	for i, s := range songs {
		params.DataHash[i] = generateDataHash(s)
		params.SongID[i] = s.SongID
		params.CelebrateVictory[i] = s.CelebrateVictory
	}

	dbRows, err := qtx.CreateFormationBossSongBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create formation boss songs: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractFormationBossSongs() ([]FormationBossSong, error) {
	songs := []FormationBossSong{}
	var err error

	for i := range l.json.monsterFormations {
		data := &l.json.monsterFormations[i].FormationData

		if data.BossMusic == nil {
			continue
		}

		data.BossMusic.SongID, err = assignFK(data.BossMusic.Song, l.Songs)
		if err != nil {
			return nil, err
		}

		songs = append(songs, *data.BossMusic)
	}

	return dedupeRows(songs, l.Hashes), nil
}


func (l *Lookup) loop4SeedEncounterAreas(qtx *database.Queries, ctx context.Context) error {
	areas, err := l.extractEncounterAreas()
	if err != nil {
		return err
	}

	params := database.CreateEncounterAreaBulkParams{
		DataHash:   	make([]string, len(areas)),
		AreaID: 		make([]int32, len(areas)),
		Specification:	make([]sql.NullString, len(areas)),
	}

	for i, a := range areas {
		params.DataHash[i] = generateDataHash(a)
		params.AreaID[i] = a.AreaID
		params.Specification[i] = h.GetNullString(a.Specification)
	}

	dbRows, err := qtx.CreateEncounterAreaBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create encounter areas: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractEncounterAreas() ([]EncounterArea, error) {
	areas := []EncounterArea{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		for j := range mf.EncounterAreas {
			area := &mf.EncounterAreas[j]

			area.AreaID, err = assignFK(area.LocationArea, l.Areas)
			if err != nil {
				return nil, err
			}

			areas = append(areas, *area)
		}
	}

	return dedupeRows(areas, l.Hashes), nil
}


func (l *Lookup) loop4SeedFormationTriggerCommands(qtx *database.Queries, ctx context.Context) error {
	commands, err := l.extractFormationTriggerCommands()
	if err != nil {
		return err
	}

	params := database.CreateFormationTriggerCommandBulkParams{
		DataHash:   		make([]string, len(commands)),
		TriggerCommandID: 	make([]int32, len(commands)),
		Condition: 			make([]sql.NullString, len(commands)),
		UseAmount: 			make([]sql.NullInt32, len(commands)),
	}

	for i, c := range commands {
		params.DataHash[i] = generateDataHash(c)
		params.TriggerCommandID[i] = c.TriggerCommandID
		params.Condition[i] = h.GetNullString(c.Condition)
		params.UseAmount[i] = h.GetNullInt32(c.UseAmount)
	}

	dbRows, err := qtx.CreateFormationTriggerCommandBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create formation trigger commands: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractFormationTriggerCommands() ([]FormationTriggerCommand, error) {
	commands := []FormationTriggerCommand{}
	var err error

	for i := range l.json.monsterFormations {
		mf := &l.json.monsterFormations[i]

		for j := range mf.TriggerCommands {
			command := & mf.TriggerCommands[j]

			command.TriggerCommandID, err = assignFK(command.AbilityReference.Untyped(), l.TriggerCommands)
			if err != nil {
				return nil, err
			}

			commands = append(commands, *command)
		}
	}

	return dedupeRows(commands, l.Hashes), nil
}