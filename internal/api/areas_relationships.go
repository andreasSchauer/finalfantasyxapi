package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AreaConnection struct {
	Area           AreaAPIResource  `json:"area"`
	ConnectionType NamedAPIResource `json:"connection_type"`
	StoryOnly      bool             `json:"story_only"`
	Notes          *string          `json:"notes,omitempty"`
}

func (ac AreaConnection) GetAPIResource() APIResource {
	return ac.Area
}

func getAreaRelationships(cfg *Config, r *http.Request, area seeding.Area) (LocRel, error) {
	characters, err := getResourcesDB(cfg, r, cfg.e.characters, area, cfg.db.GetAreaCharacterIDs)
	if err != nil {
		return LocRel{}, err
	}

	aeons, err := getResourcesDB(cfg, r, cfg.e.aeons, area, cfg.db.GetAreaAeonIDs)
	if err != nil {
		return LocRel{}, err
	}

	shops, err := getResourcesDB(cfg, r, cfg.e.shops, area, cfg.db.GetAreaShopIDs)
	if err != nil {
		return LocRel{}, err
	}

	treasures, err := getResourcesDB(cfg, r, cfg.e.treasures, area, cfg.db.GetAreaTreasureIDs)
	if err != nil {
		return LocRel{}, err
	}

	monsters, err := getResourcesDB(cfg, r, cfg.e.monsters, area, cfg.db.GetAreaMonsterIDs)
	if err != nil {
		return LocRel{}, err
	}

	formations, err := getResourcesDB(cfg, r, cfg.e.monsterFormations, area, cfg.db.GetAreaMonsterFormationIDs)
	if err != nil {
		return LocRel{}, err
	}

	sidequests, err := getLocBasedSidequests(cfg, r, area, cfg.db.GetAreaQuestIDs)
	if err != nil {
		return LocRel{}, err
	}

	music, err := getMusicLocBased(cfg, r, area, LocBasedMusicQueries{
		CueSongs:  cfg.db.GetAreaCueSongIDs,
		BmSongs:   cfg.db.GetAreaBackgroundMusicSongIDs,
		FMVSongs:  cfg.db.GetAreaFMVSongIDs,
		BossMusic: cfg.db.GetAreaBossSongIDs,
	})
	if err != nil {
		return LocRel{}, err
	}

	fmvs, err := getResourcesDB(cfg, r, cfg.e.fmvs, area, cfg.db.GetAreaFmvIDs)
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
		connType, err := newNamedAPIResourceFromType(cfg, cfg.e.connectionType.endpoint, connArea.ConnectionType, cfg.t.AreaConnectionType)
		if err != nil {
			return nil, err
		}

		connection := AreaConnection{
			Area:           locAreaToAreaAPIResource(cfg, i, connArea.LocationArea),
			ConnectionType: connType,
			StoryOnly:      connArea.StoryOnly,
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
