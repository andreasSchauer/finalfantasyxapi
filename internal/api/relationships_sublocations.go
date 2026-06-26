package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getSublocationRelationships(cfg *Config, r *http.Request, sublocation seeding.Sublocation) (Sublocation, error) {
	var rel Sublocation
	var locRel LocRel
	g, ctx := errgroup.WithContext(r.Context())
	
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.sublocations, sublocation.ID)
	if err != nil {
		return Sublocation{}, err
	}

	g.Go(func() error {
		var err error
		rel.ConnectedSublocations, err = getResourcesDbItem(cfg, r.Context(), cfg.e.sublocations, sublocation, cfg.db.GetConnectedSublocationIDs)
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.Areas, err = getResourcesDbItem(cfg, r.Context(), cfg.e.areas, sublocation, cfg.db.GetSublocationAreaIDs)
		return err
	})

	g.Go(func() error{
		var err error
		locRel.Characters, err = getResourcesDbItem(cfg, ctx, cfg.e.characters, sublocation, cfg.db.GetSublocationCharacterIDs)
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Aeons, err = getResourcesDbItem(cfg, ctx, cfg.e.aeons, sublocation, cfg.db.GetSublocationAeonIDs)
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Shops, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.shops, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Treasures, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.treasures, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeTreasure))
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Monsters, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.monsters, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeMonster))
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Formations, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.monsterFormations, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error{
		var err error
		locRel.Quests, err = runRelAvailabilityQuery(cfg, ctx, cfg.e.quests, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeQuest))
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.Music, err = getMusicLocBased(cfg, ctx, sublocation, LocBasedMusicQueries{
			CueSongs:  cfg.db.GetSublocationCueSongIDs,
			BmSongs:   cfg.db.GetSublocationBackgroundMusicSongIDs,
			FMVSongs:  cfg.db.GetSublocationFMVSongIDs,
			BossMusic: cfg.db.GetSublocationBossSongIDs,
		})
		return err
	})
	
	g.Go(func() error{
		var err error
		locRel.FMVs, err = getResourcesDbItem(cfg, ctx, cfg.e.fmvs, sublocation, cfg.db.GetSublocationFmvIDs)
		return err
	})
	
	err = g.Wait()
	if err != nil {
		return Sublocation{}, err
	}

	rel.LocRel = locRel
	return rel, nil
}
