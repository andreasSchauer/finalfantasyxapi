package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
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


func getAreaCues(cfg *Config, i handlerInput[seeding.Song, any, NamedApiResourceList], dbCues []database.GetAreaCuesRow) []LocationSong {
	songs := []LocationSong{}

	for _, cue := range dbCues {		
		song := idToNamedAPIResource(cfg, i, cue.ID)

		locationSong := LocationSong{
			Song:                   song,
			ReplacesEncounterMusic: cue.ReplacesEncounterMusic,
		}

		songs = append(songs, locationSong)
	}

	return songs
}

func getAreaBM(cfg *Config, i handlerInput[seeding.Song, any, NamedApiResourceList], dbBm []database.GetAreaBackgroundMusicRow) []LocationSong {
	songsBM := []LocationSong{}

	for _, bm := range dbBm {
		song := idToNamedAPIResource(cfg, i, bm.ID)

		locationSong := LocationSong{
			Song:                   song,
			ReplacesEncounterMusic: bm.ReplacesEncounterMusic,
		}

		songsBM = append(songsBM, locationSong)
	}

	return songsBM
}
