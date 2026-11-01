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

	availabilityParams, err := cfg.E.Locations.queryAvlRelGetParams(cfg, r, location.ID)
	if err != nil {
		return Location{}, err
	}

	g.Go(func() error {
		var err error
		rel.ConnectedLocations, err = cfg.E.Locations.dbGetRelations(cfg, ctx, location, cfg.db.GetConnectedLocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Sublocations, err = cfg.E.Sublocations.dbGetRelations(cfg, ctx, location, cfg.db.GetLocationSublocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Characters, err = cfg.E.Characters.dbGetRelations(cfg, ctx, location, cfg.db.GetLocationCharacterIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Aeons, err = cfg.E.Aeons.dbGetRelations(cfg, ctx, location, cfg.db.GetLocationAeonIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Shops, err = cfg.E.Shops.queryAvlRelRun(cfg, ctx, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Treasures, err = cfg.E.Treasures.queryAvlRelRun(cfg, ctx, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Monsters, err = cfg.E.Monsters.queryAvlRelRun(cfg, ctx, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Formations, err = cfg.E.MonsterFormations.queryAvlRelRun(cfg, ctx, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Quests, err = cfg.E.Quests.queryAvlRelRun(cfg, ctx, location, availabilityParams, getLocationRelSourceIDs(cfg, ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Music, err = getMusicLocBased(cfg, ctx, location, LocBasedMusicQueries{
			CueSongs:  cfg.db.GetLocationCueSongIDs,
			BmSongs:   cfg.db.GetLocationBackgroundMusicSongIDs,
			FMVSongs:  cfg.db.GetLocationFMVSongIDs,
			BossMusic: cfg.db.GetLocationBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		locRel.FMVs, err = cfg.E.FMVs.dbGetRelations(cfg, ctx, location, cfg.db.GetLocationFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return Location{}, err
	}

	rel.LocRel = locRel
	return rel, nil
}
