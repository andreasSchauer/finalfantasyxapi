package get

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

type locBasedMusicQueries struct {
	cueSongs  handlerinput.DBQueryIntMany
	bmSongs   handlerinput.DBQueryIntMany
	fmvSongs  handlerinput.DBQueryIntMany
	bossMusic handlerinput.DBQueryIntMany
}

func getMusicLocBased(ctx context.Context, item seedtypes.Lookupable, queries locBasedMusicQueries) (*apitypes.LocBasedMusic, error) {
	i := apilookups.E.Songs
	var music apitypes.LocBasedMusic
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		music.Cues, err = i.DBGetRelations(ctx, item, queries.cueSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.BackgroundMusic, err = i.DBGetRelations(ctx, item, queries.bmSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.FMVs, err = i.DBGetRelations(ctx, item, queries.fmvSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.BossMusic, err = i.DBGetRelations(ctx, item, queries.bossMusic)
		return err
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return &music, nil
}
