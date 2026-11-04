package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getSublocationRelationships(cfg *Config, r *http.Request, sublocation seedtypes.Sublocation) (Sublocation, error) {
	var rel Sublocation
	var locRel LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := E.Sublocations.queryAvlRelGetParams(cfg, r, sublocation.ID)
	if err != nil {
		return Sublocation{}, err
	}

	g.Go(func() error {
		var err error
		rel.ConnectedSublocations, err = E.Sublocations.dbGetRelations(cfg, r.Context(), sublocation, cfg.db.GetConnectedSublocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = E.Areas.dbGetRelations(cfg, r.Context(), sublocation, cfg.db.GetSublocationAreaIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Characters, err = E.Characters.dbGetRelations(cfg, ctx, sublocation, cfg.db.GetSublocationCharacterIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Aeons, err = E.Aeons.dbGetRelations(cfg, ctx, sublocation, cfg.db.GetSublocationAeonIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Shops, err = E.Shops.queryAvlRelRun(cfg, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Treasures, err = E.Treasures.queryAvlRelRun(cfg, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Monsters, err = E.Monsters.queryAvlRelRun(cfg, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Formations, err = E.MonsterFormations.queryAvlRelRun(cfg, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Quests, err = E.Quests.queryAvlRelRun(cfg, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(cfg, ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Music, err = getMusicLocBased(cfg, ctx, sublocation, LocBasedMusicQueries{
			CueSongs:  cfg.db.GetSublocationCueSongIDs,
			BmSongs:   cfg.db.GetSublocationBackgroundMusicSongIDs,
			FMVSongs:  cfg.db.GetSublocationFMVSongIDs,
			BossMusic: cfg.db.GetSublocationBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		locRel.FMVs, err = E.FMVs.dbGetRelations(cfg, ctx, sublocation, cfg.db.GetSublocationFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return Sublocation{}, err
	}

	rel.LocRel = locRel
	return rel, nil
}
