package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetLocationRelationships(r *http.Request, location seedtypes.Location) (apitypes.Location, error) {
	var rel apitypes.Location
	var locRel apitypes.LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := apilookups.E.Locations.QueryAvlRelGetParams(r, location.ID)
	if err != nil {
		return apitypes.Location{}, err
	}

	g.Go(func() error {
		var err error
		rel.ConnectedLocations, err = apilookups.E.Locations.DBGetRelations(ctx, location, config.Cfg.DB.GetConnectedLocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Sublocations, err = apilookups.E.Sublocations.DBGetRelations(ctx, location, config.Cfg.DB.GetLocationSublocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Characters, err = apilookups.E.Characters.DBGetRelations(ctx, location, config.Cfg.DB.GetLocationCharacterIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Aeons, err = apilookups.E.Aeons.DBGetRelations(ctx, location, config.Cfg.DB.GetLocationAeonIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Shops, err = apilookups.E.Shops.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Treasures, err = apilookups.E.Treasures.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Monsters, err = apilookups.E.Monsters.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Formations, err = apilookups.E.MonsterFormations.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Quests, err = apilookups.E.Quests.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Music, err = getMusicLocBased(ctx, location, LocBasedMusicQueries{
			CueSongs:  config.Cfg.DB.GetLocationCueSongIDs,
			BmSongs:   config.Cfg.DB.GetLocationBackgroundMusicSongIDs,
			FMVSongs:  config.Cfg.DB.GetLocationFMVSongIDs,
			BossMusic: config.Cfg.DB.GetLocationBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		locRel.FMVs, err = apilookups.E.FMVs.DBGetRelations(ctx, location, config.Cfg.DB.GetLocationFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return apitypes.Location{}, err
	}

	rel.LocRel = locRel
	return rel, nil
}
