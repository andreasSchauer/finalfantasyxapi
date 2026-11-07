package api

import (
	"net/http"
	"sync"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getArenaCreation(r *http.Request, i HandlerInput[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.NamedAPIResource], id int32) (apitypes.ArenaCreation, error) {
	creation, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	subquest, _ := lookup.GetResourceByID(creation.SubquestID, lookup.L.SubquestsID)
	ia := subquest.Completion.Reward

	monsters, err := getArenaCreationMonsters(r, creation)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	response := apitypes.ArenaCreation{
		ID:                        creation.ID,
		Name:                      creation.Name,
		Category:                  creation.Category,
		Monster:                   IDToNamedAPIResource(E.Monsters, *creation.MonsterID),
		ParentSubquest:            IDToQuestAPIResource(E.Subquests, creation.SubquestID),
		Reward:                    E.AllItems.ToResAmtFromName(ia),
		RequiredCatchAmount:       creation.Amount,
		UnlockedCreationsCategory: creation.CreationsUnlockedCategory,
		RequiredMonsters:          monsters,
	}

	return response, nil
}

func retrieveArenaCreations(r *http.Request, i HandlerInput[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryEnumList(r, enums.T.ArenaCreationCategory, ids, ts.QpnCategory, config.Cfg.DB.GetArenaCreationIDsByCategory),
	})
}

func getArenaCreationMonsters(r *http.Request, creation seedtypes.ArenaCreation) ([]apiresource.NamedAPIResource, error) {
	mf := createMonsterFilter(creation)

	if mf.IsZero() || mf.CreationsUnlockedCategory != nil {
		return nil, nil
	}

	monsterIDSlices := [][]int32{}
	g, ctx := errgroup.WithContext(r.Context())
	var mu sync.Mutex

	if mf.RequiredArea != nil {
		g.Go(func() error {
			area := database.ToNullMaCreationArea(mf.RequiredArea)
			idSlice, err := config.Cfg.DB.GetCaptureMonsterIDsByMaCreationArea(ctx, area)
			if err != nil {
				return err
			}

			mu.Lock()
			monsterIDSlices = append(monsterIDSlices, idSlice)
			mu.Unlock()
			return nil
		})
	}

	if mf.RequiredSpecies != nil {
		g.Go(func() error {
			species := database.MonsterSpecies(*mf.RequiredSpecies)
			idSlice, err := config.Cfg.DB.GetCaptureMonsterIDsBySpecies(ctx, species)
			if err != nil {
				return err
			}

			mu.Lock()
			monsterIDSlices = append(monsterIDSlices, idSlice)
			mu.Unlock()
			return nil
		})
	}

	if mf.UnderwaterOnly {
		g.Go(func() error {
			idSlice, err := config.Cfg.DB.GetCaptureMonsterIDsByIsUnderwater(ctx)
			if err != nil {
				return err
			}

			mu.Lock()
			monsterIDSlices = append(monsterIDSlices, idSlice)
			mu.Unlock()
			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	monsterIDs, err := ah.FilterIDSlices(monsterIDSlices)
	if err != nil {
		return nil, err
	}

	monsters := E.Monsters.ToAPIResources(monsterIDs)
	return monsters, nil
}
