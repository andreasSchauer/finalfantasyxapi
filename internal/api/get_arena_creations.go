package api

import (
	"net/http"
	"sync"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func (cfg *Config) getArenaCreation(r *http.Request, i handlerInput[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource], id int32) (ArenaCreation, error) {
	creation, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return ArenaCreation{}, err
	}

	subquest, _ := lookup.GetResourceByID(creation.SubquestID, lookup.L.SubquestsID)
	ia := subquest.Completion.Reward

	monsters, err := getArenaCreationMonsters(cfg, r, creation)
	if err != nil {
		return ArenaCreation{}, err
	}

	response := ArenaCreation{
		ID:                        creation.ID,
		Name:                      creation.Name,
		Category:                  creation.Category,
		Monster:                   idToNamedAPIResource(cfg, cfg.E.Monsters, *creation.MonsterID),
		ParentSubquest:            idToQuestAPIResource(cfg, cfg.E.Subquests, creation.SubquestID),
		Reward:                    cfg.E.AllItems.toResAmtFromName(cfg, ia),
		RequiredCatchAmount:       creation.Amount,
		UnlockedCreationsCategory: creation.CreationsUnlockedCategory,
		RequiredMonsters:          monsters,
	}

	return response, nil
}

func (cfg *Config) retrieveArenaCreations(r *http.Request, i handlerInput[seedtypes.ArenaCreation, ArenaCreation, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, cfg.T.ArenaCreationCategory, ids, QpnCategory, cfg.db.GetArenaCreationIDsByCategory),
	})
}

func getArenaCreationMonsters(cfg *Config, r *http.Request, creation seedtypes.ArenaCreation) ([]NamedAPIResource, error) {
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
			idSlice, err := cfg.db.GetCaptureMonsterIDsByMaCreationArea(ctx, area)
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
			idSlice, err := cfg.db.GetCaptureMonsterIDsBySpecies(ctx, species)
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
			idSlice, err := cfg.db.GetCaptureMonsterIDsByIsUnderwater(ctx)
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

	monsters := cfg.E.Monsters.toResources(cfg, monsterIDs)
	return monsters, nil
}
