package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
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

func getAreaRelationships(cfg *Config, r *http.Request, area seedtypes.Area) (LocRel, error) {
	var rel LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := cfg.E.Areas.queryAvlRelGetParams(cfg, r, area.ID)
	if err != nil {
		return LocRel{}, err
	}

	g.Go(func() error {
		var err error
		rel.Characters, err = cfg.E.Characters.dbGetRelations(cfg, ctx, area, ToIntManyNull(cfg.db.GetAreaCharacterIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Aeons, err = cfg.E.Aeons.dbGetRelations(cfg, ctx, area, ToIntManyNull(cfg.db.GetAreaAeonIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = cfg.E.Shops.queryAvlRelRun(cfg, ctx, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = cfg.E.Treasures.queryAvlRelRun(cfg, ctx, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = cfg.E.Monsters.queryAvlRelRun(cfg, ctx, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Formations, err = cfg.E.MonsterFormations.queryAvlRelRun(cfg, ctx, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Quests, err = cfg.E.Quests.queryAvlRelRun(cfg, ctx, area, availabilityParams, getAreaRelSourceIDs(cfg, ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Music, err = getMusicLocBased(cfg, ctx, area, LocBasedMusicQueries{
			CueSongs:  ToIntManyNull(cfg.db.GetAreaCueSongIDs),
			BmSongs:   cfg.db.GetAreaBackgroundMusicSongIDs,
			FMVSongs:  cfg.db.GetAreaFMVSongIDs,
			BossMusic: cfg.db.GetAreaBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = cfg.E.FMVs.dbGetRelations(cfg, ctx, area, cfg.db.GetAreaFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return LocRel{}, err
	}

	return rel, nil
}

func getAreaConnectedAreas(cfg *Config, area seedtypes.Area) ([]AreaConnection, error) {
	i := cfg.E.Areas
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

func getAreaDisplayName(area seedtypes.Area) string {
	sublocName := area.Sublocation

	if sublocName == area.Name {
		return area.Name
	}

	return fmt.Sprintf("%s - %s", sublocName, area.Name)
}
