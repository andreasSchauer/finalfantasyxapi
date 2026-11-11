package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"golang.org/x/sync/errgroup"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Sublocation(r *http.Request, i handlerinput.Main[seedtypes.Sublocation, apitypes.Sublocation, apiresource.Named], id int32) (apitypes.Sublocation, error) {
	sublocation, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Sublocation{}, err
	}

	rel, err := getSublocationRelationships(r, sublocation)
	if err != nil {
		return apitypes.Sublocation{}, err
	}

	response := apitypes.Sublocation{
		ID:                    sublocation.ID,
		Name:                  sublocation.Name,
		ParentLocation:        handlerinput.NameToNamedAPIResource(apilookups.E.Locations, sublocation.Location, nil),
		Availability:          sublocation.Availability,
		ConnectedSublocations: rel.ConnectedSublocations,
		Areas:                 rel.Areas,
		LocRel:                rel.LocRel,
	}

	return response, nil
}

func getSublocationRelationships(r *http.Request, sublocation seedtypes.Sublocation) (apitypes.Sublocation, error) {
	var rel apitypes.Sublocation
	var locRel apitypes.LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := apilookups.E.Sublocations.QueryAvlRelGetParams(r, sublocation.ID)
	if err != nil {
		return apitypes.Sublocation{}, err
	}

	g.Go(func() error {
		var err error
		rel.ConnectedSublocations, err = apilookups.E.Sublocations.DBGetRelations(r.Context(), sublocation, config.Cfg.DB.GetConnectedSublocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Areas, err = apilookups.E.Areas.DBGetRelations(r.Context(), sublocation, config.Cfg.DB.GetSublocationAreaIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Characters, err = apilookups.E.Characters.DBGetRelations(ctx, sublocation, config.Cfg.DB.GetSublocationCharacterIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Aeons, err = apilookups.E.Aeons.DBGetRelations(ctx, sublocation, config.Cfg.DB.GetSublocationAeonIDs)
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Shops, err = apilookups.E.Shops.QueryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ts.VstShop))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Treasures, err = apilookups.E.Treasures.QueryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ts.VstTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Monsters, err = apilookups.E.Monsters.QueryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ts.VstMonster))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Formations, err = apilookups.E.MonsterFormations.QueryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ts.VstMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Quests, err = apilookups.E.Quests.QueryAvlRelRun(ctx, sublocation, availabilityParams, getSublocationRelSourceIDs(ts.VstQuest))
		return err
	})

	g.Go(func() error {
		var err error
		locRel.Music, err = getMusicLocBased(ctx, sublocation, locBasedMusicQueries{
			cueSongs:  config.Cfg.DB.GetSublocationCueSongIDs,
			bmSongs:   config.Cfg.DB.GetSublocationBackgroundMusicSongIDs,
			fmvSongs:  config.Cfg.DB.GetSublocationFMVSongIDs,
			bossMusic: config.Cfg.DB.GetSublocationBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		locRel.FMVs, err = apilookups.E.FMVs.DBGetRelations(ctx, sublocation, config.Cfg.DB.GetSublocationFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return apitypes.Sublocation{}, err
	}

	rel.LocRel = locRel
	return rel, nil
}
