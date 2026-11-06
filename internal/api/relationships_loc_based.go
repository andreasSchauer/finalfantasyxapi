package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

type LocBasedMusicQueries struct {
	CueSongs  DBQueryIntMany
	BmSongs   DBQueryIntMany
	FMVSongs  DBQueryIntMany
	BossMusic DBQueryIntMany
}

func getMusicLocBased(ctx context.Context, item seedtypes.Lookupable, queries LocBasedMusicQueries) (*apitypes.LocBasedMusic, error) {
	i := E.Songs
	var music apitypes.LocBasedMusic
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		music.Cues, err = i.dbGetRelations(ctx, item, queries.CueSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.BackgroundMusic, err = i.dbGetRelations(ctx, item, queries.BmSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.FMVs, err = i.dbGetRelations(ctx, item, queries.FMVSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.BossMusic, err = i.dbGetRelations(ctx, item, queries.BossMusic)
		return err
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return &music, nil
}
