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

func getAreaRelationships(server *Server, r *http.Request, area seedtypes.Area) (LocRel, error) {
	var rel LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := E.Areas.queryAvlRelGetParams(server, r, area.ID)
	if err != nil {
		return LocRel{}, err
	}

	g.Go(func() error {
		var err error
		rel.Characters, err = E.Characters.dbGetRelations(server, ctx, area, ToIntManyNull(server.Cfg.DB.GetAreaCharacterIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Aeons, err = E.Aeons.dbGetRelations(server, ctx, area, ToIntManyNull(server.Cfg.DB.GetAreaAeonIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = E.Shops.queryAvlRelRun(server, ctx, area, availabilityParams, getAreaRelSourceIDs(server, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = E.Treasures.queryAvlRelRun(server, ctx, area, availabilityParams, getAreaRelSourceIDs(server, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = E.Monsters.queryAvlRelRun(server, ctx, area, availabilityParams, getAreaRelSourceIDs(server, ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Formations, err = E.MonsterFormations.queryAvlRelRun(server, ctx, area, availabilityParams, getAreaRelSourceIDs(server, ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Quests, err = E.Quests.queryAvlRelRun(server, ctx, area, availabilityParams, getAreaRelSourceIDs(server, ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Music, err = getMusicLocBased(server, ctx, area, LocBasedMusicQueries{
			CueSongs:  ToIntManyNull(server.Cfg.DB.GetAreaCueSongIDs),
			BmSongs:   server.Cfg.DB.GetAreaBackgroundMusicSongIDs,
			FMVSongs:  server.Cfg.DB.GetAreaFMVSongIDs,
			BossMusic: server.Cfg.DB.GetAreaBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = E.FMVs.dbGetRelations(server, ctx, area, server.Cfg.DB.GetAreaFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return LocRel{}, err
	}

	return rel, nil
}

func getAreaConnectedAreas(server *Server, area seedtypes.Area) ([]AreaConnection, error) {
	i := E.Areas
	connectedAreas := []AreaConnection{}

	for _, connArea := range area.ConnectedAreas {
		connection := AreaConnection{
			Area:           locAreaToAreaAPIResource(server, i, connArea.LocationArea),
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
