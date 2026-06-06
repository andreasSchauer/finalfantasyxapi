package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getSublocationRelationships(cfg *Config, r *http.Request, sublocation seeding.Sublocation) (LocRel, error) {
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.sublocations, sublocation.ID)
	if err != nil {
		return LocRel{}, err
	}

	characters, err := getResourcesDbItem(cfg, r, cfg.e.characters, sublocation, cfg.db.GetSublocationCharacterIDs)
	if err != nil {
		return LocRel{}, err
	}

	aeons, err := getResourcesDbItem(cfg, r, cfg.e.aeons, sublocation, cfg.db.GetSublocationAeonIDs)
	if err != nil {
		return LocRel{}, err
	}

	shops, err := runRelAvailabilityQuery(cfg, r, cfg.e.shops, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeShop))
	if err != nil {
		return LocRel{}, err
	}

	treasures, err := runRelAvailabilityQuery(cfg, r, cfg.e.treasures, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeTreasure))
	if err != nil {
		return LocRel{}, err
	}

	monsters, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsters, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeMonster))
	if err != nil {
		return LocRel{}, err
	}

	formations, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsterFormations, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeMonsterFormation))
	if err != nil {
		return LocRel{}, err
	}

	sidequests, err := getLocBasedSidequests(cfg, r, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeQuest))
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

	fmvs, err := getResourcesDbItem(cfg, r, cfg.e.fmvs, sublocation, cfg.db.GetSublocationFmvIDs)
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
