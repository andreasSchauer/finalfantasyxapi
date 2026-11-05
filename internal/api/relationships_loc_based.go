package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

type LocRel struct {
	Characters []apiresource.NamedAPIResource   `json:"characters"`
	Aeons      []apiresource.NamedAPIResource   `json:"aeons"`
	Shops      []apiresource.UnnamedAPIResource `json:"shops"`
	Treasures  []apiresource.UnnamedAPIResource `json:"treasures"`
	Monsters   []apiresource.NamedAPIResource   `json:"monsters"`
	Formations []apiresource.UnnamedAPIResource `json:"monster_formations"`
	Quests     []apiresource.QuestAPIResource   `json:"quests"`
	Music      *LocBasedMusic                   `json:"music"`
	FMVs       []apiresource.NamedAPIResource   `json:"fmvs"`
}

type LocBasedMusic struct {
	BackgroundMusic []apiresource.NamedAPIResource `json:"background_music"`
	Cues            []apiresource.NamedAPIResource `json:"cues"`
	FMVs            []apiresource.NamedAPIResource `json:"fmvs"`
	BossMusic       []apiresource.NamedAPIResource `json:"boss_fights"`
}

func (m LocBasedMusic) IsZero() bool {
	return len(m.BackgroundMusic) == 0 &&
		len(m.Cues) == 0 &&
		len(m.FMVs) == 0 &&
		len(m.BossMusic) == 0
}

type LocBasedMusicQueries struct {
	CueSongs  DBQueryIntMany
	BmSongs   DBQueryIntMany
	FMVSongs  DBQueryIntMany
	BossMusic DBQueryIntMany
}

func getMusicLocBased(ctx context.Context, item seedtypes.Lookupable, queries LocBasedMusicQueries) (*LocBasedMusic, error) {
	i := E.Songs
	var music LocBasedMusic
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
