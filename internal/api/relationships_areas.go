package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AreaConnection struct {
	Area           AreaAPIResource `json:"area"`
	ConnectionType string          `json:"connection_type"`
	IsStoryBased   bool            `json:"is_story_based"`
	Notes          *string         `json:"notes,omitempty"`
}

func (ac AreaConnection) GetAPIResource() APIResource {
	return ac.Area
}

func getAreaRelationships(cfg *Config, r *http.Request, area seeding.Area) (LocRel, error) {
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.areas, area.ID)
	if err != nil {
		return LocRel{}, err
	}

	characters, err := getResourcesDbItem(cfg, r, cfg.e.characters, area, ToIntManyNull(cfg.db.GetAreaCharacterIDs))
	if err != nil {
		return LocRel{}, err
	}

	aeons, err := getResourcesDbItem(cfg, r, cfg.e.aeons, area, ToIntManyNull(cfg.db.GetAreaAeonIDs))
	if err != nil {
		return LocRel{}, err
	}

	shops, err := runRelAvailabilityQuery(cfg, r, cfg.e.shops, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeShop))
	if err != nil {
		return LocRel{}, err
	}

	treasures, err := runRelAvailabilityQuery(cfg, r, cfg.e.treasures, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeTreasure))
	if err != nil {
		return LocRel{}, err
	}

	monsters, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsters, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeMonster))
	if err != nil {
		return LocRel{}, err
	}

	formations, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsterFormations, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeMonsterFormation))
	if err != nil {
		return LocRel{}, err
	}

	sidequests, err := getLocBasedSidequests(cfg, r, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeQuest))
	if err != nil {
		return LocRel{}, err
	}

	music, err := getMusicLocBased(cfg, r, area, LocBasedMusicQueries{
		CueSongs:  ToIntManyNull(cfg.db.GetAreaCueSongIDs),
		BmSongs:   cfg.db.GetAreaBackgroundMusicSongIDs,
		FMVSongs:  cfg.db.GetAreaFMVSongIDs,
		BossMusic: cfg.db.GetAreaBossSongIDs,
	})
	if err != nil {
		return LocRel{}, err
	}

	fmvs, err := getResourcesDbItem(cfg, r, cfg.e.fmvs, area, cfg.db.GetAreaFmvIDs)
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

func getAreaConnectedAreas(cfg *Config, area seeding.Area) ([]AreaConnection, error) {
	i := cfg.e.areas
	connectedAreas := []AreaConnection{}

	for _, connArea := range area.ConnectedAreas {
		connection := AreaConnection{
			Area:           locAreaToAreaAPIResource(cfg, i, connArea.LocationArea),
			ConnectionType: connArea.ConnectionType,
			IsStoryBased:   connArea.IsStoryBased,
			Notes:          connArea.Notes,
		}

		connectedAreas = append(connectedAreas, connection)
	}

	return connectedAreas, nil
}

func getAreaDisplayName(area seeding.Area) string {
	sublocName := area.Sublocation.Name

	if sublocName == area.Name {
		return area.Name
	}

	return fmt.Sprintf("%s - %s", sublocName, area.Name)
}
