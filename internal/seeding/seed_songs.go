package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Song struct {
	//id 			int32
	//dataHash		string
	Name 					string		`json:"name"`
	StreamingName 			*string		`json:"streaming_name"`
	InGameName 				*string		`json:"in_game_name"`
	OSTName 				*string		`json:"ost_name"`
	Translation 			*string		`json:"translation"`
	StreamingTrackNumber 	*int32		`json:"streaming_track_number"`
	MusicSphereID 			*int32		`json:"music_sphere_id"`
	OSTDisc 				*int32		`json:"ost_disc"`
	OSTTrackNumber	 		*int32		`json:"ost_track_number"`
	Credits					SongCredits	`json:"credits"`
	DurationInSeconds 		int32		`json:"duration_in_seconds"`
	CanLoop 				bool		`json:"can_loop"`
	SpecialUseCase 			*string		`json:"special_use_case"`
	CreditsID 				*int32
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
		derefOrNil(s.CreditsID),
	}
}


func (s Song) ToKeyFields() []any {
	return []any{
		s.Name,
	}
}


type SongLookup struct {
	Song
	ID 		int32
}


type SongCredits struct {
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



func (l *lookup) seedSongs(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/songs.json"

	var songs []Song
	err := loadJSONFile(string(srcPath), &songs)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, song := range songs {
			credits := song.Credits
			dbSongCredits, err := qtx.CreateSongCredit(context.Background(), database.CreateSongCreditParams{
				DataHash: 	generateDataHash(credits),
				Composer: 	getNullString(credits.Composer),
				Arranger: 	getNullString(credits.Arranger),
				Performer: 	getNullString(credits.Performer),
				Lyricist: 	getNullString(credits.Lyricist),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Credits for Song: %s: %v", song.Name, err)
			}

			song.CreditsID = &dbSongCredits.ID

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
				CreditsID: 				getNullInt32(song.CreditsID),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Song: %s: %v", song.Name, err)
			}

			key := createLookupKey(song)
			l.songs[key] = SongLookup{
				Song: 	song,
				ID: 	dbSong.ID,
			}
		}
		return nil
	})
}
