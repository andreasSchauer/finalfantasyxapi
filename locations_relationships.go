package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getLocationRelationships(cfg *Config, r *http.Request, location seeding.Location) (LocRel, error) {
	characters, err := getResourcesDB(cfg, r, cfg.e.characters, location, cfg.db.GetLocationCharacterIDs)
	if err != nil {
		return LocRel{}, err
	}

	aeons, err := getResourcesDB(cfg, r, cfg.e.aeons, location, cfg.db.GetLocationAeonIDs)
	if err != nil {
		return LocRel{}, err
	}

	shops, err := getResourcesDB(cfg, r, cfg.e.shops, location, cfg.db.GetLocationShopIDs)
	if err != nil {
		return LocRel{}, err
	}

	treasures, err := getResourcesDB(cfg, r, cfg.e.treasures, location, cfg.db.GetLocationTreasureIDs)
	if err != nil {
		return LocRel{}, err
	}

	monsters, err := getResourcesDB(cfg, r, cfg.e.monsters, location, cfg.db.GetLocationMonsterIDs)
	if err != nil {
		return LocRel{}, err
	}

	formations, err := getResourcesDB(cfg, r, cfg.e.monsterFormations, location, cfg.db.GetLocationMonsterFormationIDs)
	if err != nil {
		return LocRel{}, err
	}

	sidequests, err := getLocBasedSidequests(cfg, r, location, cfg.db.GetLocationQuestIDs)
	if err != nil {
		return LocRel{}, err
	}

	music, err := getMusicLocBased(cfg, r, location, LocBasedMusicQueries{
		CueSongs:  cfg.db.GetLocationCueSongIDs,
		BmSongs:   cfg.db.GetLocationBackgroundMusicSongIDs,
		FMVSongs:  cfg.db.GetLocationFMVSongIDs,
		BossMusic: cfg.db.GetLocationBossSongIDs,
	})
	if err != nil {
		return LocRel{}, err
	}

	fmvs, err := getResourcesDB(cfg, r, cfg.e.fmvs, location, cfg.db.GetLocationFmvIDs)
	if err != nil {
		return LocRel{}, err
	}

	rel := LocRel{
		Characters: characters,
		Aeons:      aeons,
		Shops:      shops,
		Treasures:  treasures,
		Monsters:   monsters,
		Formations: formations,
		Sidequests: sidequests,
		Music:      music,
		FMVs:       fmvs,
	}

	return rel, nil
}
