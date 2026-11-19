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

func Location(r *http.Request, i handlerinput.Main[seedtypes.Location, apitypes.Location, apiresource.Named], id int32) (apitypes.Location, error) {
	location, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Location{}, err
	}
	
	response := apitypes.Location{
		ID:                 location.ID,
		Name:               location.Name,
		Availability:       location.Availability,
	}

	err = getLocationRelationships(r, location, &response)
	if err != nil {
		return apitypes.Location{}, err
	}

	return response, nil
}

func getLocationRelationships(r *http.Request, location seedtypes.Location, response *apitypes.Location) error {
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := apilookups.E.Locations.QueryAvlRelGetParams(r, location.ID)
	if err != nil {
		return err
	}

	g.Go(func() error {
		var err error
		response.ConnectedLocations, err = apilookups.E.Locations.DBGetRelations(ctx, location, config.Cfg.DB.GetConnectedLocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.Sublocations, err = apilookups.E.Sublocations.DBGetRelations(ctx, location, config.Cfg.DB.GetLocationSublocationIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.Characters, err = apilookups.E.Characters.DBGetRelations(ctx, location, config.Cfg.DB.GetLocationCharacterIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.Aeons, err = apilookups.E.Aeons.DBGetRelations(ctx, location, config.Cfg.DB.GetLocationAeonIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.Shops, err = apilookups.E.Shops.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ts.VstShop))
		return err
	})

	g.Go(func() error {
		var err error
		response.Treasures, err = apilookups.E.Treasures.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ts.VstTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		response.Monsters, err = apilookups.E.Monsters.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ts.VstMonster))
		return err
	})

	g.Go(func() error {
		var err error
		response.Formations, err = apilookups.E.MonsterFormations.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ts.VstMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		response.Quests, err = apilookups.E.Quests.QueryAvlRelRun(ctx, location, availabilityParams, getLocationRelSourceIDs(ts.VstQuest))
		return err
	})

	g.Go(func() error {
		var err error
		response.Music, err = getMusicLocBased(ctx, location, locBasedMusicQueries{
			cueSongs:  config.Cfg.DB.GetLocationCueSongIDs,
			bmSongs:   config.Cfg.DB.GetLocationBackgroundMusicSongIDs,
			fmvSongs:  config.Cfg.DB.GetLocationFMVSongIDs,
			bossMusic: config.Cfg.DB.GetLocationBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		response.FMVs, err = apilookups.E.FMVs.DBGetRelations(ctx, location, config.Cfg.DB.GetLocationFmvIDs)
		return err
	})

	return g.Wait()
}
