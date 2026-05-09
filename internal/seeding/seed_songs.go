package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

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
		l.Songs[Key(songs[i])] = songs[i]
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
			song.Credits.ID, err = l.GetHashID(*song.Credits)
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

		l.Songs[Key(song)] = *song
		l.SongsID[song.ID] = *song
	}

	return nil
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
