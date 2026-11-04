package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

type LocRel struct {
	Characters []NamedAPIResource   `json:"characters"`
	Aeons      []NamedAPIResource   `json:"aeons"`
	Shops      []UnnamedAPIResource `json:"shops"`
	Treasures  []UnnamedAPIResource `json:"treasures"`
	Monsters   []NamedAPIResource   `json:"monsters"`
	Formations []UnnamedAPIResource `json:"monster_formations"`
	Quests     []QuestAPIResource   `json:"quests"`
	Music      *LocBasedMusic       `json:"music"`
	FMVs       []NamedAPIResource   `json:"fmvs"`
}

type LocBasedMusic struct {
	BackgroundMusic []NamedAPIResource `json:"background_music"`
	Cues            []NamedAPIResource `json:"cues"`
	FMVs            []NamedAPIResource `json:"fmvs"`
	BossMusic       []NamedAPIResource `json:"boss_fights"`
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

func getMusicLocBased(srv *Server, ctx context.Context, item seedtypes.Lookupable, queries LocBasedMusicQueries) (*LocBasedMusic, error) {
	i := E.Songs
	var music LocBasedMusic
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		music.Cues, err = i.dbGetRelations(srv, ctx, item, queries.CueSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.BackgroundMusic, err = i.dbGetRelations(srv, ctx, item, queries.BmSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.FMVs, err = i.dbGetRelations(srv, ctx, item, queries.FMVSongs)
		return err
	})

	g.Go(func() error {
		var err error
		music.BossMusic, err = i.dbGetRelations(srv, ctx, item, queries.BossMusic)
		return err
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	return &music, nil
}
