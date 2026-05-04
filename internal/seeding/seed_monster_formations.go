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

type FormationData struct {
	ID             int32              `json:"-"`
	Category       string             `json:"category"`
	Availability   string             `json:"availability"`
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

func (ea EncounterArea) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", ea),
		ea.AreaID,
		h.DerefOrNil(ea.Specification),
	}
}

func (ea EncounterArea) ToKeyFields() []any {
	return []any{
		Key(ea.LocationArea),
		ea.Specification,
	}
}

func (ea EncounterArea) GetID() int32 {
	return ea.ID
}

func (ea *EncounterArea) SetID(id int32) {
	ea.ID = id
}

func (ea EncounterArea) Error() string {
	return fmt.Sprintf("encounter location with %s, specification: %s", ea.LocationArea, h.PtrToString(ea.Specification))
}

func (ea EncounterArea) GetLocationArea() LocationArea {
	return ea.LocationArea
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

func (tc *FormationTriggerCommand) SetID(id int32) {
	tc.ID = id
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

func (l *Lookup) loop5SeedMonsterFormations(qtx *database.Queries, ctx context.Context) error {
	formations, err := l.extractMonsterFormations()
	if err != nil {
		return err
	}

	params := database.CreateMonsterFormationBulkParams{
		DataHash:           make([]string, len(formations)),
		Version:            make([]sql.NullInt32, len(formations)),
		MonsterSelectionID: make([]int32, len(formations)),
		FormationDataID:    make([]int32, len(formations)),
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
		key := Key(formations[i])
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

func (l *Lookup) completeMonsterFormations() error {
	for i := range l.json.monsterFormations {
		formation := &l.json.monsterFormations[i]

		err := assignIDs(l, formation.Monsters)
		if err != nil {
			return err
		}

		err = assignIDs(l, formation.EncounterAreas)
		if err != nil {
			return err
		}

		err = assignIDs(l, formation.TriggerCommands)
		if err != nil {
			return err
		}

		l.MonsterFormations[Key(*formation)] = *formation
		l.MonsterFormationsID[formation.ID] = *formation
	}

	return nil
}

func (l *Lookup) getMonsterFormationEncounterAreas(mf MonsterFormation) ([]EncounterArea, error) {
	return mf.EncounterAreas, nil
}

func (l *Lookup) getMonsterFormationTriggerCommands(mf MonsterFormation) ([]FormationTriggerCommand, error) {
	return mf.TriggerCommands, nil
}

func (l *Lookup) seedJuncMonsterFormationsEncounterAreas(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster formations + encounter areas"
	jParams, err := processJunctions(l, desc, l.json.monsterFormations, l.getMonsterFormationEncounterAreas)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterFormationsEncounterAreasJunctionBulk(ctx, database.CreateMonsterFormationsEncounterAreasJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterFormationID: jParams.ParentIDs,
		EncounterAreaID:    jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncMonsterFormationsTriggerCommands(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster formations + trigger commands"
	jParams, err := processJunctions(l, desc, l.json.monsterFormations, l.getMonsterFormationTriggerCommands)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterFormationsTriggerCommandsJunctionBulk(ctx, database.CreateMonsterFormationsTriggerCommandsJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterFormationID: jParams.ParentIDs,
		TriggerCommandID:   jParams.ChildIDs,
	})
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

			key := Key(*ma)
			ma.MonsterID, err = assignFK(key, l.Monsters)
			if err != nil {
				return nil, err
			}
		}

		selections = append(selections, mf.MonsterSelection)
	}

	return dedupeRows(selections, l.Hashes), nil
}

func (l *Lookup) getMonsterSelections() []MonsterSelection {
	selections := []MonsterSelection{}

	for _, formation := range l.json.monsterFormations {
		selections = append(selections, formation.MonsterSelection)
	}

	return selections
}

func (l *Lookup) getMonsterSelectionMonsterAmounts(ms MonsterSelection) ([]MonsterAmount, error) {
	return ms.Monsters, nil
}

func (l *Lookup) seedJuncMonsterSelectionMonsterAmounts(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster selection + monsters"
	jParams, err := processJunctions(l, desc, l.getMonsterSelections(), l.getMonsterSelectionMonsterAmounts)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterSelectionsMonstersJunctionBulk(ctx, database.CreateMonsterSelectionsMonstersJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterSelectionID: jParams.ParentIDs,
		MonsterAmountID:    jParams.ChildIDs,
	})
}

func (l *Lookup) loop4SeedFormationData(qtx *database.Queries, ctx context.Context) error {
	data, err := l.extractFormationData()
	if err != nil {
		return err
	}

	params := database.CreateFormationDataBulkParams{
		DataHash:       make([]string, len(data)),
		Category:       make([]database.MonsterFormationCategory, len(data)),
		Availability:   make([]database.AvailabilityType, len(data)),
		IsForcedAmbush: make([]bool, len(data)),
		CanEscape:      make([]bool, len(data)),
		BossSongID:     make([]sql.NullInt32, len(data)),
		Notes:          make([]sql.NullString, len(data)),
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


func (l *Lookup) loop3SeedFormationBossSongs(qtx *database.Queries, ctx context.Context) error {
	songs, err := l.extractFormationBossSongs()
	if err != nil {
		return err
	}

	params := database.CreateFormationBossSongBulkParams{
		DataHash:         make([]string, len(songs)),
		SongID:           make([]int32, len(songs)),
		CelebrateVictory: make([]bool, len(songs)),
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
		DataHash:      make([]string, len(areas)),
		AreaID:        make([]int32, len(areas)),
		Specification: make([]sql.NullString, len(areas)),
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
		DataHash:         make([]string, len(commands)),
		TriggerCommandID: make([]int32, len(commands)),
		Condition:        make([]sql.NullString, len(commands)),
		UseAmount:        make([]sql.NullInt32, len(commands)),
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
			command := &mf.TriggerCommands[j]

			command.TriggerCommandID, err = assignFK(command.AbilityReference.Untyped(), l.TriggerCommands)
			if err != nil {
				return nil, err
			}

			commands = append(commands, *command)
		}
	}

	return dedupeRows(commands, l.Hashes), nil
}

func (l *Lookup) getFormationTriggerCommands() []FormationTriggerCommand {
	commands := []FormationTriggerCommand{}

	for _, formation := range l.json.monsterFormations {
		commands = append(commands, formation.TriggerCommands...)
	}

	return commands
}

func (l *Lookup) getFormationTriggerCommandUsers(tc FormationTriggerCommand) ([]CharacterClass, error) {
	return getResources(tc.Users, l.CharClasses)
}

func (l *Lookup) seedJuncFormationTriggerCommandsUsers(qtx *database.Queries, ctx context.Context) error {
	const desc string = "formation trigger commands + users"
	jParams, err := processJunctions(l, desc, l.getFormationTriggerCommands(), l.getFormationTriggerCommandUsers)
	if err != nil {
		return err
	}

	return qtx.CreateFormationTriggerCommandsUsersJunctionBulk(ctx, database.CreateFormationTriggerCommandsUsersJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		TriggerCommandID: jParams.ParentIDs,
		CharacterClassID: jParams.ChildIDs,
	})
}
