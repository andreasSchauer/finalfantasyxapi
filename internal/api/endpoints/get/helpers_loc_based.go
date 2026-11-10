package get

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

type LocBasedMusicQueries struct {
	CueSongs  handlerinput.DBQueryIntMany
	BmSongs   handlerinput.DBQueryIntMany
	FMVSongs  handlerinput.DBQueryIntMany
	BossMusic handlerinput.DBQueryIntMany
}

func getMusicLocBased(ctx context.Context, item seedtypes.Lookupable, queries LocBasedMusicQueries) (*apitypes.LocBasedMusic, error) {
	i := apilookups.E.Songs
	var music apitypes.LocBasedMusic
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		music.Cues, err = i.DBGetRelations(ctx, item, queries.CueSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.BackgroundMusic, err = i.DBGetRelations(ctx, item, queries.BmSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.FMVs, err = i.DBGetRelations(ctx, item, queries.FMVSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.BossMusic, err = i.DBGetRelations(ctx, item, queries.BossMusic)
		return err
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return &music, nil
}
