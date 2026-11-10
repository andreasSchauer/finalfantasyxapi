package get

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetArea(r *http.Request, i handlerinput.Main[seedtypes.Area, apitypes.Area, apiresource.Area], id int32) (apitypes.Area, error) {
	area, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Area{}, err
	}

	connections, err := GetAreaConnectedAreas(area)
	if err != nil {
		return apitypes.Area{}, err
	}

	rel, err := getAreaRelationships(r, area)
	if err != nil {
		return apitypes.Area{}, err
	}

	response := apitypes.Area{
		ID:                area.ID,
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		DisplayName:       GetAreaDisplayName(area),
		ParentLocation:    handlerinput.NameToNamedAPIResource(apilookups.E.Locations, area.Location, nil),
		ParentSublocation: handlerinput.NameToNamedAPIResource(apilookups.E.Sublocations, area.Sublocation, nil),
		Availability:      area.Availability,
		HasSaveSphere:     area.HasSaveSphere,
		AirshipDropOff:    area.AirshipDropOff,
		HasCompSphere:     area.HasCompilationSphere,
		CanRideChocobo:    area.CanRideChocobo,
		ConnectedAreas:    connections,
		LocRel:            rel,
	}

	return response, nil
}

func getAreaRelationships(r *http.Request, area seedtypes.Area) (apitypes.LocRel, error) {
	var rel apitypes.LocRel
	g, ctx := errgroup.WithContext(r.Context())

	availabilityParams, err := apilookups.E.Areas.QueryAvlRelGetParams(r, area.ID)
	if err != nil {
		return apitypes.LocRel{}, err
	}

	g.Go(func() error {
		var err error
		rel.Characters, err = apilookups.E.Characters.DBGetRelations(ctx, area, handlerinput.ToIntManyNull(config.Cfg.DB.GetAreaCharacterIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Aeons, err = apilookups.E.Aeons.DBGetRelations(ctx, area, handlerinput.ToIntManyNull(config.Cfg.DB.GetAreaAeonIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Shops, err = apilookups.E.Shops.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ts.VstShop))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Treasures, err = apilookups.E.Treasures.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ts.VstTreasure))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Monsters, err = apilookups.E.Monsters.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ts.VstMonster))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Formations, err = apilookups.E.MonsterFormations.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ts.VstMonsterFormation))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Quests, err = apilookups.E.Quests.QueryAvlRelRun(ctx, area, availabilityParams, getAreaRelSourceIDs(ts.VstQuest))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Music, err = getMusicLocBased(ctx, area, LocBasedMusicQueries{
			CueSongs:  handlerinput.ToIntManyNull(config.Cfg.DB.GetAreaCueSongIDs),
			BmSongs:   config.Cfg.DB.GetAreaBackgroundMusicSongIDs,
			FMVSongs:  config.Cfg.DB.GetAreaFMVSongIDs,
			BossMusic: config.Cfg.DB.GetAreaBossSongIDs,
		})
		return err
	})

	g.Go(func() error {
		var err error
		rel.FMVs, err = apilookups.E.FMVs.DBGetRelations(ctx, area, config.Cfg.DB.GetAreaFmvIDs)
		return err
	})

	err = g.Wait()
	if err != nil {
		return apitypes.LocRel{}, err
	}

	return rel, nil
}

func GetAreaConnectedAreas(area seedtypes.Area) ([]apitypes.AreaConnection, error) {
	i := apilookups.E.Areas
	connectedAreas := []apitypes.AreaConnection{}

	for _, connArea := range area.ConnectedAreas {
		connection := apitypes.AreaConnection{
			Area:           handlerinput.LocAreaToAreaAPIResource(i, connArea.LocationArea),
			ConnectionType: connArea.ConnectionType,
			IsStoryBased:   connArea.IsStoryBased,
			Notes:          connArea.Notes,
		}

		connectedAreas = append(connectedAreas, connection)
	}

	return connectedAreas, nil
}

func GetAreaDisplayName(area seedtypes.Area) string {
	sublocName := area.Sublocation

	if sublocName == area.Name {
		return area.Name
	}

	return fmt.Sprintf("%s - %s", sublocName, area.Name)
}
