package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop2SeedSongs(qtx *database.Queries, ctx context.Context) error {
	songs, err := extractSongs()
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
		params.DataHash[i] = seedhelpers.GenerateDataHash(s)
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
		seedhelpers.RawJSON.Songs[i].ID = row.ID
		lookup.L.Songs[seedtypes.Key(songs[i])] = songs[i]
		lookup.L.SongsID[row.ID] = songs[i]
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractSongs() ([]seedtypes.Song, error) {
	songs := []seedtypes.Song{}
	var err error

	for i := range seedhelpers.RawJSON.Songs {
		song := &seedhelpers.RawJSON.Songs[i]

		if song.Credits != nil {
			song.Credits.ID, err = seedhelpers.GetHashID(*song.Credits)
			if err != nil {
				return nil, err
			}
		}
		songs = append(songs, *song)
	}

	return dedupeRows(songs, seedhelpers.Hashes), nil
}

func completeSongs() error {
	for i := range seedhelpers.RawJSON.Songs {
		song := &seedhelpers.RawJSON.Songs[i]

		err := assignIDs(song.BackgroundMusic)
		if err != nil {
			return err
		}

		err = assignIDs(song.Cues)
		if err != nil {
			return err
		}

		lookup.L.Songs[seedtypes.Key(song)] = *song
		lookup.L.SongsID[song.ID] = *song
	}

	return nil
}

func Loop1SeedSongCredits(qtx *database.Queries, ctx context.Context) error {
	credits := extractSongCredits()

	params := database.CreateSongCreditBulkParams{
		DataHash:  make([]string, len(credits)),
		Composer:  make([]database.NullComposer, len(credits)),
		Arranger:  make([]database.NullArranger, len(credits)),
		Performer: make([]sql.NullString, len(credits)),
		Lyricist:  make([]sql.NullString, len(credits)),
	}

	for i, c := range credits {
		params.DataHash[i] = seedhelpers.GenerateDataHash(c)
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
		seedhelpers.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractSongCredits() []seedtypes.SongCredits {
	credits := []seedtypes.SongCredits{}

	for _, song := range seedhelpers.RawJSON.Songs {
		if song.Credits != nil {
			credits = append(credits, *song.Credits)
		}
	}

	return dedupeRows(credits, seedhelpers.Hashes)
}
