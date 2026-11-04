package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getSublocationRelationships(server *Server, r *http.Request, sublocation seedtypes.Sublocation) (Sublocation, error) {
	var rel Sublocation
	var locRel LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := E.Sublocations.queryAvlRelGetParams(server, r, sublocation.ID)
	if err != nil {
		return Sublocation{}, err
	}

	g.Go(func() error {
		var err error
		rel.ConnectedSublocations, err = E.Sublocations.dbGetRelations(server, r.Context(), sublocation, server.Cfg.DB.GetConnectedSublocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = E.Areas.dbGetRelations(server, r.Context(), sublocation, server.Cfg.DB.GetSublocationAreaIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Characters, err = E.Characters.dbGetRelations(server, ctx, sublocation, server.Cfg.DB.GetSublocationCharacterIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Aeons, err = E.Aeons.dbGetRelations(server, ctx, sublocation, server.Cfg.DB.GetSublocationAeonIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Shops, err = E.Shops.queryAvlRelRun(server, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(server, ViewSourceTypeShop))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Treasures, err = E.Treasures.queryAvlRelRun(server, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(server, ViewSourceTypeTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Monsters, err = E.Monsters.queryAvlRelRun(server, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(server, ViewSourceTypeMonster))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Formations, err = E.MonsterFormations.queryAvlRelRun(server, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(server, ViewSourceTypeMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Quests, err = E.Quests.queryAvlRelRun(server, ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(server, ViewSourceTypeQuest))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Music, err = getMusicLocBased(server, ctx, sublocation, LocBasedMusicQueries{
			CueSongs:  server.Cfg.DB.GetSublocationCueSongIDs,
			BmSongs:   server.Cfg.DB.GetSublocationBackgroundMusicSongIDs,
			FMVSongs:  server.Cfg.DB.GetSublocationFMVSongIDs,
			BossMusic: server.Cfg.DB.GetSublocationBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		locRel.FMVs, err = E.FMVs.dbGetRelations(server, ctx, sublocation, server.Cfg.DB.GetSublocationFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return Sublocation{}, err
	}

	rel.LocRel = locRel
	return rel, nil
}
