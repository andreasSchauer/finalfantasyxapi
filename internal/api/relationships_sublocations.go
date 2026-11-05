package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getSublocationRelationships(r *http.Request, sublocation seedtypes.Sublocation) (Sublocation, error) {
	var rel Sublocation
	var locRel LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := E.Sublocations.queryAvlRelGetParams(r, sublocation.ID)
	if err != nil {
		return Sublocation{}, err
	}

	g.Go(func() error {
		var err error
		rel.ConnectedSublocations, err = E.Sublocations.dbGetRelations(r.Context(), sublocation, config.Cfg.DB.GetConnectedSublocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = E.Areas.dbGetRelations(r.Context(), sublocation, config.Cfg.DB.GetSublocationAreaIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Characters, err = E.Characters.dbGetRelations(ctx, sublocation, config.Cfg.DB.GetSublocationCharacterIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Aeons, err = E.Aeons.dbGetRelations(ctx, sublocation, config.Cfg.DB.GetSublocationAeonIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Shops, err = E.Shops.queryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Treasures, err = E.Treasures.queryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Monsters, err = E.Monsters.queryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Formations, err = E.MonsterFormations.queryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Quests, err = E.Quests.queryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Music, err = getMusicLocBased(ctx, sublocation, LocBasedMusicQueries{
			CueSongs:  config.Cfg.DB.GetSublocationCueSongIDs,
			BmSongs:   config.Cfg.DB.GetSublocationBackgroundMusicSongIDs,
			FMVSongs:  config.Cfg.DB.GetSublocationFMVSongIDs,
			BossMusic: config.Cfg.DB.GetSublocationBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		locRel.FMVs, err = E.FMVs.dbGetRelations(ctx, sublocation, config.Cfg.DB.GetSublocationFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return Sublocation{}, err
	}

	rel.LocRel = locRel
	return rel, nil
}
