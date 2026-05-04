package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Song struct {
	ID                   int32
	Name                 string            `json:"name"`
	StreamingName        *string           `json:"streaming_name"`
	InGameName           *string           `json:"in_game_name"`
	OSTName              *string           `json:"ost_name"`
	Translation          *string           `json:"translation"`
	StreamingTrackNumber *int32            `json:"streaming_track_number"`
	MusicSphereID        *int32            `json:"music_sphere_id"`
	OSTDisc              *int32            `json:"ost_disc"`
	OSTTrackNumber       *int32            `json:"ost_track_number"`
	Credits              *SongCredits      `json:"credits"`
	DurationInSeconds    int32             `json:"duration_in_seconds"`
	CanLoop              bool              `json:"can_loop"`
	SpecialUseCase       *string           `json:"special_use_case"`
	BackgroundMusic      []BackgroundMusic `json:"background_music"`
	Cues                 []Cue             `json:"cues"`
}

func (s Song) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Name,
		h.DerefOrNil(s.StreamingName),
		h.DerefOrNil(s.InGameName),
		h.DerefOrNil(s.OSTName),
		h.DerefOrNil(s.Translation),
		h.DerefOrNil(s.StreamingTrackNumber),
		h.DerefOrNil(s.MusicSphereID),
		h.DerefOrNil(s.OSTDisc),
		h.DerefOrNil(s.OSTTrackNumber),
		s.DurationInSeconds,
		s.CanLoop,
		h.DerefOrNil(s.SpecialUseCase),
		h.ObjPtrToID(s.Credits),
	}
}

func (s Song) GetID() int32 {
	return s.ID
}

func (s Song) Error() string {
	return fmt.Sprintf("song %s", s.Name)
}

func (s Song) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   s.ID,
		Name: s.Name,
	}
}

type SongCredits struct {
	ID        int32
	Composer  *string `json:"composer"`
	Arranger  *string `json:"arranger"`
	Performer *string `json:"performer"`
	Lyricist  *string `json:"lyricist"`
}

func (sc SongCredits) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", sc),
		h.DerefOrNil(sc.Composer),
		h.DerefOrNil(sc.Arranger),
		h.DerefOrNil(sc.Performer),
		h.DerefOrNil(sc.Lyricist),
	}
}

func (sc SongCredits) GetID() int32 {
	return sc.ID
}

func (sc SongCredits) Error() string {
	return fmt.Sprintf("song credits with composer: %v, arranger: %v, performer: %v, lyricist: %v", h.PtrToString(sc.Composer), h.PtrToString(sc.Arranger), h.PtrToString(sc.Performer), h.PtrToString(sc.Lyricist))
}

type BackgroundMusic struct {
	ID                     int32
	Condition              *string        `json:"condition"`
	ReplacesEncounterMusic bool           `json:"replaces_encounter_music"`
	LocationAreas          []LocationArea `json:"location_areas"`
}

func (bm BackgroundMusic) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", bm),
		h.DerefOrNil(bm.Condition),
		bm.ReplacesEncounterMusic,
	}
}

func (bm BackgroundMusic) GetID() int32 {
	return bm.ID
}

func (bm *BackgroundMusic) SetID(id int32) {
	bm.ID = id
}

func (bm BackgroundMusic) Error() string {
	return fmt.Sprintf("background music replacing encounter music: %t, condition: %v", bm.ReplacesEncounterMusic, h.PtrToString(bm.Condition))
}

type SongAreaJunction struct {
	StdJunction
	AreaID int32
}

func (j SongAreaJunction) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", j),
		j.ParentID,
		j.ChildID,
		j.AreaID,
	}
}

type Cue struct {
	ID                     int32
	SongID                 int32
	SceneDescription       string         `json:"scene_description"`
	TriggerLocationArea    *LocationArea  `json:"trigger_location_area"`
	IncludedAreas          []LocationArea `json:"included_areas"`
	ReplacesBGMusic        *string        `json:"replaces_bg_music"`
	EndTrigger             *string        `json:"end_trigger"`
	ReplacesEncounterMusic bool           `json:"replaces_encounter_music"`
}

func (c Cue) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", c),
		c.SceneDescription,
		c.SongID,
		h.ObjPtrToID(c.TriggerLocationArea),
		h.DerefOrNil(c.ReplacesBGMusic),
		h.DerefOrNil(c.EndTrigger),
		c.ReplacesEncounterMusic,
	}
}

func (c Cue) GetID() int32 {
	return c.ID
}

func (c *Cue) SetID(id int32) {
	c.ID = id
}

func (c Cue) Error() string {
	return fmt.Sprintf("cue for scene: %s at %v, with song id: %d", c.SceneDescription, h.PtrToString(c.TriggerLocationArea), c.SongID)
}

func (l *Lookup) loop2SeedSongs(qtx *database.Queries, ctx context.Context) error {
	songs, err := l.extractSongs()
	if err != nil {
		return err
	}

	params := database.CreateSongBulkParams{
		DataHash:             make([]string, len(songs)),
		Name:                 make([]string, len(songs)),
		StreamingName:        make([]sql.NullString, len(songs)),
		InGameName:           make([]sql.NullString, len(songs)),
		OstName:              make([]sql.NullString, len(songs)),
		Translation:          make([]sql.NullString, len(songs)),
		StreamingTrackNumber: make([]sql.NullInt32, len(songs)),
		MusicSphereID:        make([]sql.NullInt32, len(songs)),
		OstDisc:              make([]sql.NullInt32, len(songs)),
		OstTrackNumber:       make([]sql.NullInt32, len(songs)),
		DurationInSeconds:    make([]int32, len(songs)),
		CanLoop:              make([]bool, len(songs)),
		SpecialUseCase:       make([]database.NullMusicUseCase, len(songs)),
		CreditsID:            make([]sql.NullInt32, len(songs)),
	}

	for i, s := range songs {
		params.DataHash[i] = generateDataHash(s)
		params.Name[i] = s.Name
		params.StreamingName[i] = h.GetNullString(s.StreamingName)
		params.InGameName[i] = h.GetNullString(s.InGameName)
		params.OstName[i] = h.GetNullString(s.OSTName)
		params.Translation[i] = h.GetNullString(s.Translation)
		params.StreamingTrackNumber[i] = h.GetNullInt32(s.StreamingTrackNumber)
		params.MusicSphereID[i] = h.GetNullInt32(s.MusicSphereID)
		params.OstDisc[i] = h.GetNullInt32(s.OSTDisc)
		params.OstTrackNumber[i] = h.GetNullInt32(s.OSTTrackNumber)
		params.DurationInSeconds[i] = s.DurationInSeconds
		params.CanLoop[i] = s.CanLoop
		params.SpecialUseCase[i] = database.ToNullMusicUseCase(s.SpecialUseCase)
		params.CreditsID[i] = h.ObjPtrToNullInt32ID(s.Credits)
	}

	dbRows, err := qtx.CreateSongBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create songs: %v", err)
	}

	for i, row := range dbRows {
		songs[i].ID = row.ID
		l.json.songs[i].ID = row.ID
		l.Songs[songs[i].Name] = songs[i]
		l.SongsID[row.ID] = songs[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractSongs() ([]Song, error) {
	songs := []Song{}
	var err error

	for i := range l.json.songs {
		song := &l.json.songs[i]

		if song.Credits != nil {
			song.Credits.ID, err = l.getHashID(*song.Credits)
			if err != nil {
				return nil, err
			}
		}
		songs = append(songs, *song)
	}

	return dedupeRows(songs, l.Hashes), nil
}

func (l *Lookup) completeSongs() error {
	for i := range l.json.songs {
		song := &l.json.songs[i]

		err := assignIDs(l, song.BackgroundMusic)
		if err != nil {
			return err
		}

		err = assignIDs(l, song.Cues)
		if err != nil {
			return err
		}

		l.Songs[song.Name] = *song
		l.SongsID[song.ID] = *song
	}

	return nil
}

func (l *Lookup) getSongBackgroundMusic(s Song) ([]BackgroundMusic, error) {
	return s.BackgroundMusic, nil
}

func (l *Lookup) getBackgroundMusicAreas(bm BackgroundMusic) ([]Area, error) {
	return getResources(bm.LocationAreas, l.Areas)
}

func (l *Lookup) seedJuncSongsBackgroundMusic(qtx *database.Queries, ctx context.Context) error {
	const desc string = "songs + background music"
	jParams, err := processThreewayJunctions(l, desc, l.json.songs, l.getSongBackgroundMusic, l.getBackgroundMusicAreas)
	if err != nil {
		return err
	}

	return qtx.CreateSongsBackgroundMusicJunctionBulk(ctx, database.CreateSongsBackgroundMusicJunctionBulkParams{
		DataHash: jParams.DataHashes,
		SongID:   jParams.GrandParentIDs,
		BmID:     jParams.ParentIDs,
		AreaID:   jParams.ChildIDs,
	})
}

func (l *Lookup) loop1SeedSongCredits(qtx *database.Queries, ctx context.Context) error {
	credits := l.extractSongCredits()

	params := database.CreateSongCreditBulkParams{
		DataHash:  make([]string, len(credits)),
		Composer:  make([]database.NullComposer, len(credits)),
		Arranger:  make([]database.NullArranger, len(credits)),
		Performer: make([]sql.NullString, len(credits)),
		Lyricist:  make([]sql.NullString, len(credits)),
	}

	for i, c := range credits {
		params.DataHash[i] = generateDataHash(c)
		params.Composer[i] = database.ToNullComposer(c.Composer)
		params.Arranger[i] = database.ToNullArranger(c.Arranger)
		params.Performer[i] = h.GetNullString(c.Performer)
		params.Lyricist[i] = h.GetNullString(c.Lyricist)
	}

	dbRows, err := qtx.CreateSongCreditBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create song credits: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractSongCredits() []SongCredits {
	credits := []SongCredits{}

	for _, song := range l.json.songs {
		if song.Credits != nil {
			credits = append(credits, *song.Credits)
		}
	}

	return dedupeRows(credits, l.Hashes)
}

func (l *Lookup) loop1SeedBackgroundMusic(qtx *database.Queries, ctx context.Context) error {
	bm := l.extractBackgroundMusic()

	params := database.CreateBackgroundMusicBulkParams{
		DataHash:               make([]string, len(bm)),
		Condition:              make([]sql.NullString, len(bm)),
		ReplacesEncounterMusic: make([]bool, len(bm)),
	}

	for i, music := range bm {
		params.DataHash[i] = generateDataHash(music)
		params.Condition[i] = h.GetNullString(music.Condition)
		params.ReplacesEncounterMusic[i] = music.ReplacesEncounterMusic
	}

	dbRows, err := qtx.CreateBackgroundMusicBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create background music: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractBackgroundMusic() []BackgroundMusic {
	bgMusic := []BackgroundMusic{}

	for _, song := range l.json.songs {
		bgMusic = append(bgMusic, song.BackgroundMusic...)
	}

	return dedupeRows(bgMusic, l.Hashes)
}

func (l *Lookup) loop4SeedCues(qtx *database.Queries, ctx context.Context) error {
	cues, err := l.extractCues()
	if err != nil {
		return err
	}

	params := database.CreateCueBulkParams{
		DataHash:               make([]string, len(cues)),
		SongID:                 make([]int32, len(cues)),
		SceneDescription:       make([]string, len(cues)),
		TriggerAreaID:          make([]sql.NullInt32, len(cues)),
		ReplacesBgMusic:        make([]database.NullBgReplacementType, len(cues)),
		EndTrigger:             make([]sql.NullString, len(cues)),
		ReplacesEncounterMusic: make([]bool, len(cues)),
	}

	for i, c := range cues {
		params.DataHash[i] = generateDataHash(c)
		params.SongID[i] = c.SongID
		params.SceneDescription[i] = c.SceneDescription
		params.TriggerAreaID[i] = h.ObjPtrToNullInt32ID(c.TriggerLocationArea)
		params.ReplacesBgMusic[i] = database.ToNullBgReplacementType(c.ReplacesBGMusic)
		params.EndTrigger[i] = h.GetNullString(c.EndTrigger)
		params.ReplacesEncounterMusic[i] = c.ReplacesEncounterMusic
	}

	dbRows, err := qtx.CreateCueBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create cues: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractCues() ([]Cue, error) {
	cues := []Cue{}
	var err error

	for i := range l.json.songs {
		song := &l.json.songs[i]

		for j := range song.Cues {
			cue := &song.Cues[j]
			cue.SongID = song.ID

			if cue.TriggerLocationArea != nil {
				cue.TriggerLocationArea.ID, err = assignFK(*cue.TriggerLocationArea, l.Areas)
				if err != nil {
					return nil, err
				}
			}

			cues = append(cues, *cue)
		}
	}

	return dedupeRows(cues, l.Hashes), nil
}

func (l *Lookup) getCues() []Cue {
	cues := []Cue{}

	for _, song := range l.json.songs {
		cues = append(cues, song.Cues...)
	}

	return cues
}

func (l *Lookup) getCueIncludedAreas(c Cue) ([]Area, error) {
	return getResources(c.IncludedAreas, l.Areas)
}

func (l *Lookup) seedJuncCuesIncludedAreas(qtx *database.Queries, ctx context.Context) error {
	const desc string = "cues + included areas"
	jParams, err := processJunctions(l, desc, l.getCues(), l.getCueIncludedAreas)
	if err != nil {
		return err
	}

	return qtx.CreateCuesIncludedAreasJunctionBulk(ctx, database.CreateCuesIncludedAreasJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		CueID:          jParams.ParentIDs,
		IncludedAreaID: jParams.ChildIDs,
	})
}
