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

func getAreaRelationships(srv *Server, r *http.Request, area seedtypes.Area) (LocRel, error) {
	var rel LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := E.Areas.queryAvlRelGetParams(srv, r, area.ID)
	if err != nil {
		return LocRel{}, err
	}

	g.Go(func() error {
		var err error
		rel.Characters, err = E.Characters.dbGetRelations(srv, ctx, area, ToIntManyNull(srv.Cfg.DB.GetAreaCharacterIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Aeons, err = E.Aeons.dbGetRelations(srv, ctx, area, ToIntManyNull(srv.Cfg.DB.GetAreaAeonIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = E.Shops.queryAvlRelRun(srv, ctx, area, availabilityParams, getAreaRelSourceIDs(srv, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = E.Treasures.queryAvlRelRun(srv, ctx, area, availabilityParams, getAreaRelSourceIDs(srv, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = E.Monsters.queryAvlRelRun(srv, ctx, area, availabilityParams, getAreaRelSourceIDs(srv, ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Formations, err = E.MonsterFormations.queryAvlRelRun(srv, ctx, area, availabilityParams, getAreaRelSourceIDs(srv, ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Quests, err = E.Quests.queryAvlRelRun(srv, ctx, area, availabilityParams, getAreaRelSourceIDs(srv, ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Music, err = getMusicLocBased(srv, ctx, area, LocBasedMusicQueries{
			CueSongs:  ToIntManyNull(srv.Cfg.DB.GetAreaCueSongIDs),
			BmSongs:   srv.Cfg.DB.GetAreaBackgroundMusicSongIDs,
			FMVSongs:  srv.Cfg.DB.GetAreaFMVSongIDs,
			BossMusic: srv.Cfg.DB.GetAreaBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = E.FMVs.dbGetRelations(srv, ctx, area, srv.Cfg.DB.GetAreaFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return LocRel{}, err
	}

	return rel, nil
}

func getAreaConnectedAreas(srv *Server, area seedtypes.Area) ([]AreaConnection, error) {
	i := E.Areas
	connectedAreas := []AreaConnection{}

	for _, connArea := range area.ConnectedAreas {
		connection := AreaConnection{
			Area:           locAreaToAreaAPIResource(srv, i, connArea.LocationArea),
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
