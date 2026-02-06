package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getSublocationRelationships(cfg *Config, r *http.Request, sublocation seeding.Sublocation) (LocRel, error) {
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

	music, err := getMusicLocBased(cfg, r, sublocation, LocBasedMusicQueries{
		CueSongs:  cfg.db.GetSublocationCueSongIDs,
		BmSongs:   cfg.db.GetSublocationBackgroundMusicSongIDs,
		FMVSongs:  cfg.db.GetSublocationFMVSongIDs,
		BossMusic: cfg.db.GetSublocationBossSongIDs,
	})
	if err != nil {
		return LocRel{}, err
	}

	fmvs, err := getResourcesDB(cfg, r, cfg.e.fmvs, sublocation, cfg.db.GetSublocationFmvIDs)
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
