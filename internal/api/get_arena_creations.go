package api

import (
	"net/http"
	"sync"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func (srv *Server) getArenaCreation(r *http.Request, i handlerInput[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource], id int32) (ArenaCreation, error) {
	creation, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return ArenaCreation{}, err
	}

	subquest, _ := lookup.GetResourceByID(creation.SubquestID, lookup.L.SubquestsID)
	ia := subquest.Completion.Reward

	monsters, err := getArenaCreationMonsters(srv, r, creation)
	if err != nil {
		return ArenaCreation{}, err
	}

	response := ArenaCreation{
		ID:                        creation.ID,
		Name:                      creation.Name,
		Category:                  creation.Category,
		Monster:                   idToNamedAPIResource(srv, E.Monsters, *creation.MonsterID),
		ParentSubquest:            idToQuestAPIResource(srv, E.Subquests, creation.SubquestID),
		Reward:                    E.AllItems.toResAmtFromName(srv, ia),
		RequiredCatchAmount:       creation.Amount,
		UnlockedCreationsCategory: creation.CreationsUnlockedCategory,
		RequiredMonsters:          monsters,
	}

	return response, nil
}

func (srv *Server) retrieveArenaCreations(r *http.Request, i handlerInput[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.ArenaCreationCategory, ids, ts.QpnCategory, config.Cfg.DB.GetArenaCreationIDsByCategory),
	})
}

func getArenaCreationMonsters(srv *Server, r *http.Request, creation seedtypes.ArenaCreation) ([]NamedAPIResource, error) {
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

	monsterIDs, err := filterIDSlices(monsterIDSlices)
	if err != nil {
		return nil, err
	}

	monsters := E.Monsters.toResources(srv, monsterIDs)
	return monsters, nil
}
