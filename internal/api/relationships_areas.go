package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
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
	var rel LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.areas, area.ID)
	if err != nil {
		return LocRel{}, err
	}

	g.Go(func() error{
		var err error
		rel.Characters, err = getResourcesDbItem(cfg, ctx, cfg.e.characters, area, ToIntManyNull(cfg.db.GetAreaCharacterIDs))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.Aeons, err = getResourcesDbItem(cfg, ctx, cfg.e.aeons, area, ToIntManyNull(cfg.db.GetAreaAeonIDs))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.Shops, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.shops, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.Treasures, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.treasures, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeTreasure))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.Monsters, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.monsters, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeMonster))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.Formations, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.monsterFormations, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeMonsterFormation))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.Quests, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.quests, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeQuest))
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.Music, err = getMusicLocBased(cfg, ctx, area, LocBasedMusicQueries{
			CueSongs:  ToIntManyNull(cfg.db.GetAreaCueSongIDs),
			BmSongs:   cfg.db.GetAreaBackgroundMusicSongIDs,
			FMVSongs:  cfg.db.GetAreaFMVSongIDs,
			BossMusic: cfg.db.GetAreaBossSongIDs,
		})
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.FMVs, err = getResourcesDbItem(cfg, ctx, cfg.e.fmvs, area, cfg.db.GetAreaFmvIDs)
		return err
	})
	
	err = g.Wait()
	if err != nil {
		return LocRel{}, err
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
