package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getLocationRelationships(srv *Server, r *http.Request, location seedtypes.Location) (Location, error) {
	var rel Location
	var locRel LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := E.Locations.queryAvlRelGetParams(srv, r, location.ID)
	if err != nil {
		return Location{}, err
	}

	g.Go(func() error {
		var err error
		rel.ConnectedLocations, err = E.Locations.dbGetRelations(srv, ctx, location, config.Cfg.DB.GetConnectedLocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Sublocations, err = E.Sublocations.dbGetRelations(srv, ctx, location, config.Cfg.DB.GetLocationSublocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Characters, err = E.Characters.dbGetRelations(srv, ctx, location, config.Cfg.DB.GetLocationCharacterIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Aeons, err = E.Aeons.dbGetRelations(srv, ctx, location, config.Cfg.DB.GetLocationAeonIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Shops, err = E.Shops.queryAvlRelRun(srv, ctx, location, availabilityParams, getLocationRelSourceIDs(srv, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Treasures, err = E.Treasures.queryAvlRelRun(srv, ctx, location, availabilityParams, getLocationRelSourceIDs(srv, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Monsters, err = E.Monsters.queryAvlRelRun(srv, ctx, location, availabilityParams, getLocationRelSourceIDs(srv, ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Formations, err = E.MonsterFormations.queryAvlRelRun(srv, ctx, location, availabilityParams, getLocationRelSourceIDs(srv, ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Quests, err = E.Quests.queryAvlRelRun(srv, ctx, location, availabilityParams, getLocationRelSourceIDs(srv, ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Music, err = getMusicLocBased(srv, ctx, location, LocBasedMusicQueries{
			CueSongs:  config.Cfg.DB.GetLocationCueSongIDs,
			BmSongs:   config.Cfg.DB.GetLocationBackgroundMusicSongIDs,
			FMVSongs:  config.Cfg.DB.GetLocationFMVSongIDs,
			BossMusic: config.Cfg.DB.GetLocationBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		locRel.FMVs, err = E.FMVs.dbGetRelations(srv, ctx, location, config.Cfg.DB.GetLocationFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return Location{}, err
	}

	rel.LocRel = locRel
	return rel, nil
}
