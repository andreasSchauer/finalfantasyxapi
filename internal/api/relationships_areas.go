package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getAreaRelationships(r *http.Request, area seedtypes.Area) (apitypes.LocRel, error) {
	var rel apitypes.LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := apilookups.E.Areas.QueryAvlRelGetParams(r, area.ID)
	if err != nil {
		return apitypes.LocRel{}, err
	}

	g.Go(func() error {
		var err error
		rel.Characters, err = apilookups.E.Characters.DBGetRelations(ctx, area, handlerinput.ToIntManyNull(config.Cfg.DB.GetAreaCharacterIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Aeons, err = apilookups.E.Aeons.DBGetRelations(ctx, area, handlerinput.ToIntManyNull(config.Cfg.DB.GetAreaAeonIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = apilookups.E.Shops.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = apilookups.E.Treasures.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = apilookups.E.Monsters.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Formations, err = apilookups.E.MonsterFormations.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Quests, err = apilookups.E.Quests.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Music, err = getMusicLocBased(ctx, area, LocBasedMusicQueries{
			CueSongs:  handlerinput.ToIntManyNull(config.Cfg.DB.GetAreaCueSongIDs),
			BmSongs:   config.Cfg.DB.GetAreaBackgroundMusicSongIDs,
			FMVSongs:  config.Cfg.DB.GetAreaFMVSongIDs,
			BossMusic: config.Cfg.DB.GetAreaBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = apilookups.E.FMVs.DBGetRelations(ctx, area, config.Cfg.DB.GetAreaFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return apitypes.LocRel{}, err
	}

	return rel, nil
}

func getAreaConnectedAreas(area seedtypes.Area) ([]apitypes.AreaConnection, error) {
	i := apilookups.E.Areas
	connectedAreas := []apitypes.AreaConnection{}

	for _, connArea := range area.ConnectedAreas {
		connection := apitypes.AreaConnection{
			Area:           handlerinput.LocAreaToAreaAPIResource(i, connArea.LocationArea),
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
