package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getAreaRelationships(r *http.Request, area seedtypes.Area) (apitypes.LocRel, error) {
	var rel apitypes.LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := E.Areas.QueryAvlRelGetParams(r, area.ID)
	if err != nil {
		return apitypes.LocRel{}, err
	}

	g.Go(func() error {
		var err error
		rel.Characters, err = E.Characters.DBGetRelations(ctx, area, ToIntManyNull(config.Cfg.DB.GetAreaCharacterIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Aeons, err = E.Aeons.DBGetRelations(ctx, area, ToIntManyNull(config.Cfg.DB.GetAreaAeonIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = E.Shops.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = E.Treasures.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = E.Monsters.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Formations, err = E.MonsterFormations.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Quests, err = E.Quests.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Music, err = getMusicLocBased(ctx, area, LocBasedMusicQueries{
			CueSongs:  ToIntManyNull(config.Cfg.DB.GetAreaCueSongIDs),
			BmSongs:   config.Cfg.DB.GetAreaBackgroundMusicSongIDs,
			FMVSongs:  config.Cfg.DB.GetAreaFMVSongIDs,
			BossMusic: config.Cfg.DB.GetAreaBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = E.FMVs.DBGetRelations(ctx, area, config.Cfg.DB.GetAreaFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return apitypes.LocRel{}, err
	}

	return rel, nil
}

func getAreaConnectedAreas(area seedtypes.Area) ([]apitypes.AreaConnection, error) {
	i := E.Areas
	connectedAreas := []apitypes.AreaConnection{}

	for _, connArea := range area.ConnectedAreas {
		connection := apitypes.AreaConnection{
			Area:           LocAreaToAreaAPIResource(i, connArea.LocationArea),
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
