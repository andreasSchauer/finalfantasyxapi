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
		ID: 	s.ID,
		Name: 	s.Name,
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
	return fmt.Sprintf("song credits with composer: %v, arranger: %v, performer: %v, lyricist: %v", h.DerefOrNil(sc.Composer), h.DerefOrNil(sc.Arranger), h.DerefOrNil(sc.Performer), h.DerefOrNil(sc.Lyricist))
}

type BackgroundMusic struct {
	ID                     int32
	Condition              *string        `json:"condition"`
	ReplacesEncounterMusic bool           `json:"replaces_encounter_music"`
	LocationAreas          []LocationArea `json:"location_areas"`
}

func (bm BackgroundMusic) ToHashFields() []any {
	return []any{
		h.DerefOrNil(bm.Condition),
		bm.ReplacesEncounterMusic,
	}
}

func (bm BackgroundMusic) GetID() int32 {
	return bm.ID
}

func (bm BackgroundMusic) Error() string {
	return fmt.Sprintf("background music replacing encounter music: %t, condition: %v", bm.ReplacesEncounterMusic, h.DerefOrNil(bm.Condition))
}

type SongAreaJunction struct {
	Junction
	AreaID int32
}

func (j SongAreaJunction) ToHashFields() []any {
	return []any{
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

func (c Cue) Error() string {
	return fmt.Sprintf("cue for scene: %s at %v, with song id: %d", c.SceneDescription, h.DerefOrNil(c.TriggerLocationArea), c.SongID)
}

func (l *Lookup) seedSongs(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/songs.json"

	var songs []Song
	err := loadJSONFile(string(srcPath), &songs)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, song := range songs {
			dbSong, err := qtx.CreateSong(context.Background(), database.CreateSongParams{
				DataHash:             generateDataHash(song),
				Name:                 song.Name,
				StreamingName:        h.GetNullString(song.StreamingName),
				InGameName:           h.GetNullString(song.InGameName),
				OstName:              h.GetNullString(song.OSTName),
				Translation:          h.GetNullString(song.Translation),
				StreamingTrackNumber: h.GetNullInt32(song.StreamingTrackNumber),
				MusicSphereID:        h.GetNullInt32(song.MusicSphereID),
				OstDisc:              h.GetNullInt32(song.OSTDisc),
				OstTrackNumber:       h.GetNullInt32(song.OSTTrackNumber),
				DurationInSeconds:    song.DurationInSeconds,
				CanLoop:              song.CanLoop,
				SpecialUseCase:       h.NullMusicUseCase(song.SpecialUseCase),
			})
			if err != nil {
				return h.NewErr(song.Error(), err, "couldn't create song")
			}

			song.ID = dbSong.ID
			l.Songs[song.Name] = song
			l.SongsID[song.ID] = song
		}
		return nil
	})
}

func (l *Lookup) seedSongsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/songs.json"

	var songs []Song
	err := loadJSONFile(string(srcPath), &songs)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonSong := range songs {
			song, err := GetResource(jsonSong.Name, l.Songs)
			if err != nil {
				return err
			}

			song.Credits, err = seedObjPtrAssignFK(qtx, song.Credits, l.seedCredits)
			if err != nil {
				return h.NewErr(song.Error(), err)
			}

			err = qtx.UpdateSong(context.Background(), database.UpdateSongParams{
				DataHash:  generateDataHash(song),
				CreditsID: h.ObjPtrToNullInt32ID(song.Credits),
				ID:        song.ID,
			})
			if err != nil {
				return h.NewErr(song.Error(), err, "couldn't update song")
			}

			err = l.seedBackgroundMusicEntries(qtx, song)
			if err != nil {
				return h.NewErr(song.Error(), err)
			}

			err = l.seedCues(qtx, song)
			if err != nil {
				return h.NewErr(song.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedCredits(qtx *database.Queries, credits SongCredits) (SongCredits, error) {
	dbSongCredits, err := qtx.CreateSongCredit(context.Background(), database.CreateSongCreditParams{
		DataHash:  generateDataHash(credits),
		Composer:  h.NullComposer(credits.Composer),
		Arranger:  h.NullArranger(credits.Arranger),
		Performer: h.GetNullString(credits.Performer),
		Lyricist:  h.GetNullString(credits.Lyricist),
	})
	if err != nil {
		return SongCredits{}, h.NewErr(credits.Error(), err, "couldn't create song credits")
	}

	credits.ID = dbSongCredits.ID

	return credits, nil
}

func (l *Lookup) seedBackgroundMusicEntries(qtx *database.Queries, song Song) error {
	for _, bm := range song.BackgroundMusic {
		junction, err := createJunctionSeed(qtx, song, bm, l.seedBackgroundMusic)
		if err != nil {
			return err
		}

		for _, locationArea := range bm.LocationAreas {
			var err error

			saJunction := SongAreaJunction{}
			saJunction.Junction = junction
			saJunction.AreaID, err = assignFK(locationArea, l.Areas)
			if err != nil {
				return h.NewErr(bm.Error(), err)
			}

			err = qtx.CreateSongsBackgroundMusicJunction(context.Background(), database.CreateSongsBackgroundMusicJunctionParams{
				DataHash: generateDataHash(saJunction),
				SongID:   saJunction.ParentID,
				BmID:     saJunction.ChildID,
				AreaID:   saJunction.AreaID,
			})
			if err != nil {
				return h.NewErr(bm.Error(), err, "couldn't junction background music entry")
			}
		}
	}

	return nil
}

func (l *Lookup) seedCues(qtx *database.Queries, song Song) error {
	for _, cue := range song.Cues {
		cue.SongID = song.ID
		cue, err := l.seedCue(qtx, cue)
		if err != nil {
			return h.NewErr(cue.Error(), err)
		}

		for _, locationArea := range cue.IncludedAreas {
			var err error

			junction, err := createJunction(cue, locationArea, l.Areas)
			if err != nil {
				return err
			}

			err = qtx.CreateSongsCuesJunction(context.Background(), database.CreateSongsCuesJunctionParams{
				DataHash:       generateDataHash(junction),
				CueID:          junction.ParentID,
				IncludedAreaID: junction.ChildID,
			})
			if err != nil {
				return h.NewErr(cue.Error(), err, "couldn't junction cue")
			}
		}
	}

	return nil
}

func (l *Lookup) seedBackgroundMusic(qtx *database.Queries, bm BackgroundMusic) (BackgroundMusic, error) {
	dbBM, err := qtx.CreateBackgroundMusic(context.Background(), database.CreateBackgroundMusicParams{
		DataHash:               generateDataHash(bm),
		Condition:              h.GetNullString(bm.Condition),
		ReplacesEncounterMusic: bm.ReplacesEncounterMusic,
	})
	if err != nil {
		return BackgroundMusic{}, h.NewErr(bm.Error(), err, "couldn't create background music entry")
	}

	bm.ID = dbBM.ID
	l.BackgroundMusicID[bm.ID] = bm

	return bm, nil
}

func (l *Lookup) seedCue(qtx *database.Queries, cue Cue) (Cue, error) {
	if cue.TriggerLocationArea != nil {
		var err error

		cue.TriggerLocationArea.ID, err = assignFK(*cue.TriggerLocationArea, l.Areas)
		if err != nil {
			return Cue{}, h.NewErr(cue.Error(), err)
		}
	}

	dbCue, err := qtx.CreateCue(context.Background(), database.CreateCueParams{
		DataHash:               generateDataHash(cue),
		SongID:                 cue.SongID,
		SceneDescription:       cue.SceneDescription,
		TriggerAreaID:          h.ObjPtrToNullInt32ID(cue.TriggerLocationArea),
		ReplacesBgMusic:        h.NullBgReplacementType(cue.ReplacesBGMusic),
		EndTrigger:             h.GetNullString(cue.EndTrigger),
		ReplacesEncounterMusic: cue.ReplacesEncounterMusic,
	})
	if err != nil {
		return Cue{}, h.NewErr(cue.Error(), err, "couldn't create cue")
	}

	cue.ID = dbCue.ID
	l.CuesID[cue.ID] = cue

	return cue, nil
}
