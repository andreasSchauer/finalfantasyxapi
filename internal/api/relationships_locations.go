package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getLocationRelationships(cfg *Config, r *http.Request, location seeding.Location) (Location, error) {
	var rel Location
	var locRel LocRel
	g, ctx := errgroup.WithContext(r.Context())
	
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.locations, location.ID)
	if err != nil {
		return Location{}, err
	}

	g.Go(func() error {
		var err error
		rel.ConnectedLocations, err = getResourcesDbItem(cfg, ctx, cfg.e.locations, location, cfg.db.GetConnectedLocationIDs)
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.Sublocations, err = getResourcesDbItem(cfg, ctx, cfg.e.sublocations, location, cfg.db.GetLocationSublocationIDs)
		return err
	})

	g.Go(func() error{
		var err error
		locRel.Characters, err = getResourcesDbItem(cfg, ctx, cfg.e.characters, location, cfg.db.GetLocationCharacterIDs)
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Aeons, err = getResourcesDbItem(cfg, ctx, cfg.e.aeons, location, cfg.db.GetLocationAeonIDs)
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Shops, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.shops, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Treasures, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.treasures, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeTreasure))
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Monsters, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.monsters, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeMonster))
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Formations, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.monsterFormations, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeMonsterFormation))
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Quests, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.quests, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeQuest))
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Music, err = getMusicLocBased(cfg, ctx, location, LocBasedMusicQueries{
			CueSongs:  cfg.db.GetLocationCueSongIDs,
			BmSongs:   cfg.db.GetLocationBackgroundMusicSongIDs,
			FMVSongs:  cfg.db.GetLocationFMVSongIDs,
			BossMusic: cfg.db.GetLocationBossSongIDs,
		})
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.FMVs, err = getResourcesDbItem(cfg, ctx, cfg.e.fmvs, location, cfg.db.GetLocationFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return Location{}, err
	}

	rel.LocRel = locRel
	return rel, nil
}
