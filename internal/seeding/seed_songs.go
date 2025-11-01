package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Song struct {
	ID						int32
	Name 					string				`json:"name"`
	StreamingName 			*string				`json:"streaming_name"`
	InGameName 				*string				`json:"in_game_name"`
	OSTName 				*string				`json:"ost_name"`
	Translation 			*string				`json:"translation"`
	StreamingTrackNumber 	*int32				`json:"streaming_track_number"`
	MusicSphereID 			*int32				`json:"music_sphere_id"`
	OSTDisc 				*int32				`json:"ost_disc"`
	OSTTrackNumber	 		*int32				`json:"ost_track_number"`
	Credits					*SongCredits		`json:"credits"`
	DurationInSeconds 		int32				`json:"duration_in_seconds"`
	CanLoop 				bool				`json:"can_loop"`
	SpecialUseCase 			*string				`json:"special_use_case"`
	BackgroundMusic			[]BackgroundMusic	`json:"background_music"`
	Cues					[]Cue				`json:"cues"`
}

func (s Song) ToHashFields() []any {
	return []any{
		s.Name,
		derefOrNil(s.StreamingName),
		derefOrNil(s.InGameName),
		derefOrNil(s.OSTName),
		derefOrNil(s.Translation),
		derefOrNil(s.StreamingTrackNumber),
		derefOrNil(s.MusicSphereID),
		derefOrNil(s.OSTDisc),
		derefOrNil(s.OSTTrackNumber),
		s.DurationInSeconds,
		s.CanLoop,
		derefOrNil(s.SpecialUseCase),
		ObjPtrToHashID(s.Credits),
	}
}

func (s Song) GetID() int32 {
	return s.ID
}



type SongCredits struct {
	ID			int32
	Composer	*string		`json:"composer"`
	Arranger	*string		`json:"arranger"`
	Performer	*string		`json:"performer"`
	Lyricist	*string		`json:"lyricist"`
}


func (sc SongCredits) ToHashFields() []any {
	return []any{
		derefOrNil(sc.Composer),
		derefOrNil(sc.Arranger),
		derefOrNil(sc.Performer),
		derefOrNil(sc.Lyricist),
	}
}


func (sc SongCredits) GetID() int32 {
	return sc.ID
}


type BackgroundMusic struct {
	ID						int32
	Condition				*string			`json:"condition"`
	ReplacesEncounterMusic	bool			`json:"replaces_encounter_music"`
	LocationAreas			[]LocationArea	`json:"location_areas"`
}

func (bm BackgroundMusic) ToHashFields() []any {
	return []any{
		derefOrNil(bm.Condition),
		bm.ReplacesEncounterMusic,
	}
}

func (bm BackgroundMusic) GetID() int32 {
	return bm.ID
}


type SongAreaJunction struct {
	Junction
	AreaID		int32
}

func (j SongAreaJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
		j.AreaID,
	}
}


type Cue struct {
	ID						int32
	SceneDescription		string			`json:"scene_description"`
	TriggerLocationArea		*LocationArea	`json:"trigger_location_area"`
	IncludedAreas			[]LocationArea	`json:"included_areas"`
	ReplacesBGMusic			*string			`json:"replaces_bg_music"`
	EndTrigger				*string			`json:"end_trigger"`
	ReplacesEncounterMusic	bool			`json:"replaces_encounter_music"`
}

func (c Cue) ToHashFields() []any {
	return []any{
		c.SceneDescription,
		ObjPtrToHashID(c.TriggerLocationArea),
		derefOrNil(c.ReplacesBGMusic),
		derefOrNil(c.EndTrigger),
		c.ReplacesEncounterMusic,
	}
}

func (c Cue) GetID() int32 {
	return c.ID
}



func (l *lookup) seedSongs(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/songs.json"

	var songs []Song
	err := loadJSONFile(string(srcPath), &songs)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, song := range songs {
			dbSong, err := qtx.CreateSong(context.Background(), database.CreateSongParams{
				DataHash: 				generateDataHash(song),
				Name:     				song.Name,
				StreamingName: 			getNullString(song.StreamingName),
				InGameName: 			getNullString(song.InGameName),
				OstName: 				getNullString(song.OSTName),
				Translation: 			getNullString(song.Translation),
				StreamingTrackNumber: 	getNullInt32(song.StreamingTrackNumber),
				MusicSphereID: 			getNullInt32(song.MusicSphereID),
				OstDisc: 				getNullInt32(song.OSTDisc),
				OstTrackNumber: 		getNullInt32(song.OSTTrackNumber),
				DurationInSeconds: 		song.DurationInSeconds,
				CanLoop: 				song.CanLoop,
				SpecialUseCase: 		nullMusicUseCase(song.SpecialUseCase),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Song: %s: %v", song.Name, err)
			}

			song.ID = dbSong.ID
			l.songs[song.Name] = song
		}
		return nil
	})
}


func (l *lookup) seedSongsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/songs.json"

	var songs []Song
	err := loadJSONFile(string(srcPath), &songs)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonSong := range songs {
			song, err := l.getSong(jsonSong.Name)
			if err != nil {
				return err
			}

			song.Credits, err = seedObjPtrAssignFK(qtx, song.Credits, l.seedCredits)
			if err != nil {
				return fmt.Errorf("song %s: %v", song.Name, err)
			}

			err = qtx.UpdateSong(context.Background(), database.UpdateSongParams{
				DataHash: generateDataHash(song),
				CreditsID: ObjPtrToNullInt32ID(song.Credits),
				ID: song.ID,
			})
			if err != nil {
				return fmt.Errorf("couldn't update song %s: %v", song.Name, err)
			}

			err = l.seedBackgroundMusicEntries(qtx, song)
			if err != nil {
				return fmt.Errorf("song %s: %v", song.Name, err)
			}

			err = l.seedCues(qtx, song)
			if err != nil {
				return fmt.Errorf("song %s: %v", song.Name, err)
			}
		}	

		return nil
	})
}



func (l *lookup) seedCredits(qtx *database.Queries, credits SongCredits) (SongCredits, error) {
	dbSongCredits, err := qtx.CreateSongCredit(context.Background(), database.CreateSongCreditParams{
		DataHash: 	generateDataHash(credits),
		Composer: 	getNullString(credits.Composer),
		Arranger: 	getNullString(credits.Arranger),
		Performer: 	getNullString(credits.Performer),
		Lyricist: 	getNullString(credits.Lyricist),
	})
	if err != nil {
		return SongCredits{}, fmt.Errorf("couldn't create credits: %v", err)
	}

	credits.ID = dbSongCredits.ID

	return credits, nil
}


func (l *lookup) seedBackgroundMusicEntries (qtx *database.Queries, song Song) error {
	for _, bm := range song.BackgroundMusic {
		junction, err := createJunctionSeed(qtx, song, bm, l.seedBackgroundMusic)
		if err != nil {
			return err
		}

		for _, locationArea := range bm.LocationAreas {
			var err error
			
			saJunction := SongAreaJunction{}
			saJunction.Junction = junction
			saJunction.AreaID, err = assignFK(locationArea, l.getArea)
			if err != nil {
				return err
			}
			
			err = qtx.CreateSongsBackgroundMusicJunction(context.Background(), database.CreateSongsBackgroundMusicJunctionParams{
				DataHash: 	generateDataHash(saJunction),
				SongID: 	saJunction.ParentID,
				BmID: 		saJunction.ChildID,
				AreaID: 	saJunction.AreaID,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}


func (l *lookup) seedCues (qtx *database.Queries, song Song) error {
	for _, cue := range song.Cues {
		junction, err := createJunctionSeed(qtx, song, cue, l.seedCue)
		if err != nil {
			return err
		}

		for _, locationArea := range cue.IncludedAreas {
			var err error
			
			saJunction := SongAreaJunction{}
			saJunction.Junction = junction
			saJunction.AreaID, err = assignFK(locationArea, l.getArea)
			if err != nil {
				return err
			}

			err = qtx.CreateSongsCuesJunction(context.Background(), database.CreateSongsCuesJunctionParams{
				DataHash: 	generateDataHash(saJunction),
				SongID: 	saJunction.ParentID,
				CueID: 		saJunction.ChildID,
				AreaID: 	saJunction.AreaID,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}


func (l *lookup) seedBackgroundMusic(qtx *database.Queries, bm BackgroundMusic) (BackgroundMusic, error) {
	dbBM, err := qtx.CreateBackgroundMusic(context.Background(), database.CreateBackgroundMusicParams{
		DataHash: 				generateDataHash(bm),
		Condition: 				getNullString(bm.Condition),
		ReplacesEncounterMusic: bm.ReplacesEncounterMusic,
	})
	if err != nil {
		return BackgroundMusic{}, fmt.Errorf("couldn't create background music: %v", err)
	}

	bm.ID = dbBM.ID

	return bm, nil
}


func (l *lookup) seedCue(qtx *database.Queries, cue Cue) (Cue, error) {
	if cue.TriggerLocationArea != nil {
		var err error
		
		cue.TriggerLocationArea.ID, err = assignFK(*cue.TriggerLocationArea, l.getArea)
		if err != nil {
			return Cue{}, err
		}
	}

	dbCue, err := qtx.CreateCue(context.Background(), database.CreateCueParams{
		DataHash: 				generateDataHash(cue),
		SceneDescription: 		cue.SceneDescription,
		AreaID: 				ObjPtrToNullInt32ID(cue.TriggerLocationArea),
		ReplacesBgMusic: 		nullBgReplacementType(cue.ReplacesBGMusic),
		EndTrigger: 			getNullString(cue.EndTrigger),
		ReplacesEncounterMusic: cue.ReplacesEncounterMusic,
	})
	if err != nil {
		return Cue{}, fmt.Errorf("couldn't create cue: %v", err)
	}

	cue.ID = dbCue.ID

	return cue, nil
}