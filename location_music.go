package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type LocationMusic struct {
	BackgroundMusic []LocationSong     `json:"background_music"`
	Cues            []LocationSong     `json:"cues"`
	FMVs            []NamedAPIResource `json:"fmvs"`
	BossFights      []NamedAPIResource `json:"boss_fights"`
}

func (m LocationMusic) IsZero() bool {
	return 	len(m.BackgroundMusic) 	== 0 &&
			len(m.Cues) 			== 0 &&
			len(m.FMVs) 			== 0 &&
			len(m.BossFights) 		== 0
}

type LocationSong struct {
	Song                   NamedAPIResource `json:"song"`
	ReplacesEncounterMusic bool             `json:"replaces_encounter_music"`
}

func (ls LocationSong) GetAPIResource() APIResource {
	return ls.Song
}

// these two functions can be generalized with ids
func (cfg *Config) getAreaCues(dbCues []database.GetAreaCuesRow) []LocationSong {
	songsCues := []LocationSong{}

	for _, cue := range dbCues {
		song := cfg.newNamedAPIResourceSimple(cfg.e.songs.endpoint, cue.ID, cue.Name)

		locationSong := LocationSong{
			Song:                   song,
			ReplacesEncounterMusic: cue.ReplacesEncounterMusic,
		}

		songsCues = append(songsCues, locationSong)
	}

	return songsCues
}

func (cfg *Config) getAreaBM(dbBm []database.GetAreaBackgroundMusicRow) []LocationSong {
	songsBM := []LocationSong{}

	for _, bm := range dbBm {
		song := cfg.newNamedAPIResourceSimple(cfg.e.songs.endpoint, bm.ID, bm.Name)

		locationSong := LocationSong{
			Song:                   song,
			ReplacesEncounterMusic: bm.ReplacesEncounterMusic,
		}

		songsBM = append(songsBM, locationSong)
	}

	return songsBM
}
