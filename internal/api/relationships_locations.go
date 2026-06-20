package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getLocationRelationships(cfg *Config, r *http.Request, location seeding.Location) (LocRel, error) {
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.locations, location.ID)
	if err != nil {
		return LocRel{}, err
	}

	characters, err := getResourcesDbItem(cfg, r, cfg.e.characters, location, cfg.db.GetLocationCharacterIDs)
	if err != nil {
		return LocRel{}, err
	}

	aeons, err := getResourcesDbItem(cfg, r, cfg.e.aeons, location, cfg.db.GetLocationAeonIDs)
	if err != nil {
		return LocRel{}, err
	}

	shops, err := runRelAvailabilityQuery(cfg, r, cfg.e.shops, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeShop))
	if err != nil {
		return LocRel{}, err
	}

	treasures, err := runRelAvailabilityQuery(cfg, r, cfg.e.treasures, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeTreasure))
	if err != nil {
		return LocRel{}, err
	}

	monsters, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsters, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeMonster))
	if err != nil {
		return LocRel{}, err
	}

	formations, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsterFormations, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeMonsterFormation))
	if err != nil {
		return LocRel{}, err
	}

	quests, err := runRelAvailabilityQuery(cfg, r, cfg.e.quests, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeQuest))
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

	fmvs, err := getResourcesDbItem(cfg, r, cfg.e.fmvs, location, cfg.db.GetLocationFmvIDs)
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
		Quests:     quests,
		Music:      music,
		FMVs:       fmvs,
	}

	return rel, nil
}
