package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getSublocationRelationships(cfg *Config, r *http.Request, sublocation seeding.SubLocation) (LocRel, error) {
	characters, err := getResourcesDB(cfg, r, cfg.e.characters, sublocation, cfg.db.GetSublocationCharacterIDs)
	if err != nil {
		return LocRel{}, err
	}

	aeons, err := getResourcesDB(cfg, r, cfg.e.aeons, sublocation, cfg.db.GetSublocationAeonIDs)
	if err != nil {
		return LocRel{}, err
	}

	shops, err := getResourcesDB(cfg, r, cfg.e.shops, sublocation, cfg.db.GetSublocationShopIDs)
	if err != nil {
		return LocRel{}, err
	}

	treasures, err := getResourcesDB(cfg, r, cfg.e.treasures, sublocation, cfg.db.GetSublocationTreasureIDs)
	if err != nil {
		return LocRel{}, err
	}

	monsters, err := getResourcesDB(cfg, r, cfg.e.monsters, sublocation, cfg.db.GetSublocationMonsterIDs)
	if err != nil {
		return LocRel{}, err
	}

	formations, err := getResourcesDB(cfg, r, cfg.e.monsterFormations, sublocation, cfg.db.GetSublocationMonsterFormationIDs)
	if err != nil {
		return LocRel{}, err
	}

	sidequests, err := getLocBasedSidequests(cfg, r, sublocation, cfg.db.GetSublocationQuestIDs)
	if err != nil {
		return LocRel{}, err
	}

	music, err := getSublocationMusic(cfg, r, sublocation)
	if err != nil {
		return LocRel{}, err
	}

	fmvs, err := getResourcesDB(cfg, r, cfg.e.fmvs, sublocation, cfg.db.GetSublocationFmvIDs)
	if err != nil {
		return LocRel{}, err
	}

	rel := LocRel{
		Characters:     characters,
		Aeons:          aeons,
		Shops:          shops,
		Treasures:      treasures,
		Monsters:       monsters,
		Formations:     formations,
		Sidequests:     sidequests,
		Music:          h.ObjPtrOrNil(music),
		FMVs:           fmvs,
	}

	return rel, nil
}



func getSublocationMusic(cfg *Config, r *http.Request, item seeding.LookupableID) (LocationMusic, error) {
	i := cfg.e.songs

	cueSongs, err := getSublocationCueSongs(cfg, r, i, item)
	if err != nil {
		return LocationMusic{}, err
	}

	bmSongs, err := getSublocationBMSongs(cfg, r, i, item)
	if err != nil {
		return LocationMusic{}, err
	}

	music, err := completeLocationMusic(cfg, r, i, item, cueSongs, bmSongs, LocationMusicQueries{
		FMVSongs:  cfg.db.GetSublocationFMVSongIDs,
		BossMusic: cfg.db.GetSublocationBossSongIDs,
	})
	if err != nil {
		return LocationMusic{}, err
	}

	return music, nil
}

func getSublocationCueSongs(cfg *Config, r *http.Request, i handlerInput[seeding.Song, any, NamedAPIResource, NamedApiResourceList], item seeding.LookupableID) ([]LocationSong, error) {
	dbCueSongs, err := cfg.db.GetSublocationCues(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get cues of %s", item), err)
	}

	cueSongs := []LocationSong{}
	for _, song := range dbCueSongs {
		cueSongs = append(cueSongs, newLocationSong(cfg, i, song.ID, song.ReplacesEncounterMusic))
	}

	return cueSongs, nil
}

func getSublocationBMSongs(cfg *Config, r *http.Request, i handlerInput[seeding.Song, any, NamedAPIResource, NamedApiResourceList], item seeding.LookupableID) ([]LocationSong, error) {
	dbBMSongs, err := cfg.db.GetSublocationBackgroundMusic(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get cues of %s", item), err)
	}

	bmSongs := []LocationSong{}
	for _, song := range dbBMSongs {
		bmSongs = append(bmSongs, newLocationSong(cfg, i, song.ID, song.ReplacesEncounterMusic))
	}

	return bmSongs, nil
}
